// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingConnectionKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_CONNECTION_ID:
		t = ktypeString
	case NM_SETTING_CONNECTION_UUID:
		t = ktypeString
	case NM_SETTING_CONNECTION_TYPE:
		t = ktypeString
	case NM_SETTING_CONNECTION_AUTOCONNECT:
		t = ktypeBoolean
	case NM_SETTING_CONNECTION_PERMISSIONS:
		t = ktypeArrayString
	case NM_SETTING_CONNECTION_TIMESTAMP:
		t = ktypeUint64
	case NM_SETTING_CONNECTION_READ_ONLY:
		t = ktypeBoolean
	case NM_SETTING_CONNECTION_ZONE:
		t = ktypeString
	case NM_SETTING_CONNECTION_MASTER:
		t = ktypeString
	case NM_SETTING_CONNECTION_SLAVE_TYPE:
		t = ktypeString
	case NM_SETTING_CONNECTION_SECONDARIES:
		t = ktypeArrayString
	}
	return
}

// Check is key in current setting field
func isKeyInSettingConnection(key string) bool {
	switch key {
	case NM_SETTING_CONNECTION_ID:
		return true
	case NM_SETTING_CONNECTION_UUID:
		return true
	case NM_SETTING_CONNECTION_TYPE:
		return true
	case NM_SETTING_CONNECTION_AUTOCONNECT:
		return true
	case NM_SETTING_CONNECTION_PERMISSIONS:
		return true
	case NM_SETTING_CONNECTION_TIMESTAMP:
		return true
	case NM_SETTING_CONNECTION_READ_ONLY:
		return true
	case NM_SETTING_CONNECTION_ZONE:
		return true
	case NM_SETTING_CONNECTION_MASTER:
		return true
	case NM_SETTING_CONNECTION_SLAVE_TYPE:
		return true
	case NM_SETTING_CONNECTION_SECONDARIES:
		return true
	}
	return false
}

// Get key's default value
func getSettingConnectionKeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_CONNECTION_ID:
		valueJSON = `""`
	case NM_SETTING_CONNECTION_UUID:
		valueJSON = `""`
	case NM_SETTING_CONNECTION_TYPE:
		valueJSON = `""`
	case NM_SETTING_CONNECTION_AUTOCONNECT:
		valueJSON = `true`
	case NM_SETTING_CONNECTION_PERMISSIONS:
		valueJSON = `null`
	case NM_SETTING_CONNECTION_TIMESTAMP:
		valueJSON = `0`
	case NM_SETTING_CONNECTION_READ_ONLY:
		valueJSON = `false`
	case NM_SETTING_CONNECTION_ZONE:
		valueJSON = `""`
	case NM_SETTING_CONNECTION_MASTER:
		valueJSON = `""`
	case NM_SETTING_CONNECTION_SLAVE_TYPE:
		valueJSON = `""`
	case NM_SETTING_CONNECTION_SECONDARIES:
		valueJSON = `null`
	}
	return
}

// Get JSON value generally
func generalGetSettingConnectionKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingConnectionKeyJSON: invalide key", key)
	case NM_SETTING_CONNECTION_ID:
		value = getSettingConnectionIdJSON(data)
	case NM_SETTING_CONNECTION_UUID:
		value = getSettingConnectionUuidJSON(data)
	case NM_SETTING_CONNECTION_TYPE:
		value = getSettingConnectionTypeJSON(data)
	case NM_SETTING_CONNECTION_AUTOCONNECT:
		value = getSettingConnectionAutoconnectJSON(data)
	case NM_SETTING_CONNECTION_PERMISSIONS:
		value = getSettingConnectionPermissionsJSON(data)
	case NM_SETTING_CONNECTION_TIMESTAMP:
		value = getSettingConnectionTimestampJSON(data)
	case NM_SETTING_CONNECTION_READ_ONLY:
		value = getSettingConnectionReadOnlyJSON(data)
	case NM_SETTING_CONNECTION_ZONE:
		value = getSettingConnectionZoneJSON(data)
	case NM_SETTING_CONNECTION_MASTER:
		value = getSettingConnectionMasterJSON(data)
	case NM_SETTING_CONNECTION_SLAVE_TYPE:
		value = getSettingConnectionSlaveTypeJSON(data)
	case NM_SETTING_CONNECTION_SECONDARIES:
		value = getSettingConnectionSecondariesJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingConnectionKeyJSON(data connectionData, key, valueJSON string) {
	switch key {
	default:
		logger.Error("generalSetSettingConnectionKeyJSON: invalide key", key)
	case NM_SETTING_CONNECTION_ID:
		setSettingConnectionIdJSON(data, valueJSON)
	case NM_SETTING_CONNECTION_UUID:
		setSettingConnectionUuidJSON(data, valueJSON)
	case NM_SETTING_CONNECTION_TYPE:
		setSettingConnectionTypeJSON(data, valueJSON)
	case NM_SETTING_CONNECTION_AUTOCONNECT:
		setSettingConnectionAutoconnectJSON(data, valueJSON)
	case NM_SETTING_CONNECTION_PERMISSIONS:
		setSettingConnectionPermissionsJSON(data, valueJSON)
	case NM_SETTING_CONNECTION_TIMESTAMP:
		setSettingConnectionTimestampJSON(data, valueJSON)
	case NM_SETTING_CONNECTION_READ_ONLY:
		setSettingConnectionReadOnlyJSON(data, valueJSON)
	case NM_SETTING_CONNECTION_ZONE:
		setSettingConnectionZoneJSON(data, valueJSON)
	case NM_SETTING_CONNECTION_MASTER:
		setSettingConnectionMasterJSON(data, valueJSON)
	case NM_SETTING_CONNECTION_SLAVE_TYPE:
		setSettingConnectionSlaveTypeJSON(data, valueJSON)
	case NM_SETTING_CONNECTION_SECONDARIES:
		setSettingConnectionSecondariesJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingConnectionIdExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ID)
}
func isSettingConnectionUuidExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_UUID)
}
func isSettingConnectionTypeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TYPE)
}
func isSettingConnectionAutoconnectExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_AUTOCONNECT)
}
func isSettingConnectionPermissionsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS)
}
func isSettingConnectionTimestampExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TIMESTAMP)
}
func isSettingConnectionReadOnlyExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_READ_ONLY)
}
func isSettingConnectionZoneExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ZONE)
}
func isSettingConnectionMasterExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_MASTER)
}
func isSettingConnectionSlaveTypeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SLAVE_TYPE)
}
func isSettingConnectionSecondariesExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SECONDARIES)
}

// Ensure field and key exists and not empty
func ensureFieldSettingConnectionExists(data connectionData, errs FieldKeyErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_CONNECTION_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_CONNECTION_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_CONNECTION_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_CONNECTION_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_CONNECTION_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_CONNECTION_SETTING_NAME))
	}
}
func ensureSettingConnectionIdNoEmpty(data connectionData, errs FieldKeyErrors) {
	if !isSettingConnectionIdExists(data) {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ID, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingConnectionId(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ID, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingConnectionUuidNoEmpty(data connectionData, errs FieldKeyErrors) {
	if !isSettingConnectionUuidExists(data) {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_UUID, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingConnectionUuid(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_UUID, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingConnectionTypeNoEmpty(data connectionData, errs FieldKeyErrors) {
	if !isSettingConnectionTypeExists(data) {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TYPE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingConnectionType(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TYPE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingConnectionAutoconnectNoEmpty(data connectionData, errs FieldKeyErrors) {
	if !isSettingConnectionAutoconnectExists(data) {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_AUTOCONNECT, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingConnectionPermissionsNoEmpty(data connectionData, errs FieldKeyErrors) {
	if !isSettingConnectionPermissionsExists(data) {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingConnectionPermissions(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingConnectionTimestampNoEmpty(data connectionData, errs FieldKeyErrors) {
	if !isSettingConnectionTimestampExists(data) {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TIMESTAMP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingConnectionReadOnlyNoEmpty(data connectionData, errs FieldKeyErrors) {
	if !isSettingConnectionReadOnlyExists(data) {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_READ_ONLY, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingConnectionZoneNoEmpty(data connectionData, errs FieldKeyErrors) {
	if !isSettingConnectionZoneExists(data) {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ZONE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingConnectionZone(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ZONE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingConnectionMasterNoEmpty(data connectionData, errs FieldKeyErrors) {
	if !isSettingConnectionMasterExists(data) {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_MASTER, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingConnectionMaster(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_MASTER, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingConnectionSlaveTypeNoEmpty(data connectionData, errs FieldKeyErrors) {
	if !isSettingConnectionSlaveTypeExists(data) {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SLAVE_TYPE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingConnectionSlaveType(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SLAVE_TYPE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingConnectionSecondariesNoEmpty(data connectionData, errs FieldKeyErrors) {
	if !isSettingConnectionSecondariesExists(data) {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SECONDARIES, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingConnectionSecondaries(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SECONDARIES, NM_KEY_ERROR_EMPTY_VALUE)
	}
}

// Getter
func getSettingConnectionId(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ID).(string)
	return
}
func getSettingConnectionUuid(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_UUID).(string)
	return
}
func getSettingConnectionType(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TYPE).(string)
	return
}
func getSettingConnectionAutoconnect(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_AUTOCONNECT).(bool)
	return
}
func getSettingConnectionPermissions(data connectionData) (value []string) {
	value, _ = getSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS).([]string)
	return
}
func getSettingConnectionTimestamp(data connectionData) (value uint64) {
	value, _ = getSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TIMESTAMP).(uint64)
	return
}
func getSettingConnectionReadOnly(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_READ_ONLY).(bool)
	return
}
func getSettingConnectionZone(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ZONE).(string)
	return
}
func getSettingConnectionMaster(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_MASTER).(string)
	return
}
func getSettingConnectionSlaveType(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SLAVE_TYPE).(string)
	return
}
func getSettingConnectionSecondaries(data connectionData) (value []string) {
	value, _ = getSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SECONDARIES).([]string)
	return
}

// Setter
func setSettingConnectionId(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ID, value)
}
func setSettingConnectionUuid(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_UUID, value)
}
func setSettingConnectionType(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TYPE, value)
}
func setSettingConnectionAutoconnect(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_AUTOCONNECT, value)
}
func setSettingConnectionPermissions(data connectionData, value []string) {
	setSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS, value)
}
func setSettingConnectionTimestamp(data connectionData, value uint64) {
	setSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TIMESTAMP, value)
}
func setSettingConnectionReadOnly(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_READ_ONLY, value)
}
func setSettingConnectionZone(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ZONE, value)
}
func setSettingConnectionMaster(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_MASTER, value)
}
func setSettingConnectionSlaveType(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SLAVE_TYPE, value)
}
func setSettingConnectionSecondaries(data connectionData, value []string) {
	setSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SECONDARIES, value)
}

// JSON Getter
func getSettingConnectionIdJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ID, getSettingConnectionKeyType(NM_SETTING_CONNECTION_ID))
	return
}
func getSettingConnectionUuidJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_UUID, getSettingConnectionKeyType(NM_SETTING_CONNECTION_UUID))
	return
}
func getSettingConnectionTypeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TYPE, getSettingConnectionKeyType(NM_SETTING_CONNECTION_TYPE))
	return
}
func getSettingConnectionAutoconnectJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_AUTOCONNECT, getSettingConnectionKeyType(NM_SETTING_CONNECTION_AUTOCONNECT))
	return
}
func getSettingConnectionPermissionsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS, getSettingConnectionKeyType(NM_SETTING_CONNECTION_PERMISSIONS))
	return
}
func getSettingConnectionTimestampJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TIMESTAMP, getSettingConnectionKeyType(NM_SETTING_CONNECTION_TIMESTAMP))
	return
}
func getSettingConnectionReadOnlyJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_READ_ONLY, getSettingConnectionKeyType(NM_SETTING_CONNECTION_READ_ONLY))
	return
}
func getSettingConnectionZoneJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ZONE, getSettingConnectionKeyType(NM_SETTING_CONNECTION_ZONE))
	return
}
func getSettingConnectionMasterJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_MASTER, getSettingConnectionKeyType(NM_SETTING_CONNECTION_MASTER))
	return
}
func getSettingConnectionSlaveTypeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SLAVE_TYPE, getSettingConnectionKeyType(NM_SETTING_CONNECTION_SLAVE_TYPE))
	return
}
func getSettingConnectionSecondariesJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SECONDARIES, getSettingConnectionKeyType(NM_SETTING_CONNECTION_SECONDARIES))
	return
}

// JSON Setter
func setSettingConnectionIdJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ID, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_ID))
}
func setSettingConnectionUuidJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_UUID, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_UUID))
}
func setSettingConnectionTypeJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TYPE, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_TYPE))
}
func setSettingConnectionAutoconnectJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_AUTOCONNECT, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_AUTOCONNECT))
}
func setSettingConnectionPermissionsJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_PERMISSIONS))
}
func setSettingConnectionTimestampJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TIMESTAMP, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_TIMESTAMP))
}
func setSettingConnectionReadOnlyJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_READ_ONLY, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_READ_ONLY))
}
func setSettingConnectionZoneJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ZONE, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_ZONE))
}
func setSettingConnectionMasterJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_MASTER, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_MASTER))
}
func setSettingConnectionSlaveTypeJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SLAVE_TYPE, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_SLAVE_TYPE))
}
func setSettingConnectionSecondariesJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SECONDARIES, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_SECONDARIES))
}

// Remover
func removeSettingConnectionId(data connectionData) {
	removeSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ID)
}
func removeSettingConnectionUuid(data connectionData) {
	removeSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_UUID)
}
func removeSettingConnectionType(data connectionData) {
	removeSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TYPE)
}
func removeSettingConnectionAutoconnect(data connectionData) {
	removeSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_AUTOCONNECT)
}
func removeSettingConnectionPermissions(data connectionData) {
	removeSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS)
}
func removeSettingConnectionTimestamp(data connectionData) {
	removeSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TIMESTAMP)
}
func removeSettingConnectionReadOnly(data connectionData) {
	removeSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_READ_ONLY)
}
func removeSettingConnectionZone(data connectionData) {
	removeSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ZONE)
}
func removeSettingConnectionMaster(data connectionData) {
	removeSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_MASTER)
}
func removeSettingConnectionSlaveType(data connectionData) {
	removeSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SLAVE_TYPE)
}
func removeSettingConnectionSecondaries(data connectionData) {
	removeSettingKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SECONDARIES)
}
