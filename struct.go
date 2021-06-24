package main

// video's owner info (user)
//owner[name] = @
type owner struct {
	Name          string           `json:"name"`          //username
	Videos        map[string]video `json:"videos"`        //video list
	ReputationRaw float64          `json:"reputationRaw"` //reputation raw value
	IsVoter       bool             `json:"isVoter"`       //is this user owner?
	Identity      interface{}      `json:"identity"`      //owner cert
}

// video info
//video[id] = @
type video struct {
	Id                      string                  `json:"id"`       //video serial
	Owner                   owner                   `json:"owner"`    //owner info
	Metadata                string                  `json:"metadata"` //video metadata
	ContractedInfo          transferContractedInfo  `json:"contractInfo"`
	ForTransferContractInfo forTransferContractInfo `json:"forTransferContractInfo"`
}

//After Transfer Contract
//transfer ContractedInfo
//transferContractInfo[video Id] = @
type transferContractedInfo struct {
	Contractor   string                  `json:"contractor"`    //contract presenter
	Contractee   string                  `json:"contractee"`    //parent video's owner
	ContractInfo forTransferContractInfo `json:"contractClass"` //for transfer contract Info
	ParentVideo  string                  `json:"parentVideo"`
}

//Before Transfer Contract
//@ RF(royalty free) 		1, fee 0
//@ RE(royalty exist) 		2, fee exist
//@ CC(creative commons) 	3, fee 0
//@@ADD Fee Contract
type forTransferContractInfo struct {
	ContractClass int     `json:"contractClass"`
	ContractFee   float64 `json:"contractFee"`
}

//Alert Event params
type contractAlert struct {
	Contractor interface{} `json:"contractor"` //contract presenter
	Contractee interface{} `json:"contractee"` //parent video's owner
	Video      string      `json:"video"`
}

type transferContractWaitingList struct {
	Contractor interface{} `json:"contractor"` //contract presenter
	Contractee interface{} `json:"contractee"` //parent video's owner
	Video      string      `json:"video"`
	Isfine     bool        `json:"isfine"`
}

var contractSerial = 0
