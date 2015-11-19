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
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	User     string
	Password string
	Address  string
}

// password is in /etc/cfm/.cfmsecret user is cfm

func NewContext(user, password, address string) Context {
	var ctx Context
	ctx.User = user
	ctx.Password = password
	ctx.Address = address
	return ctx
}

type AddNodeInput struct {
	Node struct {
		Count           int    `json:"count"`
		HardwareProfile string `json:"hardwareProfile"`
		SoftwareProfile string `json:"softwareProfile"`
	} `json:"node"`
}

type inner struct {
	Count           int    `json:"count"`
	HardwareProfile string `json:"hardwareProfile"`
	SoftwareProfile string `json:"softwareProfile"`
}

type AddNodeResult struct {
	AddHostSession string `json:"addHostSession"`
}

// AddNode adds a new UniCloud nodes to the cluster by the given hardware and software
// profiles.
func (ctx Context) AddNode(count int, hardwareprofile, softwareprofile string) (AddNodeResult, error) {
	var result AddNodeResult

	// we trust the signature
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	addInput := AddNodeInput{inner{count, hardwareprofile, softwareprofile}}
	input, errMarshal := json.Marshal(addInput)
	if errMarshal != nil {
		return result, errMarshal
	}
	url := fmt.Sprintf("https://%s:8443/v1/nodes", ctx.Address)
	req, errRequest := http.NewRequest("POST", url, bytes.NewBuffer(input))
	if errRequest != nil {
		return result, errRequest
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(ctx.User, ctx.Password)

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}

type DeleteResult struct {
	DeleteNodeFailed          []interface{} `json:"DeleteNodeFailed"`
	NodesDeleted              []string      `json:"NodesDeleted"`
	SoftwareProfileHardLocked []interface{} `json:"SoftwareProfileHardLocked"`
	SoftwareProfileLocked     []interface{} `json:"SoftwareProfileLocked"`
}

// DeleteNode deletes a UniCloud node by name
func (ctx Context) DeleteNode(name string) (DeleteResult, error) {
	var result DeleteResult

	// we trust the signature
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	url := fmt.Sprintf("https://%s:8443/v1/nodes/%s/True", ctx.Address, name)
	req, errRequest := http.NewRequest("DELETE", url, nil)
	if errRequest != nil {
		return result, errRequest
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(ctx.User, ctx.Password)
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}

// GetNode returns the details of a specific UniCloud node.
func (ctx Context) GetNode(name string) (UniCloudNode, error) {
	var node UniCloudNode

	// we trust the signature
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	url := fmt.Sprintf("https://%s:8443/v1/nodes/%s", ctx.Address, name)
	req, errRequest := http.NewRequest("GET", url, nil)
	if errRequest != nil {
		return node, errRequest
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(ctx.User, ctx.Password)
	resp, err := client.Do(req)
	if err != nil {
		return node, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&node); err != nil {
		return node, err
	}
	return node, nil
}

// GetAllNodes returns all UniCloud nodes available.
func (ctx Context) GetAllNodes() (UniCloudNodes, error) {
	var nodes UniCloudNodes

	// we trust the signature
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	url := fmt.Sprintf("https://%s:8443/v1/nodes", ctx.Address)
	req, errRequest := http.NewRequest("GET", url, nil)
	if errRequest != nil {
		return nodes, errRequest
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(ctx.User, ctx.Password)
	resp, err := client.Do(req)
	if err != nil {
		return nodes, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&nodes); err != nil {
		return nodes, err
	}
	return nodes, nil
}
