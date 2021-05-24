package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

//@Param targetVideo serial
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

	transactionCreator, _ := stub.GetCreator()

	alert := contractAlert{
		Contractor: transactionCreator,
		Contractee: videoInfo.Owner.Identity,
		Video:      videoInfo.Id,
	}
	alert2, _ := json.Marshal(alert)
	err := stub.SetEvent("transferContractAlert", alert2)
	if err != nil {
		return "", err
	}
}
