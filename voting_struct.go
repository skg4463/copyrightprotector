package main

type Vote struct {
	Uuid string `json:"uuid"`
}

var PRIMARYKEY = [3]string{"Parties", "Votes", "Candidates"}

type Party struct {
	Id            string   `json:"id"`            //unique serial
	Name          string   `json:"name"`          //username or video id
	Voter         bool     `json:"voter"`         //voter?
	Contents      bool     `json:"Contents"`      //Contents?
	VotesToAssign []string `json:"votestoassign"` //제출투표
	VotesReceived []string `json:"votesreceived"` //받은투표
	flagProposer  string   `json:"flagProposer"`  //flager
	repu          int      `json:"repu"`          //repetaion int
}
