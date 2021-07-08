package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric/common/util"
)

//invoke another chaincode in same channel
//@@add args (chaincode info)
func (t *copyrightprotector) _(stub shim.ChaincodeStubInterface) (string, error) {
	chainCodeArgs := util.ToChaincodeArgs("anotherCCFunc", "paramA")
	response := stub.InvokeChaincode("anotherCCFunc", chainCodeArgs, "channelname")
	if response.Status != shim.OK {
		return "", fmt.Errorf("error, %s", response.Message)
	}

	return "", nil
}

func (dcc *DecodedChainCode) getDataArrayStrings(stub shim.ChaincodeStubInterface, dataKey string, args []string) ([]string, error) {
	var err error
	var empty []string
	if len(args) != 0 {
		err = errors.New("{\"Error\":\"Incorrect number of arguments\", \"Function\":\"getDataArrayStrings\"}")
		fmt.Printf("\t *** %s", err)
		return empty, err
	}
	arrayBytes, err := stub.GetState(dataKey)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return empty, err
	}
	var outputArray []string
	err = json.Unmarshal(arrayBytes, &outputArray)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return empty, err
	}
	return outputArray, nil
}

func (dcc *DecodedChainCode) saveStringToDataArray(stub shim.ChaincodeStubInterface, dataKey string, addString string, ledger []string) ([]byte, error) {
	var err error
	// Add the string to the array
	ledger = append(ledger, addString)
	// err = dcc.saveLedger(stub, dataKey, ledger)
	if err = dcc.saveLedger(stub, dataKey, ledger); err != nil {
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	return nil, nil
}

func (dcc *DecodedChainCode) saveLedger(stub shim.ChaincodeStubInterface, dataKey string, ledger []string) error {
	var err error
	// Marshall the ledger to bytes
	bytesToWrite, err := json.Marshal(&ledger)
	if err != nil {
		return err
	}
	// Save the array.
	err = stub.PutState(dataKey, bytesToWrite)
	if err != nil {
		return err
	}
	return nil
}
