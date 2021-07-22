package main

type Vote struct {
	Uuid string `json:"uuid"`
}

var PRIMARYKEY = [3]string{"Parties", "Votes", "Candidates"}

type Party struct {
	Id            string   `json:"id"`            //unique serial
	Name          string   `json:"name"`          //username
	Voter         bool     `json:"voter"`         //voter?
	Candidate     bool     `json:"candidate"`     //candidate?
	VotesToAssign []string `json:"votestoassign"` //제출투표
	VotesReceived []string `json:"votesreceived"` //받은투표
}
