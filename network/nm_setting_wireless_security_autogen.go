// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingWirelessSecurityKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_WIRELESS_SECURITY_KEY_MGMT:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_SECURITY_AUTH_ALG:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_PROTO:
		t = ktypeArrayString
	case NM_SETTING_WIRELESS_SECURITY_PAIRWISE:
		t = ktypeArrayString
	case NM_SETTING_WIRELESS_SECURITY_GROUP:
		t = ktypeArrayString
	case NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY0:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY1:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY2:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY3:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_SECURITY_PSK:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS:
		t = ktypeUint32
	}
	return
}

// Check is key in current setting field
func isKeyInSettingWirelessSecurity(key string) bool {
	switch key {
	case NM_SETTING_WIRELESS_SECURITY_KEY_MGMT:
		return true
	case NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX:
		return true
	case NM_SETTING_WIRELESS_SECURITY_AUTH_ALG:
		return true
	case NM_SETTING_WIRELESS_SECURITY_PROTO:
		return true
	case NM_SETTING_WIRELESS_SECURITY_PAIRWISE:
		return true
	case NM_SETTING_WIRELESS_SECURITY_GROUP:
		return true
	case NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME:
		return true
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD:
		return true
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS:
		return true
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY0:
		return true
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY1:
		return true
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY2:
		return true
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY3:
		return true
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS:
		return true
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE:
		return true
	case NM_SETTING_WIRELESS_SECURITY_PSK:
		return true
	case NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS:
		return true
	}
	return false
}

// Get key's default value
func getSettingWirelessSecurityKeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_WIRELESS_SECURITY_KEY_MGMT:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX:
		valueJSON = `0`
	case NM_SETTING_WIRELESS_SECURITY_AUTH_ALG:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_SECURITY_PROTO:
		valueJSON = `null`
	case NM_SETTING_WIRELESS_SECURITY_PAIRWISE:
		valueJSON = `null`
	case NM_SETTING_WIRELESS_SECURITY_GROUP:
		valueJSON = `null`
	case NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS:
		valueJSON = `0`
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY0:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY1:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY2:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY3:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS:
		valueJSON = `0`
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE:
		valueJSON = `0`
	case NM_SETTING_WIRELESS_SECURITY_PSK:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS:
		valueJSON = `0`
	}
	return
}

// Get JSON value generally
func generalGetSettingWirelessSecurityKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingWirelessSecurityKeyJSON: invalide key", key)
	case NM_SETTING_WIRELESS_SECURITY_KEY_MGMT:
		value = getSettingWirelessSecurityKeyMgmtJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX:
		value = getSettingWirelessSecurityWepTxKeyidxJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_AUTH_ALG:
		value = getSettingWirelessSecurityAuthAlgJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_PROTO:
		value = getSettingWirelessSecurityProtoJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_PAIRWISE:
		value = getSettingWirelessSecurityPairwiseJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_GROUP:
		value = getSettingWirelessSecurityGroupJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME:
		value = getSettingWirelessSecurityLeapUsernameJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD:
		value = getSettingWirelessSecurityLeapPasswordJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS:
		value = getSettingWirelessSecurityLeapPasswordFlagsJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY0:
		value = getSettingWirelessSecurityWepKey0JSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY1:
		value = getSettingWirelessSecurityWepKey1JSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY2:
		value = getSettingWirelessSecurityWepKey2JSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY3:
		value = getSettingWirelessSecurityWepKey3JSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS:
		value = getSettingWirelessSecurityWepKeyFlagsJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE:
		value = getSettingWirelessSecurityWepKeyTypeJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_PSK:
		value = getSettingWirelessSecurityPskJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS:
		value = getSettingWirelessSecurityPskFlagsJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingWirelessSecurityKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingWirelessSecurityKeyJSON: invalide key", key)
	case NM_SETTING_WIRELESS_SECURITY_KEY_MGMT:
		err = setSettingWirelessSecurityKeyMgmtJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX:
		err = setSettingWirelessSecurityWepTxKeyidxJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_AUTH_ALG:
		err = setSettingWirelessSecurityAuthAlgJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_PROTO:
		err = setSettingWirelessSecurityProtoJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_PAIRWISE:
		err = setSettingWirelessSecurityPairwiseJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_GROUP:
		err = setSettingWirelessSecurityGroupJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME:
		err = setSettingWirelessSecurityLeapUsernameJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD:
		err = setSettingWirelessSecurityLeapPasswordJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS:
		err = setSettingWirelessSecurityLeapPasswordFlagsJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY0:
		err = setSettingWirelessSecurityWepKey0JSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY1:
		err = setSettingWirelessSecurityWepKey1JSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY2:
		err = setSettingWirelessSecurityWepKey2JSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY3:
		err = setSettingWirelessSecurityWepKey3JSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS:
		err = setSettingWirelessSecurityWepKeyFlagsJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE:
		err = setSettingWirelessSecurityWepKeyTypeJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_PSK:
		err = setSettingWirelessSecurityPskJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS:
		err = setSettingWirelessSecurityPskFlagsJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingWirelessSecurityKeyMgmtExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT)
}
func isSettingWirelessSecurityWepTxKeyidxExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX)
}
func isSettingWirelessSecurityAuthAlgExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_AUTH_ALG)
}
func isSettingWirelessSecurityProtoExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PROTO)
}
func isSettingWirelessSecurityPairwiseExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PAIRWISE)
}
func isSettingWirelessSecurityGroupExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_GROUP)
}
func isSettingWirelessSecurityLeapUsernameExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME)
}
func isSettingWirelessSecurityLeapPasswordExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD)
}
func isSettingWirelessSecurityLeapPasswordFlagsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS)
}
func isSettingWirelessSecurityWepKey0Exists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY0)
}
func isSettingWirelessSecurityWepKey1Exists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY1)
}
func isSettingWirelessSecurityWepKey2Exists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY2)
}
func isSettingWirelessSecurityWepKey3Exists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY3)
}
func isSettingWirelessSecurityWepKeyFlagsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS)
}
func isSettingWirelessSecurityWepKeyTypeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE)
}
func isSettingWirelessSecurityPskExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK)
}
func isSettingWirelessSecurityPskFlagsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS)
}

// Ensure field and key exists and not empty
func ensureFieldSettingWirelessSecurityExists(data connectionData, errs fieldErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_WIRELESS_SECURITY_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME))
	}
}
func ensureSettingWirelessSecurityKeyMgmtNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityKeyMgmtExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSecurityKeyMgmt(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecurityWepTxKeyidxNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityWepTxKeyidxExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWirelessSecurityAuthAlgNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityAuthAlgExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_AUTH_ALG, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSecurityAuthAlg(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_AUTH_ALG, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecurityProtoNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityProtoExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PROTO, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSecurityProto(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PROTO, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecurityPairwiseNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityPairwiseExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PAIRWISE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSecurityPairwise(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PAIRWISE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecurityGroupNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityGroupExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_GROUP, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSecurityGroup(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_GROUP, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecurityLeapUsernameNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityLeapUsernameExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSecurityLeapUsername(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecurityLeapPasswordNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityLeapPasswordExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSecurityLeapPassword(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecurityLeapPasswordFlagsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityLeapPasswordFlagsExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWirelessSecurityWepKey0NoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityWepKey0Exists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY0, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSecurityWepKey0(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY0, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecurityWepKey1NoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityWepKey1Exists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY1, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSecurityWepKey1(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY1, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecurityWepKey2NoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityWepKey2Exists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY2, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSecurityWepKey2(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY2, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecurityWepKey3NoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityWepKey3Exists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY3, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSecurityWepKey3(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY3, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecurityWepKeyFlagsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityWepKeyFlagsExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWirelessSecurityWepKeyTypeNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityWepKeyTypeExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWirelessSecurityPskNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityPskExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSecurityPsk(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecurityPskFlagsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecurityPskFlagsExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingWirelessSecurityKeyMgmt(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT).(string)
	return
}
func getSettingWirelessSecurityWepTxKeyidx(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX).(uint32)
	return
}
func getSettingWirelessSecurityAuthAlg(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_AUTH_ALG).(string)
	return
}
func getSettingWirelessSecurityProto(data connectionData) (value []string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PROTO).([]string)
	return
}
func getSettingWirelessSecurityPairwise(data connectionData) (value []string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PAIRWISE).([]string)
	return
}
func getSettingWirelessSecurityGroup(data connectionData) (value []string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_GROUP).([]string)
	return
}
func getSettingWirelessSecurityLeapUsername(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME).(string)
	return
}
func getSettingWirelessSecurityLeapPassword(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD).(string)
	return
}
func getSettingWirelessSecurityLeapPasswordFlags(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS).(uint32)
	return
}
func getSettingWirelessSecurityWepKey0(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY0).(string)
	return
}
func getSettingWirelessSecurityWepKey1(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY1).(string)
	return
}
func getSettingWirelessSecurityWepKey2(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY2).(string)
	return
}
func getSettingWirelessSecurityWepKey3(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY3).(string)
	return
}
func getSettingWirelessSecurityWepKeyFlags(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS).(uint32)
	return
}
func getSettingWirelessSecurityWepKeyType(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE).(uint32)
	return
}
func getSettingWirelessSecurityPsk(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK).(string)
	return
}
func getSettingWirelessSecurityPskFlags(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS).(uint32)
	return
}

// Setter
func setSettingWirelessSecurityKeyMgmt(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT, value)
}
func setSettingWirelessSecurityWepTxKeyidx(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX, value)
}
func setSettingWirelessSecurityAuthAlg(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_AUTH_ALG, value)
}
func setSettingWirelessSecurityProto(data connectionData, value []string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PROTO, value)
}
func setSettingWirelessSecurityPairwise(data connectionData, value []string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PAIRWISE, value)
}
func setSettingWirelessSecurityGroup(data connectionData, value []string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_GROUP, value)
}
func setSettingWirelessSecurityLeapUsername(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME, value)
}
func setSettingWirelessSecurityLeapPassword(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD, value)
}
func setSettingWirelessSecurityLeapPasswordFlags(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS, value)
}
func setSettingWirelessSecurityWepKey0(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY0, value)
}
func setSettingWirelessSecurityWepKey1(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY1, value)
}
func setSettingWirelessSecurityWepKey2(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY2, value)
}
func setSettingWirelessSecurityWepKey3(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY3, value)
}
func setSettingWirelessSecurityWepKeyFlags(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS, value)
}
func setSettingWirelessSecurityWepKeyType(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE, value)
}
func setSettingWirelessSecurityPsk(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK, value)
}
func setSettingWirelessSecurityPskFlags(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS, value)
}

// JSON Getter
func getSettingWirelessSecurityKeyMgmtJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_KEY_MGMT))
	return
}
func getSettingWirelessSecurityWepTxKeyidxJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX))
	return
}
func getSettingWirelessSecurityAuthAlgJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_AUTH_ALG, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_AUTH_ALG))
	return
}
func getSettingWirelessSecurityProtoJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PROTO, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PROTO))
	return
}
func getSettingWirelessSecurityPairwiseJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PAIRWISE, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PAIRWISE))
	return
}
func getSettingWirelessSecurityGroupJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_GROUP, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_GROUP))
	return
}
func getSettingWirelessSecurityLeapUsernameJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME))
	return
}
func getSettingWirelessSecurityLeapPasswordJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD))
	return
}
func getSettingWirelessSecurityLeapPasswordFlagsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS))
	return
}
func getSettingWirelessSecurityWepKey0JSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY0, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY0))
	return
}
func getSettingWirelessSecurityWepKey1JSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY1, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY1))
	return
}
func getSettingWirelessSecurityWepKey2JSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY2, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY2))
	return
}
func getSettingWirelessSecurityWepKey3JSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY3, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY3))
	return
}
func getSettingWirelessSecurityWepKeyFlagsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS))
	return
}
func getSettingWirelessSecurityWepKeyTypeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE))
	return
}
func getSettingWirelessSecurityPskJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PSK))
	return
}
func getSettingWirelessSecurityPskFlagsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS))
	return
}

// JSON Setter
func setSettingWirelessSecurityKeyMgmtJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_KEY_MGMT))
}
func setSettingWirelessSecurityWepTxKeyidxJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX))
}
func setSettingWirelessSecurityAuthAlgJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_AUTH_ALG, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_AUTH_ALG))
}
func setSettingWirelessSecurityProtoJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PROTO, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PROTO))
}
func setSettingWirelessSecurityPairwiseJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PAIRWISE, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PAIRWISE))
}
func setSettingWirelessSecurityGroupJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_GROUP, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_GROUP))
}
func setSettingWirelessSecurityLeapUsernameJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME))
}
func setSettingWirelessSecurityLeapPasswordJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD))
}
func setSettingWirelessSecurityLeapPasswordFlagsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS))
}
func setSettingWirelessSecurityWepKey0JSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY0, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY0))
}
func setSettingWirelessSecurityWepKey1JSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY1, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY1))
}
func setSettingWirelessSecurityWepKey2JSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY2, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY2))
}
func setSettingWirelessSecurityWepKey3JSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY3, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY3))
}
func setSettingWirelessSecurityWepKeyFlagsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS))
}
func setSettingWirelessSecurityWepKeyTypeJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE))
}
func setSettingWirelessSecurityPskJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PSK))
}
func setSettingWirelessSecurityPskFlagsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS))
}

// Logic JSON Setter

// Remover
func removeSettingWirelessSecurityKeyMgmt(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT)
}
func removeSettingWirelessSecurityWepTxKeyidx(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX)
}
func removeSettingWirelessSecurityAuthAlg(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_AUTH_ALG)
}
func removeSettingWirelessSecurityProto(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PROTO)
}
func removeSettingWirelessSecurityPairwise(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PAIRWISE)
}
func removeSettingWirelessSecurityGroup(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_GROUP)
}
func removeSettingWirelessSecurityLeapUsername(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME)
}
func removeSettingWirelessSecurityLeapPassword(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD)
}
func removeSettingWirelessSecurityLeapPasswordFlags(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS)
}
func removeSettingWirelessSecurityWepKey0(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY0)
}
func removeSettingWirelessSecurityWepKey1(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY1)
}
func removeSettingWirelessSecurityWepKey2(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY2)
}
func removeSettingWirelessSecurityWepKey3(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY3)
}
func removeSettingWirelessSecurityWepKeyFlags(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS)
}
func removeSettingWirelessSecurityWepKeyType(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE)
}
func removeSettingWirelessSecurityPsk(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK)
}
func removeSettingWirelessSecurityPskFlags(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS)
}
