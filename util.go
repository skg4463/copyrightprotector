package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
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
