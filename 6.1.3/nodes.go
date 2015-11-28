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
	"encoding/xml"
	"fmt"
)

type UniCloudNodes struct {
	Nodes []struct {
		Bootfrom     int         `json:"bootFrom"`
		Lastupdate   string      `json:"lastUpdate"`
		Name         string      `json:"name"`
		Parentnodeid interface{} `json:"parentNodeId"`
		Myunits      int         `json:"myUnits"`
		Nics         []struct {
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
		State             string `json:"state"`
		Rank              int    `json:"rank"`
		ID                int    `json:"id"`
		Softwareprofileid int    `json:"softwareProfileId"`
		Lockedstate       string `json:"lockedState"`
		Hardwareprofile   struct {
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
			Description                 string        `json:"description"`
			Admins                      []interface{} `json:"admins"`
			Nameformat                  string        `json:"nameFormat"`
			Networks                    []interface{} `json:"networks"`
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
		} `json:"hardwareprofile"`
		Destspid          interface{} `json:"destSPId"`
		Hardwareprofileid int         `json:"hardwareProfileId"`
		Isidle            bool        `json:"isIdle"`
		Addhostsession    interface{} `json:"addHostSession"`
		Maxchildunits     int         `json:"maxChildUnits"`
		Rack              int         `json:"rack"`
		Softwareprofile   struct {
			Kernel       interface{}   `json:"kernel"`
			Description  string        `json:"description"`
			Kernelparams interface{}   `json:"kernelParams"`
			Type         string        `json:"type"`
			Admins       []interface{} `json:"admins"`
			Name         string        `json:"name"`
			Kitsources   []interface{} `json:"kitsources"`
			Initrd       interface{}   `json:"initrd"`
			Minnodes     interface{}   `json:"minNodes"`
			Components   []interface{} `json:"components"`
			Isidle       bool          `json:"isIdle"`
			Osid         int           `json:"osId"`
			Nodes        []interface{} `json:"nodes"`
			Packages     []interface{} `json:"packages"`
			Os           struct {
				Osfamily struct {
					Version string `json:"version"`
					Arch    string `json:"arch"`
					Name    string `json:"name"`
					ID      int    `json:"id"`
				} `json:"osFamily"`
				Version string `json:"version"`
				Arch    string `json:"arch"`
				Name    string `json:"name"`
				ID      int    `json:"id"`
			} `json:"os"`
			ID               int `json:"id"`
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
				Description                 string        `json:"description"`
				Admins                      []interface{} `json:"admins"`
				Nameformat                  string        `json:"nameFormat"`
				Networks                    []interface{} `json:"networks"`
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
			Partitions []interface{} `json:"partitions"`
		} `json:"softwareprofile"`
	} `json:"nodes"`
}

func (ucn UniCloudNodes) PrintJSON() error {
	doc, err := json.Marshal(ucn)
	if err == nil {
		fmt.Println(string(doc))
	}
	return err
}

func (ucn UniCloudNodes) PrintXML() error {
	doc, err := xml.Marshal(ucn)
	if err == nil {
		fmt.Println(string(doc))
	}
	return err
}
