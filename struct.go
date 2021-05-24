package main

// video's owner info (user)
//owner[name] =
type owner struct {
	Name          string           `json:"name"`
	Videos        map[string]video `json:"videos"`
	ReputationRaw float64          `json:"reputationRaw"`
	IsVoter       bool             `json:"isVoter"`
	Identity      []byte           `json:"identity"`
}

// video info
//video[id] =
type video struct {
	Id                      string                  `json:"id"`
	Owner                   owner                   `json:"owner"`
	Metadata                string                  `json:"metadata"`
	ContractInfo            transferContractInfo    `json:"contractInfo"`
	ForTransferContractInfo forTransferContractInfo `json:"forTransferContractInfo"`
}

//contractClass = RF, RE, CC
type transferContractInfo struct {
	Contractor   string                  `json:"contractor"` //contract presenter
	Contractee   string                  `json:"contractee"` //parent video's owner
	ContractInfo forTransferContractInfo `json:"contractClass"`
	ParentVideo  string                  `json:"parentVideo"`
}

//@ RF(royalty free) 		1, fee 0
//@ RE(royalty exist) 		2, fee exist
//@ CC(creative commons) 	3 fee 0
type forTransferContractInfo struct {
	ContractClass int     `json:"contractClass"`
	ContractFee   float64 `json:"contractFee"`
}

type contractAlert struct {
	Contractor []byte `json:"contractor"`
	Contractee []byte `json:"contractee"`
	Video      string `json:"video"`
}

type transferContractWaitingList struct {
	Contractor []byte `json:"contractor"`
	Contractee []byte `json:"contractee"`
	Video      string `json:"video"`
	Isfine     bool   `json:"isfine"`
}
