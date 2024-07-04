package models

type Card struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Question  string `json:"question" bson:"question"`
	Answer    string `json:"answer" bson:"answer"`
	Workspace string `json:"workspace" bson:"workspace"`
}
