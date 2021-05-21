package main

// video owner info
type owner struct {
	Name          string           `json:"Name"`
	Videos        map[string]video `json:"Videos"`
	ReputationRaw float64          `json:"ReputationRaw"`
	IsVoter       bool             `json:"IsVoter"`
}

// video info
///each video transaction
type video struct {
	Id       string `json:"Id"`
	Owner    string `json:"Owner"`
	Metadata string `json:"Metadata"`
	//contract info
	//transfer contract info
}

//contractClass =
//@ RF(royalty free) 0,
//@ RE(royalty exist) 1,
//@ CC(creative commons) 2
type transferContract struct {
	contractor    owner   `json:"contractor"` //contract presenter
	contractee    owner   `json:"contractee"` //video owner
	contractClass int     `json:"contractClass"`
	contractFee   float64 `json:"contractFee"`
}
