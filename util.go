package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/core/ledger"
)

//qscc
func getTransactionByID(vledger ledger.PeerLedger, txid []byte) peer.Response {
	if txid == nil {
		return shim.Error("TXID is nil")
	}

	processedTran, err := vledger.GetTransactionByID(string(txid))

	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get transaction with TXID %s, error %s", string(txid), err))
	}

	bytes, err := json.Marshal(processedTran)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(bytes)
}

//param : owner id
func reputationRawCalc(stub shim.ChaincodeStubInterface, args []string) (int, error) {
	if len(args) != 1 {
		return 0, fmt.Errorf("Incorrect arguments, Expecting 1 arguments")
	}
	ownerAsBytes, _ := stub.GetState("Owners")
	ownerData := map[string]owner{}
	_ = json.Unmarshal(ownerAsBytes, &ownerData)
	rr := ownerData[args[0]].ReputationRaw

	return int(rr), fmt.Errorf("reputationRawCalc")
}

//owner authorization
func getCreatorCert(stub shim.ChaincodeStubInterface) (interface{}, error) {
	serializedid, _ := stub.GetCreator()

	sid := &msp.SerializedIdentity{}
	err := json.Unmarshal(serializedid, &sid)
	if err != nil {
		return "", fmt.Errorf("Unmarshal error")
	}

	bl, _ := pem.Decode(sid.IdBytes)
	if bl == nil {
		return "", fmt.Errorf("bl is nil")
	}

	cert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		return "", fmt.Errorf("Unable to parse certificate")
	}

	//certification
	return cert, err
}
