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
	"strconv"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage:")
		fmt.Println("addNodes <password> <installer-adddress> <amount> <hardware-profile> <software-profile>")
		os.Exit(1)
	}

	amount, errConv := strconv.Atoi(os.Args[3])
	if errConv != nil {
		fmt.Println("Error: Third argument is not a number.")
		os.Exit(1)
	}

	// password and hostaddress as arguments
	ctx := unicloud.NewContext("cfm", os.Args[1], os.Args[2])

	// adding 2 new nodes with hardware profile EC2 and software profile Compute
	if result, err := ctx.AddNode(amount, os.Args[4], os.Args[5]); err != nil {
		fmt.Printf("Error when adding %d additional nodes: %s\n", amount, err)
		os.Exit(1)
	} else {
		fmt.Printf("Successfully added %d nodes.\n", amount)
		fmt.Printf("Add Host Session: %s\n", result.AddHostSession)
	}
}
