/**
 * Copyright (c) 2014 Deepin, Inc.
 *               2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

// "NM_SETTING_CONNECTION_SETTING_NAME" -> "../nm_setting_connection_gen.go"
func getBackEndFilePath(sectionName string) (filePath string) {
	sectionName = strings.Replace(sectionName, "NM_SETTING_ALIAS_", "NM_SETTING_", -1) // remove virtual section tag
	fileName := strings.TrimSuffix(sectionName, "_SETTING_NAME")
	fileName = strings.ToLower(fileName) + "_gen.go"
	filePath = path.Join(argBackEndGoDir, fileName)
	return
}

// "NM_SETTING_VS_GENERAL" -> "../../../dss/modules/network/edit_autogen/EditSectionGeneral.qml"
func getFrontEndFilePath(vsectionName string) (filePath string) {
	fileName := "EditSection" + ToVsClassName(vsectionName) + ".qml"
	filePath = path.Join(argFrontEndQmlDir, fileName)
	return
}

func writeOrDisplayResultForBackEnd(file, content string) {
	if argWriteOutput {
		writeBackEndFile(file, content)
	} else {
		fmt.Println(content)
		fmt.Println()
	}
}

func writeOrDisplayResultForFrontEnd(file, content string) {
	if argWriteOutput {
		writeFrontEndFile(file, content)
	} else {
		fmt.Println(content)
		fmt.Println()
	}
}

func writeBackEndFile(file, content string) {
	// write to .go file and execute gofmt
	err := ioutil.WriteFile(file, []byte(content), 0644)
	if err != nil {
		fmt.Println("error, write file failed:", err)
		return
	}
	execAndWait(10, "gofmt", "-w", file)
	fmt.Println("GEN " + file)
}

func writeFrontEndFile(file, content string) {
	// write to .go file and execute gofmt
	err := ioutil.WriteFile(file, []byte(content), 0644)
	if err != nil {
		fmt.Println("error, write file failed:", err)
		return
	}
	fmt.Println("GEN " + file)
}

func GetAllKeysInVsection(vsectionName string) (keys []string) {
	// get virtual keys which related section is this virtual section
	// first, and they should be controller key, such as "vk-vpn-type"
	for _, vkey := range nmVkeys {
		if vkey.RelatedSection == vsectionName {
			keys = appendStrArrayUnique(keys, vkey.Name)
		}
	}
	vsectionInfo := getVsectionInfo(vsectionName)
	for _, section := range vsectionInfo.RelatedSections {
		keys = appendStrArrayUnique(keys, GetAllKeysInSection(section)...)
	}
	return
}

// GetAllKeysInSection return all keys that will be used by front-end.
func GetAllKeysInSection(sectionName string) (keys []string) {
	for _, nmSection := range nmSections {
		if nmSection.Name == sectionName {
			for _, k := range nmSection.Keys {
				vksNames := getRelatedVks(k.Name)
				if len(vksNames) > 0 {
					// if virtual key is a enable wrapper, both
					// virtual key and real key will be appended.
					for _, vkName := range vksNames {
						keys = appendStrArrayUnique(keys, vkName)
						if IsEnableWrapperVkey(vkName) {
							keys = appendStrArrayUnique(keys, k.Name)
						}
					}
				} else {
					keys = append(keys, k.Name)
				}
			}
			break
		}
	}
	return
}

// get all related virtual keys of real key
func getRelatedVks(keyName string) (vks []string) {
	for _, vk := range nmVkeys {
		if isStringInArray(keyName, vk.RelatedKeys) {
			vks = append(vks, vk.Name)
		}
	}
	return
}

func ToKeyDisplayName(keyName string) (displayName string) {
	var keyValue string
	if isVk(keyName) {
		vkInfo := getVkInfo(keyName)
		displayName = vkInfo.DisplayName
		keyValue = vkInfo.Value
		if displayName == "<default>" {
			keyInfo := getKeyInfo(vkInfo.RelatedKeys[0])
			displayName = keyInfo.DisplayName
		}
	} else {
		keyInfo := getKeyInfo(keyName)
		displayName = keyInfo.DisplayName
		keyValue = keyInfo.Value
	}
	if displayName != "<default>" {
		return
	}
	return "!!" + keyValue
}

// "NM_SETTING_802_1X_EAP" -> "eap"
func ToKeyValue(keyName string) (keyValue string) {
	if isVk(keyName) {
		keyValue = getVkInfo(keyName).Value
	} else {
		keyValue = getKeyInfo(keyName).Value
	}
	return
}

func ToKeyAlwaysUpdate(keyName string) (alwaysUpdate bool) {
	if isVk(keyName) {
		alwaysUpdate = getVkInfo(keyName).AlwaysUpdate
	} else {
		alwaysUpdate = getKeyInfo(keyName).AlwaysUpdate
	}
	return
}

func ToKeyUseValueRange(keyName string) (useValueRange bool) {
	if isVk(keyName) {
		useValueRange = getVkInfo(keyName).UseValueRange
	} else {
		useValueRange = getKeyInfo(keyName).UseValueRange
	}
	return
}

func ToKeyMinValue(keyName string) (minValue int) {
	if isVk(keyName) {
		minValue = getVkInfo(keyName).MinValue
	} else {
		minValue = getKeyInfo(keyName).MinValue
	}
	return
}

func ToKeyMaxValue(keyName string) (maxValue int) {
	if isVk(keyName) {
		maxValue = getVkInfo(keyName).MaxValue
	} else {
		maxValue = getKeyInfo(keyName).MaxValue
	}
	return
}

// "NM_SETTING_802_1X_EAP" -> "802-1x", "NM_SETTING_VS_GENERAL" -> "<none>"
func ToKeyRelatedSectionValue(keyName string) (sectionValue string) {
	if isVk(keyName) {
		sectionName := getVkInfo(keyName).RelatedSection
		sectionValue = ToSectionValue(sectionName)
	} else {
		sectionValue = getKeyRelatedSectionInfo(keyName).Value
	}
	return
}
func getKeyRelatedSectionInfo(keyName string) (sectionInfo NMSectionStruct) {
	for _, section := range nmSections {
		for _, key := range section.Keys {
			if key.Name == keyName {
				sectionInfo = section
				return
			}
		}
	}
	fmt.Println("invalid key name", keyName)
	os.Exit(1)
	return
}

// "NM_SETTING_802_1X_SETTING_NAME" -> "802-1x", "NM_SETTING_VS_GENERAL" -> "vs-general"
func ToSectionValue(sectionName string) (sectionValue string) {
	if isVirtualSection(sectionName) {
		sectionValue = getVsectionInfo(sectionName).Value
	} else {
		sectionValue = getSectionInfo(sectionName).Value
	}
	return
}

func GetKeyWidgetProps(keyName string) (prop map[string]string) {
	if isVk(keyName) {
		prop = getVkInfo(keyName).WidgetProps
	} else {
		prop = getKeyInfo(keyName).WidgetProps
	}
	return
}

func IsKeyUsedByFrontEnd(keyName string) (used bool) {
	if isVk(keyName) {
		used = getVkInfo(keyName).UsedByFrontEnd
	} else {
		used = getKeyInfo(keyName).UsedByFrontEnd
	}
	return
}

func isVirtualSection(sectionName string) bool {
	for _, vsection := range nmVsections {
		if vsection.Name == sectionName {
			return true
		}
	}
	return false
}

func getSectionInfo(sectionName string) (sectionInfo NMSectionStruct) {
	for _, section := range nmSections {
		if section.Name == sectionName {
			sectionInfo = section
			return
		}
	}
	fmt.Println("invalid section name", sectionName)
	os.Exit(1)
	return
}

func getVsectionInfo(vsectionName string) (vsectionInfo NMVsectionStruct) {
	for _, vsection := range nmVsections {
		if vsection.Name == vsectionName {
			vsectionInfo = vsection
			return
		}
	}
	fmt.Println("invalid vsection name", vsectionName)
	os.Exit(1)
	return
}

func getKeyInfo(keyName string) (keyInfo NMKeyStruct) {
	for _, section := range nmSections {
		for _, key := range section.Keys {
			if key.Name == keyName {
				keyInfo = key
				return
			}
		}
	}
	fmt.Println("invalid key name", keyName)
	os.Exit(1)
	return
}

func getVkInfo(vkName string) (vkInfo NMVkeyStruct) {
	for _, vk := range nmVkeys {
		if vk.Name == vkName {
			vkInfo = vk
			return
		}
	}
	fmt.Println("invalid key name", vkName)
	os.Exit(1)
	return
}

// check if target key is a virtual key
func isVk(keyName string) (ok bool) {
	for _, vk := range nmVkeys {
		if vk.Name == keyName {
			return true
		}
	}
	return false
}

func isWrapperVkey(vkName string) (ok bool) {
	if getVkInfo(vkName).VkType == "vkTypeWrapper" {
		return true
	}
	return false
}

func IsEnableWrapperVkey(vkName string) (ok bool) {
	if getVkInfo(vkName).VkType == "vkTypeEnableWrapper" {
		return true
	}
	return false
}

func isControllerVkey(vkName string) (ok bool) {
	if getVkInfo(vkName).VkType == "vkTypeController" {
		return true
	}
	return false
}

// NM_SETTING_CONNECTION_SETTING_NAME -> ConnectionSetting, NM_SETTING_VK_VPN_L2TP_SETTING_NAME -> VpnL2tp
func ToSectionFuncBaseName(name string) (funcName string) {
	name = strings.Replace(name, "NM_SETTING_ALIAS_", "NM_SETTING_", -1) // remove virtual section tag
	funcName = strings.TrimPrefix(name, "NM_")
	funcName = strings.TrimSuffix(funcName, "_SETTING_NAME")
	funcName = strings.Replace(funcName, "_", " ", -1)
	funcName = ToCaplitalize(funcName)
	funcName = strings.Replace(funcName, " ", "", -1)
	return
}

// NM_SETTING_CONNECTION_ID -> SettingConnectionId
func ToKeyFuncBaseName(name string) (funcName string) {
	funcName = strings.TrimPrefix(name, "NM_")
	funcName = strings.Replace(funcName, "_", " ", -1)
	funcName = ToCaplitalize(funcName)
	funcName = strings.Replace(funcName, " ", "", -1)
	return
}

// "hello world" -> "Hello World", "HELLO WORLD" -> "Hello World"
func ToCaplitalize(str string) (capstr string) {
	scaner := bufio.NewScanner(strings.NewReader(str))
	scaner.Split(bufio.ScanWords)
	for scaner.Scan() {
		word := scaner.Text()
		if len(word) > 1 {
			capstr = capstr + " " + strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		} else if len(word) == 1 {
			capstr = capstr + " " + strings.ToUpper(word)
		}
	}
	capstr = strings.TrimSpace(capstr)
	return
}

// "VPN_L2TP" -> "VpnL2tp", "vk-key-mgmt" -> "VkKeyMgmt", "vk key mgmt" -> "VkKeyMgmt", "vk_key_mgmt" -> "VkKeyMgmt"
func ToClassName(str string) (className string) {
	className = strings.Replace(str, "_", " ", -1)
	className = strings.Replace(className, "-", " ", -1)
	className = ToCaplitalize(className)
	className = strings.Replace(className, " ", "", -1)
	return
}

// "NM_SETTING_VS_GENERAL" -> "General"
func ToVsClassName(vsectionName string) (className string) {
	vsectionName = strings.TrimPrefix(vsectionName, "NM_SETTING_VS_")
	className = ToClassName(vsectionName)
	return
}

func unmarshalJSONFile(jsonFile string, data interface{}) {
	fileContent, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("error, open file failed:", err)
		os.Exit(1)
	}

	err = json.Unmarshal(fileContent, data)
	if err != nil {
		fmt.Printf("error, unmarshal json %s failed: %v\n", jsonFile, err)
		os.Exit(1)
	}
}

func execAndWait(timeout int, name string, arg ...string) (stdout, stderr string, err error) {
	cmd := exec.Command(name, arg...)
	var bufStdout, bufStderr bytes.Buffer
	cmd.Stdout = &bufStdout
	cmd.Stderr = &bufStderr
	err = cmd.Start()
	if err != nil {
		return
	}

	// wait for process finished
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		if err = cmd.Process.Kill(); err != nil {
			return
		}
		<-done
		err = fmt.Errorf("time out and process was killed")
	case err = <-done:
		stdout = bufStdout.String()
		stderr = bufStderr.String()
		if err != nil {
			return
		}
	}
	return
}

func isStringInArray(s string, list []string) bool {
	for _, i := range list {
		if i == s {
			return true
		}
	}
	return false
}

func appendStrArrayUnique(a1 []string, a2 ...string) (a []string) {
	a = a1
	for _, s := range a2 {
		if !isStringInArray(s, a) {
			a = append(a, s)
		}
	}
	return
}
