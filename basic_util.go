package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric/common/util"
)

//invoke another chaincode in same channel
//@@add args (chaincode info)
func _(stub shim.ChaincodeStubInterface) (string, error) {
	chainCodeArgs := util.ToChaincodeArgs("anotherCCFunc", "paramA")
	response := stub.InvokeChaincode("anotherCCFunc", chainCodeArgs, "channelname")
	if response.Status != shim.OK {
		return "", fmt.Errorf("error, %s", response.Message)
	}

	return "", nil
}
