package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func initLedger(stub shim.ChaincodeStubInterface) (string, error) {
	owners := map[string]owner{}
	videos := map[string]video{}

	owners["nakhoon"] = owner{
		Name:   "Nakhoon",
		Videos: map[string]video{},
	}
	videos["0"] = video{
		Id:       "DCPoTnakAe0",
		Owner:    "nakhoon",
		Metadata: "meta1",
	}
	videos["1"] = video{
		Id:       "id",
		Owner:    "owner",
		Metadata: "metadata",
	}

	ownersAsBytes, _ := json.Marshal(owners)
	videosAsBytes, _ := json.Marshal(videos)
	err := stub.PutState("Owners", ownersAsBytes)
	er := stub.PutState("Videos", videosAsBytes)
	if (err != nil) && (er != nil) {
		return "", fmt.Errorf("failed to intialize ledger")
	}
	return string(videosAsBytes), err
}
