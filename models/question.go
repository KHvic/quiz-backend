package models

import (
	"database/sql/driver"
	"encoding/json"
)

// QuestionType represent type of question
type QuestionType int

const (
	// FillBlank ...
	FillBlank = 1
	// MCQ ...
	MCQ = 2
	// MultiSelect ...
	MultiSelect = 3
)

// Options represent options type
type Options [][]string

// Value indicate how database store Options type
func (o Options) Value() (driver.Value, error) {
	bytes, _ := json.Marshal(o)
	return string(bytes), nil
}

// Scan indicate how application read Options type
func (o *Options) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	return json.Unmarshal(bytes, o)
}

// Answer represent answer type
type Answer [][]int

// Value indicate how database store Answer type
func (a Answer) Value() (driver.Value, error) {
	bytes, _ := json.Marshal(a)
	return string(bytes), nil
}

// Scan indicate how application read Answer type
func (a *Answer) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	return json.Unmarshal(bytes, a)
}

// Question ...
type Question struct {
	ID          int64
	Category    string
	SubCategory string
	Type        QuestionType
	Description string
	Options     Options // e.g [["option a", "option b"], ["option c, option d"]]
	Answer      Answer  // e.g [[1][0]]
	Image       *string
}

// TableName ...
func (*Question) TableName() string {
	return "question"
}
