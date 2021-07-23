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

//party struct extract to 'voter' and 'content'
//each Owner that isVoter is true flag have (party)voter struct
type Voter struct {
	Id              string   `json:"id"`               //unique serial
	Name            string   `json:"name"`             //username
	VotesToAssign   []string `json:"votestoassign"`    //assigned VOTE arr
	reputationScore int      `json:"reputation_score"` //reputation score for voting score recording
}

//Voter, Contents boolean class !CHECK!

//each Video have (party)contents struct
type Contents struct {
	Id            string   `json:"id"`            //unique serial
	Name          string   `json:"name"`          //video ID in struct VIDEO
	votesReceived []string `json:"votesreceived"` //received vote to Voter
	flagProposer  string   `json:"flag_proposer"` //flag proposer's OWNER namespace
}
