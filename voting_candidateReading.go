package main

import (
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"sort"
	"strconv"
)

type Candidates []Party // To assign the sorting functions

func readAllCandidates(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var emptyArgs []string
	if len(args) != 0 {
		err = errors.New("{\"Error\":\"Expecting 0 arguments, got " + strconv.Itoa(len(args)))
		PrintErrorFull("", err)
		return nil, err
	}

	// Get candidates main ledger.
	candidateIds, err := getDataArrayStrings(stub, PRIMARYKEY[2], emptyArgs)
	if err != nil {
		PrintErrorFull("readAllCandidates - getDataArrayStrings", err)
		return nil, err
	}

	// Iterate over all candidates to get the full details
	if len(candidateIds) > 0 {
		// Initialise an empty slice for the output
		var candidatesLedger []Party
		// Iterate over all parties and return the party object.
		for _, candidateId := range candidateIds {
			thisCandidate, err := getParty(stub, []string{candidateId})
			if err != nil {
				PrintErrorFull("readAllCandidates - getParty", err)
				return nil, err
			}
			candidatesLedger = append(candidatesLedger, thisCandidate)
		}
		// Sort the ledger by... number of votes received. (len(VotesReceived))
		sort.Sort(Candidates(candidatesLedger))
		// This gives us an slice with parties. Translate to bytes and return
		partiesLedgerBytes, err := json.Marshal(&candidatesLedger)
		if err != nil {
			PrintErrorFull("readAllCandidates - Marshal", err)
			return nil, err
		}
		PrintSuccess("Retrieved full information for all Parties.")
		return partiesLedgerBytes, nil
	} else {
		return nil, nil
	}
} // readAllCandidates
