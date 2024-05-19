package models

type Note struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	OwnerId string `json:"owner_id"`
}
