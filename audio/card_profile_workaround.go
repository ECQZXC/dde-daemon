/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package audio

import (
	"pkg.deepin.io/lib/pulse"
	"pkg.deepin.io/lib/strv"
	"sort"
)

const (
	CardBuildin   = 0
	CardBluethooh = 1
	CardUnknow    = 2

	PropDeviceFromFactor = "device.form_factor"
	PropDeviceBus        = "device.bus"
)

func cardType(c *pulse.Card) int {
	if c.PropList[PropDeviceFromFactor] == "internal" {
		return CardBuildin
	}
	if c.PropList[PropDeviceBus] == "bluetooth" {
		return CardBluethooh
	}
	return CardUnknow
}

func profileBlacklist(c *pulse.Card) strv.Strv {
	var blacklist []string
	switch cardType(c) {
	case CardBluethooh:
		// TODO: bluez not full support headset_head_unit, please skip
		blacklist = []string{"off", "headset_head_unit"}
	default:
		// CardBuildin, CardUnknow and other
		blacklist = []string{"off"}
	}
	return strv.Strv(blacklist)
}

//select New Card Profile By priority, protocl.
func selectNewCardProfile(c *pulse.Card) {
	blacklist := profileBlacklist(c)
	if !blacklist.Contains(c.ActiveProfile.Name) {
		logger.Debug("use profile:", c.ActiveProfile)
		return
	}

	var profiles cProfileInfos2
	for _, p := range c.Profiles {
		// skip the profile in the blacklist
		if blacklist.Contains(p.Name) {
			continue
		}
		profiles = append(profiles, p)
	}

	// sort profiles by priority
	logger.Debug("[selectNewCardProfile] before sort:", profiles)
	sort.Sort(profiles)
	logger.Debug("[selectNewCardProfile] after sort:", profiles)

	// if card is bluetooth device, switch to profile a2dp_sink
	// only 'a2dp_sink' in bluetooth profiles because of blacklist
	if len(profiles) > 0 {
		logger.Debug("re-select card profile:", profiles[0], c.ActiveProfile.Name)
		if c.ActiveProfile.Name != profiles[0].Name {
			c.SetProfile(profiles[0].Name)
		}
	}
}
