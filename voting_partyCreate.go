package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"strconv"
)

//user to vote member
//args: id, name, voter, contentsID
//votestoassign, votesreceived
func createParty(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	var err error
	var emptyArgs []string
	if len(args) != 4 {
		err = errors.New("{\"Error\":\"Expecting 4 arguments, got " + strconv.Itoa(len(args)))
		fmt.Printf("\t *** %s", err)
		return "", err
	}
	// The partyId needs to be unique
	partyId := args[0]

	// Get all the party
	partyIds, err := getDataArrayStrings(stub, PRIMARYKEY[0], emptyArgs)
	if err != nil {
		PrintErrorFull("createParty - getDataArrayStrings", err)
		return "", err
	}

	// Get all the candidate
	candidateIds, err := getDataArrayStrings(stub, PRIMARYKEY[2], emptyArgs)
	if err != nil {
		PrintErrorFull("createParty - getDataArrayStrings", err)
		return "", err
	}

	// Check if the partyId exists in the current ledger of partiㅛ
	partyExists := IsElementInSlice(partyIds, partyId)
	if partyExists == false {
		voter, err := strconv.ParseBool(args[2])
		if err != nil {
			fmt.Printf("\t *** %s", err)
			return "", err
		}
		Contents, err := strconv.ParseBool(args[3])
		if err != nil {
			fmt.Printf("\t *** %s", err)
			return "", err
		}

		// Create a new party
		var newParty = Party{
			Id:       partyId,
			Name:     args[1],
			Voter:    voter,
			Contents: Contents,
		}

		// Save new party
		if err = newParty.save(stub); err != nil {
			fmt.Printf("\t *** %s", err)
			return "", err
		}
		// Add party to the ledger.
		_, err = saveStringToDataArray(stub, PRIMARYKEY[0], partyId, partyIds)
		if err != nil {
			PrintErrorFull("createParty - saveStringToDataArray", err)
			return "", err
		}
		// If it is a candidate, add the the candidates-ledger
		if newParty.Contents {
			_, err = saveStringToDataArray(stub, PRIMARYKEY[2], partyId, candidateIds)
			if err != nil {
				PrintErrorFull("createParty - saveStringToDataArray", err)
				return "", err
			}
		}

		PrintSuccess("Added a new party: " + partyId)
		return "", nil
	} else { //if exist
		err = errors.New(partyId + "` already exists.")
		PrintErrorFull("createParty", err)
		return "", err
	}

	return "", nil
}
