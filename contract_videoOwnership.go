package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func (t *copyrightprotector) videoOwnership(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting two arguments (user name, video id)")
	}
	videosAsBytes, _ := stub.GetState("Videos")
	ownersAsBytes, _ := stub.GetState("Owners")

	videos := map[string]video{}
	owners := map[string]owner{}

	_ = json.Unmarshal(videosAsBytes, &videos)
	_ = json.Unmarshal(ownersAsBytes, &owners)

	videoowner := owners[args[0]]
	selectedvideo := videos[args[1]]

	videoowner.Videos[selectedvideo.Id] = selectedvideo

	owners[args[0]] = videoowner

	updatedownerAsBytes, _ := json.Marshal(owners)
	err := stub.PutState("Owners", updatedownerAsBytes)

	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return string(updatedownerAsBytes), nil
}
