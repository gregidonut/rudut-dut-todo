package todo

import "time"

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
	mongoID    string
	postNumber int
	date       time.Time
	title      string
	content    string
	progress   *Progress
}

func NewTodo(postNumber int, date time.Time, id, title, content string) *Todo {
	return &Todo{
		mongoID:    id,
		postNumber: postNumber,
		date:       date,
		title:      title,
		content:    content,
		progress:   NewProgress(),
	}
}
