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

//transferContract info
type transferContract struct {
	Owner string `json:"Owner"`
}
