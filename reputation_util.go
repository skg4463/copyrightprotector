package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

//reputation raw (float64) -> persentage (int)
//param : owner id
//reputationRawCalc
func _(stub shim.ChaincodeStubInterface, args []string) (int, error) {
	if len(args) != 1 {
		return 0, fmt.Errorf("Incorrect arguments, Expecting 1 arguments")
	}
	ownerAsBytes, _ := stub.GetState("Owners")
	ownerData := map[string]owner{}
	_ = json.Unmarshal(ownerAsBytes, &ownerData)
	rr := ownerData[args[0]].ReputationRaw

	return int(rr), fmt.Errorf("reputationRawCalc")
}
