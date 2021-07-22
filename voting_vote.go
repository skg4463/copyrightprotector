package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/nu7hatch/gouuid"
	"strconv"
)

func createVotesAndAssignToAll(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	var err error
	if len(args) != 0 {
		err = errors.New("{\"Error\":\"Expecting 0 arguments, got " + strconv.Itoa(len(args)))
		fmt.Printf("\t *** %s", err)
		return "", err
	}

	var emptyArgs []string

	// Get all parties
	partiesLedgerBytes, err := readAllParties(stub, emptyArgs)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return "", err
	}

	var partiesLedger []Party

	if err := json.Unmarshal(partiesLedgerBytes, &partiesLedger); err != nil {
		fmt.Printf("\t *** %s", err)
		return "", err
	}

	// Iterate over parties
	for _, party := range partiesLedger {

		// Filter for party.Voter == true
		// party == voter
		if party.Voter {
			// create a new vote with uuid
			u4, err := uuid.NewV4()
			if err != nil {
				fmt.Printf("\t *** %s", err)
				return "", err
			}
			vote := Vote{
				Uuid: u4.String(),
			}
			// save new vote in blockchain
			if err = vote.save(stub); err != nil {
				fmt.Printf("\t *** %s", err)
				return "", err
			}
			// assign new vote to voting party
			args := []string{party.Id, vote.Uuid, "", ""}
			if _, err := updateParty(stub, args); err != nil {
				fmt.Printf("\t *** %s", err)
				return "", err
			}
		}
	}
	return "", nil
}

func (v *Vote) save(stub shim.ChaincodeStubInterface) error {
	var err error
	voteBytesToWrite, err := json.Marshal(&v)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return err
	}
	err = stub.PutState(v.Uuid, voteBytesToWrite)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return err
	}
	fmt.Printf("\t --- Saved vote %+v to blockchain\n", &v)
	return nil
}
