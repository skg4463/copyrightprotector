package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func (slice Candidates) Len() int {
	return len(slice)
}

func (slice Candidates) Less(i, j int) bool {
	return len(slice[i].VotesReceived) > len(slice[j].VotesReceived)
}

func (slice Candidates) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Query entry point for queries
//func Query(stub shim.ChaincodeStubInterface, fn string, args []string) ([]byte, error) {
//	// Handle different functions
//	if fn == "read" { // read a variable
//		return read(stub, fn, args)
//	} else if fn == "readParty" {
//		return readParty(stub, fn, args)
//	} else if fn == "readAllParties" {
//		return readAllParties(stub, fn, args)
//	} else if fn == "readAllCandidates" {
//		return readAllCandidates(stub, fn, args)
//	}
//	fmt.Println("\t*** ERROR: Query function did not find ChainCode function: " + fn)
//	return nil, errors.New(" --- QUERY ERROR: Received unknown function query")
//} back commit

// ============================================================================================================================

// Function that reads the bytes associated with a data-key and returns the byte-array.
func read(stub shim.ChaincodeStubInterface, fn string, args []string) ([]byte, error) {
	var err error
	if len(args) != 1 { // needs a data key to read.
		err = errors.New("{\"Error\":\"Incorrect number of arguments\", \"Function\":\"" + fn + "\"}")
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	dataKey := args[0]
	dataBytes, err := stub.GetState(dataKey)
	if dataBytes == nil { // deals with non existing data keys.
		err = errors.New("{\"Error\":\"State " + dataKey + " does not exist\", \"Function\":\"" + fn + "\"}")
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	if err != nil {
		err = errors.New("{\"Error\":\"Failed to get state for " + dataKey + "\", \"Function\":\"" + fn + "\"}")
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	fmt.Println("\t--- Read the ledger: " + dataKey)
	return dataBytes, nil
}

func readAll(stub shim.ChaincodeStubInterface, fn string, args []string) ([]byte, error) {
	var err error
	var emptyArgs []string
	if len(args) != 0 {
		err = errors.New("{\"Error\":\"Incorrect number of arguments\", \"Function\":\"" + fn + "\"}")
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	// get all parties - returns an array of strings.
	partiesLedger, err := getDataArrayStrings(stub, PRIMARYKEY[0], emptyArgs)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	// get all votes
	votesLedger, err := getDataArrayStrings(stub, PRIMARYKEY[1], emptyArgs)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	// Create a map of both
	m := map[string][]string{
		PRIMARYKEY[0]: partiesLedger,
		PRIMARYKEY[1]: votesLedger,
	}
	// Cast to JSON
	mStr, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	// Return as bytes.
	fmt.Println("\t--- Read the main ledgers ")
	out := []byte(string(mStr))
	return out, nil
}

func getDataArrayStrings(stub shim.ChaincodeStubInterface, dataKey string, args []string) ([]string, error) {
	var err error
	var empty []string
	if len(args) != 0 {
		err = errors.New("{\"Error\":\"Incorrect number of arguments\", \"Function\":\"getDataArrayStrings\"}")
		fmt.Printf("\t *** %s", err)
		return empty, err
	}
	arrayBytes, err := stub.GetState(dataKey)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return empty, err
	}
	var outputArray []string
	err = json.Unmarshal(arrayBytes, &outputArray)
	if err != nil {
		fmt.Printf("\t *** %s", err)
		return empty, err
	}
	return outputArray, nil
}

func saveStringToDataArray(stub shim.ChaincodeStubInterface, dataKey string, addString string, ledger []string) ([]byte, error) {
	var err error
	// Add the string to the array
	ledger = append(ledger, addString)
	// err = dcc.saveLedger(stub, dataKey, ledger)
	if err = saveLedger(stub, dataKey, ledger); err != nil {
		fmt.Printf("\t *** %s", err)
		return nil, err
	}
	return nil, nil
}

func saveLedger(stub shim.ChaincodeStubInterface, dataKey string, ledger []string) error {
	var err error
	// Marshall the ledger to bytes
	bytesToWrite, err := json.Marshal(&ledger)
	if err != nil {
		return err
	}
	// Save the array.
	err = stub.PutState(dataKey, bytesToWrite)
	if err != nil {
		return err
	}
	return nil
}

//---

// HashSHA256 Function to return the SHA256 for
func HashSHA256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	sha256Hash := hex.EncodeToString(h.Sum(nil))
	return sha256Hash
}

//---

func PrintError(str string) {
	fmt.Println("\t \x1b[31;1m  *** " + str + " \x1b[0m")
}

func PrintErrorFull(str string, err error) {
	fmt.Println("\t \x1b[31;1m  *** "+str+": %s \x1b[0m", err)
}

func PrintStatus(str string) {
	fmt.Println("\x1b[32;1m  ~~~ " + str + "\x1b[0m")
}

func PrintSuccess(str string) {
	fmt.Println("\t \x1b[35;1m  \u2713\u2713\u2713 " + str + "\x1b[0m")
}

func PrintLine() {
	fmt.Println("\x1b[33;1m \n  ---------------------------------------------------------------- \x1b[0m")
}

//---

func IsElementInSlice(slice []string, element string) bool {
	// Initialise return as false.
	check := false
	// Iterate over the list to see if the value is in there.
	for _, val := range slice {
		if val == element {
			check = true
			return check
		}
	}
	// Failsafe return.
	return check
}

// FindElementIndex Function to return the index of an element in a slice.
func FindElementIndex(slice []string, element string) int {
	// Initialise return as false.
	ix := -1
	// Iterate over the list to see if the value is in there.
	for i, val := range slice {
		if val == element {
			return i
		}
	}
	// Failsafe return.
	return ix
}

// DeleteElementFromSlice Function to delete an element from a slice.?
func DeleteElementFromSlice(slice []string, element string) []string {
	var emptySlice []string
	if len(slice) == 0 {
		return emptySlice
	}
	if !IsElementInSlice(slice, element) { // It is not in the slice.
		return slice
	}
	if len(slice) == 1 {
		return emptySlice
	}
	ix := FindElementIndex(slice, element)
	slice = append(slice[:ix], slice[ix+1:]...)
	return slice
}
