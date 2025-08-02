package types

// A Job requested by a client
type Job struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Body            string `json:"body"`
	CreatedDateTime string `json:"createdDateTime"`
}
