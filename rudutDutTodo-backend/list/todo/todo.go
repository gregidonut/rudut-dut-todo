package todo

import (
	"encoding/json"
	"fmt"
	"time"
)

type Progress struct {
	Todo       bool
	InProgress bool
	Finished   bool
}

func NewProgress() *Progress {
	return &Progress{
		Todo:       true,
		InProgress: false,
		Finished:   false,
	}
}

// MakeSureOneOfThree will be useful to make sure that only
// one of three states is used
func (p *Progress) MakeSureOneOfThree() error {
	var states int

	// Must be checked separately since a switch statement will return
	// if only one or more of them were true.
	if p.Todo {
		states++
	}
	if p.InProgress {
		states++
	}
	if p.Finished {
		states++
	}

	if states > 1 {
		return MoreThanOneStateErr
	}

	return nil
}

type Todo struct {
	MongoID    string
	PostNumber int
	Date       time.Time
	Title      string
	Content    string
	Progress   *Progress
}

func NewTodo(object json.RawMessage) (*Todo, error) {
	var returnVal Todo
	var aux map[string]interface{}

	err := json.Unmarshal(object, &aux)
	if err != nil {
		return &Todo{}, fmt.Errorf("%v: %v\n", TodoInstantiationErr, err)
	}

	mongoID, ok := aux["_id"]
	if ok {
		returnVal.MongoID = mongoID.(string)
	}
	err = json.Unmarshal(object, &returnVal)
	if err != nil {
		return &Todo{}, fmt.Errorf("%v: %v\n", TodoInstantiationErr, err)
	}

	return &returnVal, nil
}
