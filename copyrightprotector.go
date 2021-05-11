package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type copyrightprotector struct {
}

type owner struct {
	Name   string           `json:"Name"`
	Videos map[string]video `json:"Videos"`
}

type video struct {
	Id       string `json:"Id"`
	Owner    string `json:"Owner"`
	Metadata string `json:"Metadata"`
}

func (t *copyrightprotector) Init(shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *copyrightprotector) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error

	switch fn {
	case "initLedger":
		result, err = initLedger(stub)
	case "queryOwner":
		result, err = queryOwner(stub, args)
	case "queryVideo":
		result, err = queryVideo(stub, args)
	case "addVideo":
		result, err = addVideo(stub, args)
	case "purchaseVideo":
		result, err = purchaseVideo(stub, args)
	case "queryOwner1":
		result, err = queryOwner1(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(result))
}

func initLedger(stub shim.ChaincodeStubInterface) (string, error) {
	owners := map[string]owner{}
	videos := map[string]video{}

	owners["nakhoon"] = owner{
		Name:   "Nakhoon",
		Videos: map[string]video{},
	}
	videos["0"] = video{
		Id:       "DCPoTnakAe0",
		Owner:    "nakhoon",
		Metadata: "meta1",
	}
	videos["1"] = video{
		Id:       "id",
		Owner:    "owner",
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

func queryOwner(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	owners := map[string]owner{}
	ownersAsBytes, err := stub.GetState("Owners")
	_ = json.Unmarshal(ownersAsBytes, &owners)

	selectedownerAsBytes, err := json.Marshal(owners[args[0]])

	if err != nil {
		return "", fmt.Errorf("Failed to get video: %s with error: %s", args[0], err)
	}
	if selectedownerAsBytes == nil {
		return "", fmt.Errorf("video not found: %s", args[0])
	}
	fmt.Println(selectedownerAsBytes)
	return string(selectedownerAsBytes), nil
}

func queryVideo(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting an id")
	}

	videos := map[string]video{}
	videosAsBytes, err := stub.GetState("Videos")
	_ = json.Unmarshal(videosAsBytes, &videos)
	selectedvideoAsBytes, err := json.Marshal(videos[args[0]])

	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if selectedvideoAsBytes == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(selectedvideoAsBytes), nil
}

func addVideo(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 3 {
		return "", fmt.Errorf("Incorrect arguments. Expecting 3 arguments (id, owner, metadata)")
	}
	videos := map[string]video{}
	videossAsBytes, _ := stub.GetState("Videos")
	_ = json.Unmarshal(videossAsBytes, &videos)

	videos[args[0]] = video{
		Id:       args[0],
		Owner:    args[1],
		Metadata: args[2],
	}

	updatedvideoAsBytes, _ := json.Marshal(videos)
	err := stub.PutState("Videos", updatedvideoAsBytes)
	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return string(updatedvideoAsBytes), nil
}

func purchaseVideo(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting two arguments (user name, and video id)")
	}
	videosAsBytes, _ := stub.GetState("Videos")
	ownersAsBytes, _ := stub.GetState("Owners")

	videos := map[string]video{}
	owners := map[string]owner{}

	_ = json.Unmarshal(videosAsBytes, &videos)
	_ = json.Unmarshal(ownersAsBytes, &owners)

	videoowner := owners[args[0]]
	selectedvideo := videos[args[1]]

	videoowner.Videos[selectedvideo.Id] = selectedvideo

	owners[args[0]] = videoowner

	updatedownerAsBytes, _ := json.Marshal(owners)
	err := stub.PutState("Owners", updatedownerAsBytes)

	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return string(updatedownerAsBytes), nil
}

func main() {
	err := shim.Start(new(copyrightprotector))
	if err != nil {
		fmt.Println("Could not start copyrightprotector Chaincode")
	} else {
		fmt.Println("copyrightprotector Chaincode successfully started")
	}
}
