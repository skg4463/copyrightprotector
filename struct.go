package main

// video's owner info (user)
//owner[name] =
type owner struct {
	Name          string           `json:"name"`
	Videos        map[string]video `json:"videos"`
	ReputationRaw float64          `json:"reputationRaw"`
	IsVoter       bool             `json:"isVoter"`
}

// video info
//video[id] =
type video struct {
	Id                      string                  `json:"id"`
	Owner                   string                  `json:"owner"`
	Metadata                string                  `json:"metadata"`
	contractInfo            transferContract        `json:"contractInfo"`
	forTransferContractInfo forTransferContractInfo `json:"forTransferContractInfo"`
	//contract info
	//transfer contract info
}

//contractClass = RF, RE, CC
type transferContract struct {
	contractor   string                  `json:"contractor"` //contract presenter
	contractee   string                  `json:"contractee"` //parent video's owner
	contractInfo forTransferContractInfo `json:"contractClass"`
	parentVideo  string                  `json:"parentVideo"`
}

//@ RF(royalty free) 0, fee 0
//@ RE(royalty exist) 1, fee exist
//@ CC(creative commons) 2 fee 0
type forTransferContractInfo struct {
	contractClass int     `json:"contractClass"`
	contractFee   float64 `json:"contractFee"`
}
