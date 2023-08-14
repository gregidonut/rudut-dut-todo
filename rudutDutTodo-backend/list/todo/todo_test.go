package todo_test

import (
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/todo"
	"reflect"
	"testing"
)

func TestNewProgress(t *testing.T) {
	tests := []struct {
		name string
		want *todo.Progress
	}{
		{
			name: "InitialProgressIsTodo",
			want: &todo.Progress{
				Todo:       true,
				InProgress: false,
				Finished:   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := todo.NewProgress()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProgress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgress_makeSureOneOfThree(t *testing.T) {
	type fields struct {
		todo       bool
		inProgress bool
		finished   bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &todo.Progress{
				Todo:       tt.fields.todo,
				InProgress: tt.fields.inProgress,
				Finished:   tt.fields.finished,
			}
			if err := p.MakeSureOneOfThree(); (err != nil) != tt.wantErr {
				t.Errorf("makeSureOneOfThree() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewTodo(t *testing.T) {
	type args struct {
		postNumber int
		id         string
		title      string
		content    string
		date       string
	}
	tests := []struct {
		name string
		args args
		want *todo.Todo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := todo.NewTodo(tt.args.postNumber, tt.args.id, tt.args.title, tt.args.content, tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}
