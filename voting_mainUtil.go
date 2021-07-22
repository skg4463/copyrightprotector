package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"strconv"
)

//												party ID
func getParty(stub shim.ChaincodeStubInterface, args []string) (Party, error) {
	var party Party
	var err error
	if len(args) != 1 {
		err = errors.New("{\"Error\":\"Incorrect number of arguments\", \"Function\":\"getParty\"}")
		fmt.Printf("\t *** %s", err)
		return party, err
	}
	partyId := args[0]
	partyBytes, err := stub.GetState(partyId)
	if partyBytes == nil {
		err = errors.New("{\"Error\":\"State " + partyId + " does not exist\", \"Function\":\"getParty\"}")
		fmt.Printf("\t *** %s", err)
		return party, err
	}
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return party, err
	}
	err = json.Unmarshal(partyBytes, &party)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return party, err
	}
	return party, nil
}

// Id, VotesToAssign, VotesTransferred, VotesReceived
func updateParty(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	var err error
	if len(args) != 4 {
		err = errors.New("{\"Error\":\"Expecting 4 arguments, got " + strconv.Itoa(len(args)))
		fmt.Printf("\t *** %s", err)
		return "", err
	}

	// Load party
	partyId := args[0]
	party, err := getParty(stub, []string{partyId})
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return "", err
	}

	//voter---------------
	// add voting weight by reputationCalc(owner[party.name].reputationRaw)
	// if party is a voter
	// add vote uuid to VotesToAssign slice
	voteToAssign := args[1]
	if party.Voter && voteToAssign != "" {
		party.VotesToAssign = append(party.VotesToAssign, voteToAssign)
	}

	// if party is a voter and vote transfer
	// delete from VotesToAssign slice
	voteTransferred := args[2]
	if party.Voter && voteTransferred != "" {
		// check if vote exists
		args := []string{voteTransferred}
		_, err := readVote(stub, args)
		if err != nil {
			fmt.Printf("\t *** %s", err)
			return "", err
		}
		for i, v := range party.VotesToAssign {
			if v == voteTransferred {
				party.VotesToAssign = append(party.VotesToAssign[:i], party.VotesToAssign[i+1:]...)
			}
		}
	}

	//candidate---------------
	// if party is a candidate
	// add vote uuid to VotesReceived slice
	voteReceived := args[3]
	if party.Contents && voteReceived != "" {
		party.VotesReceived = append(party.VotesReceived, voteReceived)
	}

	// Save the new party.
	if err = party.save(stub); err != nil {
		fmt.Printf("\t *** %s", err)
		return "", err
	}
	fmt.Printf("\t --- Updated Party %s\n", partyId)
	return "", nil
}

func (p *Party) save(stub shim.ChaincodeStubInterface) error {
	var err error
	partyBytesToWrite, err := json.Marshal(&p)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return err
	}
	err = stub.PutState(p.Id, partyBytesToWrite)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return err
	}
	fmt.Printf("\t --- Saved party %v to blockchain\n", &p.Id)
	return nil
}

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
