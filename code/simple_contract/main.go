/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"contracts"
	"fmt"

	"github.com/awjh-ibm/fabric-go-developer-api/contractapi"
)

func main() {
	simpleContract := new(contracts.Simple)

	cc := contractapi.CreateNewChaincode(simpleContract)

	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting LettersOfCredit chaincode: %s", err)
	}
}
