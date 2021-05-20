package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func addVideo(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 3 {
		return "", fmt.Errorf("Incorrect arguments. Expecting 3 arguments (id, owner, metadata)")
	}
	videos := map[string]video{}
	videossAsBytes, _ := stub.GetState("Videos")
	_ = json.Unmarshal(videossAsBytes, &videos)

	videos[args[0]] = video{
		Id:       args[0],
		Owner:    args[1],
		Metadata: args[2],
	}

	updatedvideoAsBytes, _ := json.Marshal(videos)
	err := stub.PutState("Videos", updatedvideoAsBytes)
	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return string(updatedvideoAsBytes), nil
}
