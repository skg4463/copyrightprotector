package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func (t *copyrightprotector) initLedger(stub shim.ChaincodeStubInterface) (string, error) {
	owners := map[string]owner{}
	videos := map[string]video{}

	owners["nakhoon"] = owner{
		Name:   "Nakhoon",
		Videos: map[string]video{},
	}
	owners["owner"] = owner{
		Name:   "Nakhoon",
		Videos: map[string]video{},
	}

	videos["DCPoTnakAe0"] = video{
		Id:       "DCPoTnakAe0",
		Owner:    owners["nakhoon"],
		Metadata: "meta1",
		ContractedInfo: transferContractedInfo{
			Contractor: "",
			Contractee: "",
			ContractInfo: forTransferContractInfo{
				ContractClass: 0,
				ContractFee:   0,
			},
			ParentVideo: "",
		},
	}
	videos["id"] = video{
		Id:       "id",
		Owner:    owners["owner"],
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
