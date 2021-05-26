package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"strconv"
)

//@Param targetVideo serial
//@Emit contractAlert
func transferContractPresent(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments, Expecting 1 arguments")
	}

	videoAsBytes, _ := stub.GetState("Videos")
	videos := map[string]video{}
	_ = json.Unmarshal(videoAsBytes, &videos)

	videoInfo := videos[args[0]]

	if videoInfo.ForTransferContractInfo.ContractFee != 0 {
		return "", fmt.Errorf("Incorrect Contract Call, this Video is not Creative Commons")
	}

	transactionCreator, _ := getCreatorCert(stub)

	alertStruct := contractAlert{
		Contractor: transactionCreator,
		Contractee: videoInfo.Owner.Identity, //indexed
		Video:      videoInfo.Id,
	}
	alert, _ := json.Marshal(alertStruct)
	err := stub.SetEvent("transferContractAlert", alert)
	if err != nil {
		return "", err
	}

	contractAsBytes, _ := stub.GetState("Contracts")
	contracts := map[string]transferContractWaitingList{}
	_ = json.Unmarshal(contractAsBytes, &contracts)

	contracts[strconv.Itoa(contractCount)] = transferContractWaitingList{
		Contractor: transactionCreator,
		Contractee: videoInfo.Owner.Identity,
		Video:      videoInfo.Id,
		Isfine:     false,
	}
	contractCount++

	updatedContracts, _ := json.Marshal(contracts)
	err = stub.PutState("Contracts", updatedContracts)
	if err != nil {
		return "", fmt.Errorf("Failed to set : %s", args[0])
	}

	return string(updatedContracts), nil
}
