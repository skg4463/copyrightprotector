package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func (t *copyrightprotector) addVideo(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 3 {
		return "", fmt.Errorf("Incorrect arguments. Expecting 3 arguments (id, owner, metadata)")
	}

	owners := map[string]owner{}
	videos := map[string]video{}
	ownersAsBytes, _ := json.Marshal(owners)
	videossAsBytes, _ := stub.GetState("Videos")
	_ = json.Unmarshal(ownersAsBytes, &owners)
	_ = json.Unmarshal(videossAsBytes, &videos)

	videos[args[0]] = video{
		Id:       args[0],
		Owner:    owners[args[1]],
		Metadata: args[2],
	}

	updatedvideoAsBytes, _ := json.Marshal(videos)
	err := stub.PutState("Videos", updatedvideoAsBytes)
	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return string(updatedvideoAsBytes), nil
}
