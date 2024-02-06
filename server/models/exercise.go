package model

type Example struct {
	input  string
	output string
}

type Exercise struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Examples    []Example
}
type Answer struct {
	Function   string `json:"function" bson:"function"`
	Lenguage   string `json:"lenguage" bson:"lenguage"`
	ID         string `json:"id" bson:"_id,omitempty"`
	ExerciseId string `json:"exerciseId" bson:"exerciseId"`
}
