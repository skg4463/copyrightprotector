package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type copyrightprotector struct {
}

type owner struct {
	Name   string           `json:"Name"`
	Videos map[string]video `json:"Videos"`
}

type video struct {
	Id       string `json:"Id"`
	Owner    string `json:"Owner"`
	Metadata string `json:"Metadata"`
}

func (t *copyrightprotector) Init(shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *copyrightprotector) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error

	switch fn {
	case "initLedger":
		result, err = initLedger(stub)
	case "queryOwner":
		result, err = queryOwner(stub, args)
	case "queryVideo":
		result, err = queryVideo(stub, args)
	case "addVideo":
		result, err = addVideo(stub, args)
	case "purchaseVideo":
		result, err = purchaseVideo(stub, args)
	case "anotherCCFunc":
		chainCodeArgs := util.ToChaincodeArgs("anotherCCFunc", "paramA")
		response := stub.InvokeChaincode("anotherCCFunc", chainCodeArgs, "channelname")
		if response.Status != shim.OK {
			return shim.Error(response.Message)
		}
	}
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(result))
}

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

func main() {
	err := shim.Start(new(copyrightprotector))
	if err != nil {
		fmt.Println("Could not start copyrightprotector Chaincode")
	} else {
		fmt.Println("copyrightprotector Chaincode successfully started")
	}
}
