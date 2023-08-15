package todo_test

import (
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/todo"
	"reflect"
	"testing"
	"time"
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
		name        string
		fields      fields
		wantErr     bool
		expectedErr error
	}{
		{
			name: "OnlyTodoIsTrue",
			fields: fields{
				todo:       true,
				inProgress: false,
				finished:   false,
			},
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name: "TodoAndInProgIsTrue",
			fields: fields{
				todo:       true,
				inProgress: true,
				finished:   false,
			},
			wantErr:     true,
			expectedErr: todo.MoreThanOneStateErr,
		},
		{
			name: "OnlyFinishedIsTrue",
			fields: fields{
				todo:       false,
				inProgress: false,
				finished:   true,
			},
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name: "FinishedAndInProgIsTrue",
			fields: fields{
				todo:       false,
				inProgress: true,
				finished:   true,
			},
			wantErr:     true,
			expectedErr: todo.MoreThanOneStateErr,
		},
		{
			name: "TodoAndFinishedIsTrue",
			fields: fields{
				todo:       true,
				inProgress: false,
				finished:   true,
			},
			wantErr:     true,
			expectedErr: todo.MoreThanOneStateErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &todo.Progress{
				Todo:       tt.fields.todo,
				InProgress: tt.fields.inProgress,
				Finished:   tt.fields.finished,
			}
			err := p.MakeSureOneOfThree()
			if tt.wantErr {
				if err == nil {
					t.Fatalf(
						"expected err: %q, but didn't get one\n",
						tt.expectedErr,
					)
				}

				if tt.expectedErr != err {
					t.Fatalf(
						"expected err: %q, but got different err: %q\n",
						tt.expectedErr,
						err,
					)
				}

				return
			}
			if err != nil {
				t.Fatalf(
					"did not expect err but got: %q\n",
					err,
				)
			}
		})
	}
}

func TestNewTodo(t *testing.T) {
	type args struct {
		postNumber int
		date       time.Time
		id         string
		title      string
		content    string
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
			got := todo.NewTodo(tt.args.postNumber, tt.args.date, tt.args.id, tt.args.title, tt.args.content)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}
