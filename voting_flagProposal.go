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
} // end of dcc.createvotesAndAssignToAll

func readVote(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	if len(args) != 1 { // id
		err = errors.New("{\"Error\":\"Expecting 1 arguments, got " + strconv.Itoa(len(args)))
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	id := args[0]
	var returnSlice []Vote
	party, err := getVote(stub, []string{id})
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
	fmt.Printf("\t--- Retrieved full information for Vote %s\n", id)
	return returnSliceBytes, nil
}

func getVote(stub shim.ChaincodeStubInterface, args []string) (Vote, error) {
	var vote Vote // We need to have an empty vote ready to return in case of an error.
	var err error
	if len(args) != 1 { // Only needs a vote id.
		err = errors.New("{\"Error\":\"Incorrect number of arguments\", \"Function\":\"getVote\"}")
		fmt.Printf("\t *** %s", err)
		return vote, err
	}
	voteId := args[0]
	voteBytes, err := stub.GetState(voteId)
	if voteBytes == nil {
		err = errors.New("{\"Error\":\"State " + voteId + " does not exist\", \"Function\":\"getVote\"}")
		fmt.Printf("\t *** %s", err)
		return vote, err
	}
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return vote, err
	}
	err = json.Unmarshal(voteBytes, &vote)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return vote, err
	}
	return vote, nil
} // end of dcc.getVote

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
