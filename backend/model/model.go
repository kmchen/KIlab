package model

type Interviewer struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	TimeSlot []TimeSlot
}

type TimeSlot struct {
	Time int64  `json:"time"`
	Id   uint32 `json:"Id"`
}

type Candidate struct {
	Id       uint32 `json:"id"`
	Name     string `json:"name"`
	TimeSlot []TimeSlot
}

type Payload struct {
	Interviewer string  `json:"Interviewer, omitempty"`
	Candidate   string  `json:"Candidate, omitempty"`
	Time        []int64 `json:"time"`
}
