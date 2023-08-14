package todo

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
	return nil
}

type Todo struct {
	mongoID    string
	postNumber int
	date       string
	title      string
	content    string
	progress   *Progress
}

func NewTodo(postNumber int, id, title, content, date string) *Todo {
	return &Todo{
		mongoID:    id,
		postNumber: postNumber,
		date:       date,
		title:      title,
		content:    content,
		progress:   NewProgress(),
	}
}
