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

package main

import (
	"fmt"
	"github.com/dgruber/gounicloud/6.1.3"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please call it with password and installer address as arguments.")
		os.Exit(1)
	}

	// password and hostaddress as arguments
	ctx := unicloud.NewContext("cfm", os.Args[1], os.Args[2])

	if sprofiles, err := ctx.GetSoftwareProfiles(); err != nil {
		fmt.Printf("Error when calling GetSoftwareProfiles(): %s\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Following software profiles found:")
		for _, sp := range sprofiles.Softwareprofiles {
			fmt.Printf("%s\n", sp.Name)
		}
	}
}
