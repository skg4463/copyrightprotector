package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"strconv"
)

func readParty(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	if len(args) != 1 { // id
		err = errors.New("{\"Error\":\"Expecting 1 arguments, got " + strconv.Itoa(len(args)))
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	id := args[0]
	var returnSlice []Party
	party, err := getParty(stub, []string{id})
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	returnSlice = append(returnSlice, party)
	// This gives us an slice with parties. Translate to bytes and return
	returnSliceBytes, err := json.Marshal(&returnSlice)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	fmt.Printf("\t--- Retrieved full information for Party %s", id)
	return returnSliceBytes, nil
}

func readAllParties(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var emptyArgs []string
	if len(args) != 0 {
		err = errors.New("{\"Error\":\"Expecting 0 arguments, got " + strconv.Itoa(len(args)))
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	// Get all parties - returns an slice of strings - partyIds
	partyIds, err := getDataArrayStrings(stub, PRIMARYKEY[0], emptyArgs)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	if len(partyIds) > 0 {
		// Initialise an empty slice for the output
		var partiesLedger []Party
		// Iterate over all parties and return the party object.
		for _, partyId := range partyIds {
			thisParty, err := getParty(stub, []string{partyId})
			if err != nil {
				fmt.Printf("\t *** %s", err)
				return nil, err
			}
			partiesLedger = append(partiesLedger, thisParty)
		}
		// This gives us an slice with parties. Translate to bytes and return
		partiesLedgerBytes, err := json.Marshal(&partiesLedger)
		if err != nil {
			fmt.Printf("\t *** %s", err)
			return nil, err
		}
		fmt.Println("\t--- Retrieved full information for all Parties.")
		return partiesLedgerBytes, nil
	} else {
		return nil, nil
	}
	return nil, nil // redundancy
}
