package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

// queryowner(owner name) =
func queryOwner(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	owners := map[string]owner{}
	ownersAsBytes, err := stub.GetState("Owners")
	_ = json.Unmarshal(ownersAsBytes, &owners)

	selectedownerAsBytes, err := json.Marshal(owners[args[0]])

	if err != nil {
		return "", fmt.Errorf("Failed to get video: %s with error: %s", args[0], err)
	}
	if selectedownerAsBytes == nil {
		return "", fmt.Errorf("video not found: %s", args[0])
	}
	fmt.Println(selectedownerAsBytes)
	return string(selectedownerAsBytes), nil
}
