package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type copyrightprotector struct {
}

func (t *copyrightprotector) Init(shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *copyrightprotector) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	var result string
	var err error

	switch function {
	//init func
	case "initLedger":
		result, err = initLedger(stub)
	//query func
	case "queryOwner":
		result, err = queryOwner(stub, args)
	case "queryVideo":
		result, err = queryVideo(stub, args)
	//adding func
	case "addVideo":
		result, err = addVideo(stub, args)
	//contract func
	case "transferContractPresent":
		result, err = transferContractPresent(stub, args)
	case "videoOwnership":
		result, err = videoOwnership(stub, args)
	//vointg func
	case "createParty":
		result, err = createParty(stub, args)
	case "createVotesAndAssignToAll":
		result, err = createVotesAndAssignToAll(stub, args)
	case "updateParty":
		result, err = updateParty(stub, args)
	}

	if err != nil {
		return shim.Error(err.Error())
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
