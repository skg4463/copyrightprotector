package main

type Vote struct {
	Uuid string `json:"uuid"`
}

var PRIMARYKEY = [3]string{"Parties", "Votes", "Candidates"}

type Party struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Voter         bool     `json:"voter"`
	Candidate     bool     `json:"candidate"`
	VotesToAssign []string `json:"votestoassign"`
	VotesReceived []string `json:"votesreceived"`
	CandidateUrl  string   `json:"candidateUrl"`
	ScreenshotUrl string   `json:"screenshotUrl"`
}
