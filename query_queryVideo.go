package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func (t *copyrightprotector) queryVideo(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting an id")
	}

	videos := map[string]video{}
	videosAsBytes, err := stub.GetState("Videos")
	_ = json.Unmarshal(videosAsBytes, &videos)
	selectedvideoAsBytes, err := json.Marshal(videos[args[0]])

	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if selectedvideoAsBytes == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(selectedvideoAsBytes), nil
}
