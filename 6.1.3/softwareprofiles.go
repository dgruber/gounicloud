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

type SoftwareProfiles struct {
	Softwareprofiles []struct {
		Kernel       string        `json:"kernel"`
		Description  string        `json:"description"`
		Kernelparams interface{}   `json:"kernelParams"`
		Type         string        `json:"type"`
		Admins       []interface{} `json:"admins"`
		Name         string        `json:"name"`
		Kitsources   []interface{} `json:"kitsources"`
		Initrd       string        `json:"initrd"`
		Minnodes     interface{}   `json:"minNodes"`
		Components   []struct {
			Integrationmodulepath interface{}   `json:"integrationModulePath"`
			Packagelist           []interface{} `json:"packageList"`
			Description           string        `json:"description"`
			Osfamilycomponentlist []interface{} `json:"osFamilyComponentList"`
			Oscomponentlist       []interface{} `json:"osComponentList"`
			Version               string        `json:"version"`
			ID                    int           `json:"id"`
			Name                  string        `json:"name"`
		} `json:"components"`
		Isidle   bool          `json:"isIdle"`
		Osid     int           `json:"osId"`
		Nodes    []interface{} `json:"nodes"`
		Packages []interface{} `json:"packages"`
		Os       struct {
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
			Kernel                      string        `json:"kernel"`
			Softwareoverrideallowed     bool          `json:"softwareOverrideAllowed"`
			Maxunits                    int           `json:"maxUnits"`
			Name                        string        `json:"name"`
			Bcastenabled                bool          `json:"bcastEnabled"`
			Localbootparams             string        `json:"localBootParams"`
			Nics                        []interface{} `json:"nics"`
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
			Initrd                      string        `json:"initrd"`
			Hypervisorsoftwareprofileid interface{}   `json:"hypervisorSoftwareProfileId"`
		} `json:"hardwareprofiles"`
		Partitions []struct {
			Softwareprofileid  int         `json:"softwareProfileId"`
			Preserve           bool        `json:"preserve"`
			Indirectattachment string      `json:"indirectAttachment"`
			Name               string      `json:"name"`
			ID                 int         `json:"id"`
			Size               int         `json:"size"`
			Directattachment   string      `json:"directAttachment"`
			Fstype             string      `json:"fsType"`
			Maxsize            interface{} `json:"maxSize"`
			Disksize           int         `json:"diskSize"`
			Device             string      `json:"device"`
			Mountpoint         string      `json:"mountPoint"`
			Grow               bool        `json:"grow"`
			Bootloader         bool        `json:"bootLoader"`
			Options            interface{} `json:"options"`
			Sanvolume          interface{} `json:"sanVolume"`
		} `json:"partitions"`
	} `json:"softwareprofiles"`
}

func (sp SoftwareProfiles) PrintXML() error {
	doc, err := xml.Marshal(sp)
	if err == nil {
		fmt.Println(string(doc))
	}
	return err
}

func (sp SoftwareProfiles) PrintJSON() error {
	doc, err := json.Marshal(sp)
	if err == nil {
		fmt.Println(string(doc))
	}
	return err
}
