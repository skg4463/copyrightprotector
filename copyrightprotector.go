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
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	chainCodeArgs := util.ToChaincodeArgs("anotherCCFunc", "paramA")
	response := stub.InvokeChaincode("anotherCCFunc", chainCodeArgs, "channelname")
	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	return shim.Success([]byte(result))
}

func main() {
	err := shim.Start(new(copyrightprotector))
	if err != nil {
		fmt.Println("Could not start 'copyrightprotector' Chaincode")
	} else {
		fmt.Println("copyrightprotector Chaincode successfully started")
	}
}
