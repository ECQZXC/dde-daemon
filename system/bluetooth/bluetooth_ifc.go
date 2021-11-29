package bluetooth

import (
	"fmt"
	"strconv"

	"github.com/godbus/dbus"
	sysbtagent "github.com/linuxdeepin/go-dbus-factory/com.deepin.system.bluetooth.agent"
	login1 "github.com/linuxdeepin/go-dbus-factory/org.freedesktop.login1"
	"github.com/linuxdeepin/go-lib/dbusutil"
)

func (b *SysBluetooth) ConnectDevice(devPath dbus.ObjectPath, adapterPath dbus.ObjectPath) *dbus.Error {
	device, err := b.getDevice(devPath)
	if err != nil {
		logger.Debug("getDevice failed:", err)
		adapter, err := b.getAdapter(adapterPath)
		if err != nil {
			logger.Debug("getAdapter failed:", err)
			return dbusutil.ToError(err)
		}
		adapter.startDiscovery()
		adapter.scanReadyToConnectDeviceTimeoutFlag = true
		b.prepareToConnectedMu.Lock()
		b.prepareToConnectedDevice = devPath
		b.prepareToConnectedMu.Unlock()
	} else {
		go func() {
			err := device.Connect()
			if err != nil {
				logger.Warning(err)
			}
		}()
	}
	return nil
}

func (b *SysBluetooth) DisconnectDevice(devPath dbus.ObjectPath) *dbus.Error {
	device, err := b.getDevice(devPath)
	if err != nil {
		return dbusutil.ToError(err)
	}
	go device.Disconnect()
	return nil
}

func (b *SysBluetooth) RemoveDevice(adapterPath, devPath dbus.ObjectPath) *dbus.Error {
	adapter, err := b.getAdapter(adapterPath)
	if err != nil {
		return dbusutil.ToError(err)
	}
	// find remove device from map
	removeDev, err := b.getDevice(devPath)
	if err != nil {
		logger.Warningf("failed to get device, err: %v", err)
		return dbusutil.ToError(err)
	}
	// check if device connect state is connecting, if is, mark remove state as true
	deviceState := removeDev.getState()
	if deviceState == deviceStateConnecting {
		removeDev.markNeedRemove(true)
	} else {
		// connection finish, allow removing device directly
		b.removeBackupDevice(devPath)
		err = adapter.core.Adapter().RemoveDevice(0, devPath)
		if err != nil {
			logger.Warningf("failed to remove device %q from adapter %q: %v",
				devPath, adapterPath, err)
			return dbusutil.ToError(err)
		}
	}
	return nil
}

func (b *SysBluetooth) SetDeviceAlias(device dbus.ObjectPath, alias string) *dbus.Error {
	d, err := b.getDevice(device)
	if err != nil {
		return dbusutil.ToError(err)
	}
	err = d.core.Alias().Set(0, alias)
	if err != nil {
		return dbusutil.ToError(err)
	}
	return nil
}

func (b *SysBluetooth) SetDeviceTrusted(device dbus.ObjectPath, trusted bool) *dbus.Error {
	d, err := b.getDevice(device)
	if err != nil {
		return dbusutil.ToError(err)
	}
	err = d.core.Trusted().Set(0, trusted)
	if err != nil {
		return dbusutil.ToError(err)
	}
	return nil
}

// GetDevices return all device objects that marshaled by json.
func (b *SysBluetooth) GetDevices(adapterPath dbus.ObjectPath) (devicesJSON string, busErr *dbus.Error) {
	_, err := b.getAdapter(adapterPath)
	if err != nil {
		return "", dbusutil.ToError(err)
	}

	var deviceMap = make(map[dbus.ObjectPath]*backupDevice)
	b.backupDevicesMu.Lock()
	for _, d := range b.backupDevices[adapterPath] {
		deviceMap[d.Path] = d
	}
	b.backupDevicesMu.Unlock()

	b.devicesMu.Lock()
	for _, d := range b.devices[adapterPath] {
		deviceMap[d.Path] = newBackupDevice(d)
	}
	b.devicesMu.Unlock()

	var devices []*backupDevice
	for _, d := range deviceMap {
		devices = append(devices, d)
	}
	devicesJSON = marshalJSON(devices)
	return
}

// GetAdapters return all adapter objects that marshaled by json.
func (b *SysBluetooth) GetAdapters() (adaptersJSON string, err *dbus.Error) {
	adapters := make([]*adapter, 0, len(b.adapters))
	b.adaptersMu.Lock()
	for _, a := range b.adapters {
		adapters = append(adapters, a)
	}
	b.adaptersMu.Unlock()
	adaptersJSON = marshalJSON(adapters)
	return
}

func (b *SysBluetooth) RequestDiscovery(adapterPath dbus.ObjectPath) *dbus.Error {
	adapter, err := b.getAdapter(adapterPath)
	if err != nil {
		return dbusutil.ToError(err)
	}

	if !adapter.Powered {
		err = fmt.Errorf("'%s' power off", adapter)
		return dbusutil.ToError(err)
	}

	discovering, err := adapter.core.Adapter().Discovering().Get(0)
	if err != nil {
		return dbusutil.ToError(err)
	}

	if discovering {
		// if adapter is discovering now, just return
		return nil
	}

	adapter.startDiscovery()

	return nil
}

func (b *SysBluetooth) SetAdapterPowered(adapterPath dbus.ObjectPath,
	powered bool) *dbus.Error {

	logger.Debug("SetAdapterPowered", adapterPath, powered)

	adapter, err := b.getAdapter(adapterPath)
	if err != nil {
		return dbusutil.ToError(err)
	}
	adapter.discoveringFinished = false

	err = adapter.core.Adapter().Powered().Set(0, powered)
	if err != nil {
		logger.Warningf("failed to set %s powered: %v", adapter, err)
		return dbusutil.ToError(err)
	}
	_bt.config.setAdapterConfigPowered(adapter.Address, powered)

	return nil
}

func (b *SysBluetooth) SetAdapterAlias(adapterPath dbus.ObjectPath, alias string) *dbus.Error {
	adapter, err := b.getAdapter(adapterPath)
	if err != nil {
		return dbusutil.ToError(err)
	}

	err = adapter.core.Adapter().Alias().Set(0, alias)
	if err != nil {
		logger.Warningf("failed to set %s alias: %v", adapter, err)
		return dbusutil.ToError(err)
	}

	return nil
}

func (b *SysBluetooth) SetAdapterDiscoverable(adapterPath dbus.ObjectPath,
	discoverable bool) *dbus.Error {
	logger.Debug("SetAdapterDiscoverable", adapterPath, discoverable)

	adapter, err := b.getAdapter(adapterPath)
	if err != nil {
		return dbusutil.ToError(err)
	}

	if !adapter.Powered {
		err = fmt.Errorf("'%s' power off", adapter)
		return dbusutil.ToError(err)
	}

	err = adapter.core.Adapter().Discoverable().Set(0, discoverable)
	if err != nil {
		logger.Warningf("failed to set %s discoverable: %v", adapter, err)
		return dbusutil.ToError(err)
	}
	_bt.config.setAdapterConfigDiscoverable(adapter.Address, discoverable)

	return nil
}

func (b *SysBluetooth) SetAdapterDiscovering(adapterPath dbus.ObjectPath,
	discovering bool) *dbus.Error {
	logger.Debug("SetAdapterDiscovering", adapterPath, discovering)

	adapter, err := b.getAdapter(adapterPath)
	if err != nil {
		return dbusutil.ToError(err)
	}

	if !adapter.Powered {
		err = fmt.Errorf("'%s' power off", adapter)
		return dbusutil.ToError(err)
	}

	if discovering {
		adapter.startDiscovery()
	} else {
		err = adapter.core.Adapter().StopDiscovery(0)
		if err != nil {
			logger.Warningf("failed to stop discovery for %s: %v", adapter, err)
			return dbusutil.ToError(err)
		}
	}

	return nil
}

func (b *SysBluetooth) SetAdapterDiscoverableTimeout(adapterPath dbus.ObjectPath,
	discoverableTimeout uint32) *dbus.Error {
	logger.Debug("SetAdapterDiscoverableTimeout", adapterPath, discoverableTimeout)

	adapter, err := b.getAdapter(adapterPath)
	if err != nil {
		return dbusutil.ToError(err)
	}

	err = adapter.core.Adapter().DiscoverableTimeout().Set(0, discoverableTimeout)
	if err != nil {
		logger.Warningf("failed to set %s discoverableTimeout: %v", adapter, err)
		return dbusutil.ToError(err)
	}

	return nil
}

func (b *SysBluetooth) getActiveUserAgent() sysbtagent.Agent {
	return b.userAgents.getActiveAgent()
}

func (b *SysBluetooth) RegisterAgent(sender dbus.Sender, agentPath dbus.ObjectPath) *dbus.Error {
	uid, err := b.service.GetConnUID(string(sender))
	if err != nil {
		logger.Warning(err)
		return dbusutil.ToError(err)
	}

	uidStr := strconv.Itoa(int(uid))
	b.userAgents.addUser(uidStr)

	sessionDetails, err := b.loginManager.ListSessions(0)
	if err != nil {
		logger.Warning(err)
		return dbusutil.ToError(err)
	}

	sysBus := b.service.Conn()
	for _, detail := range sessionDetails {
		if detail.UID == uid {
			session, err := login1.NewSession(sysBus, detail.Path)
			if err != nil {
				logger.Warning(err)
				continue
			}
			newlyAdded := b.userAgents.addSession(uidStr, session)
			if newlyAdded {
				b.watchSession(uidStr, session)
			}
		}
	}

	agent, err := sysbtagent.NewAgent(sysBus, string(sender), agentPath)
	if err != nil {
		logger.Warning(err)
		return dbusutil.ToError(err)
	}
	b.userAgents.addAgent(uidStr, agent)
	// 防止第一次进入系统，此时无设备连接，但是 needFixBtPoweredStatus 为true，此时打开蓝牙添加设备后paired为true
	// 此时在调用addDevice接口后，会走不必要的逻辑，因此将此标志位置为false规避
	b.needFixBtPoweredStatus = false
	go b.tryConnectPairedDevices("")

	logger.Debugf("agent registered, sender: %q, agentPath: %q", sender, agentPath)
	return nil
}

func (b *SysBluetooth) UnregisterAgent(sender dbus.Sender, agentPath dbus.ObjectPath) *dbus.Error {
	uid, err := b.service.GetConnUID(string(sender))
	if err != nil {
		logger.Warning(err)
		return dbusutil.ToError(err)
	}

	uidStr := strconv.Itoa(int(uid))
	err = b.userAgents.removeAgent(uidStr, agentPath)
	if err != nil {
		logger.Warning(err)
		return dbusutil.ToError(err)
	}
	logger.Debugf("agent unregistered, sender: %q, agentPath: %q", sender, agentPath)
	return nil
}

func (b *SysBluetooth) DebugInfo() (info string, busErr *dbus.Error) {
	info = fmt.Sprintf("adapters: %s\ndevices: %s", marshalJSON(b.adapters), marshalJSON(b.devices))
	return info, nil
}

//ClearUnpairedDevice will remove all device in unpaired list
func (b *SysBluetooth) ClearUnpairedDevice() *dbus.Error {
	logger.Debug("ClearUnpairedDevice")
	var removeDevices []*device
	b.devicesMu.Lock()
	for _, devices := range b.devices {
		for _, d := range devices {
			if !d.Paired {
				logger.Info("remove unpaired device", d)
				removeDevices = append(removeDevices, d)
			}
		}
	}
	b.devicesMu.Unlock()

	for _, d := range removeDevices {
		err := b.RemoveDevice(d.AdapterPath, d.Path)
		if err != nil {
			logger.Warning(err)
		}
	}
	return nil
}