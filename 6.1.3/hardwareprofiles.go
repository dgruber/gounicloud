/*
   Copyright 2015 Daniel Gruber, Univa

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package unicloud

import (
	"encoding/json"
	"fmt"
)

type HardwareProfiles struct {
	Hardwareprofiles []struct {
		Kernel                  interface{} `json:"kernel"`
		Softwareoverrideallowed bool        `json:"softwareOverrideAllowed"`
		Maxunits                int         `json:"maxUnits"`
		Name                    string      `json:"name"`
		Bcastenabled            bool        `json:"bcastEnabled"`
		Localbootparams         interface{} `json:"localBootParams"`
		Nics                    []struct {
			Networkid int `json:"networkId"`
			Network   struct {
				Suffix    interface{} `json:"suffix"`
				Startip   interface{} `json:"startIp"`
				ID        int         `json:"id"`
				Device    interface{} `json:"device"`
				Netmask   string      `json:"netmask"`
				Gateway   interface{} `json:"gateway"`
				Increment int         `json:"increment"`
				Address   string      `json:"address"`
				Usingdhcp bool        `json:"usingDhcp"`
				Type      string      `json:"type"`
				Options   interface{} `json:"options"`
				Name      string      `json:"name"`
			} `json:"network"`
			IP            string      `json:"ip"`
			Boot          bool        `json:"boot"`
			Nodeid        int         `json:"nodeId"`
			Mac           interface{} `json:"mac"`
			Networkdevice struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"networkdevice"`
			ID int `json:"id"`
		} `json:"nics"`
		Description string        `json:"description"`
		Admins      []interface{} `json:"admins"`
		Nameformat  string        `json:"nameFormat"`
		Networks    []struct {
			Suffix        interface{} `json:"suffix"`
			Startip       interface{} `json:"startIp"`
			ID            int         `json:"id"`
			Device        interface{} `json:"device"`
			Netmask       string      `json:"netmask"`
			Gateway       interface{} `json:"gateway"`
			Networkdevice struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"networkdevice"`
			Increment int         `json:"increment"`
			Address   string      `json:"address"`
			Usingdhcp bool        `json:"usingDhcp"`
			Type      string      `json:"type"`
			Options   interface{} `json:"options"`
			Name      string      `json:"name"`
		} `json:"networks"`
		Installtype                 string        `json:"installType"`
		Cost                        int           `json:"cost"`
		Location                    string        `json:"location"`
		Idlesoftwareprofileid       interface{}   `json:"idleSoftwareProfileId"`
		ID                          int           `json:"id"`
		Nodes                       []interface{} `json:"nodes"`
		Mcastenabled                bool          `json:"mcastEnabled"`
		Kernelparams                interface{}   `json:"kernelParams"`
		Initrd                      interface{}   `json:"initrd"`
		Hypervisorsoftwareprofileid interface{}   `json:"hypervisorSoftwareProfileId"`
	} `json:"hardwareprofiles"`
}

func (hp HardwareProfiles) PrintXML() error {
	doc, err := json.Marshal(hp)
	if err == nil {
		fmt.Println(string(doc))
	}
	return err
}
