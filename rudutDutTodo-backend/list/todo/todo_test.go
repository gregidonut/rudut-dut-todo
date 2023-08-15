package todo_test

import (
	"encoding/json"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/todo"
	"github.com/jaswdr/faker"
	"math/rand"
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
	}{
		{
			name: "Initial",
			args: args{
				postNumber: 0,
				date:       time.Now(),
				id:         randStringBytes(16),
				title:      faker.New().Lorem().Sentence(4),
				content:    faker.New().Lorem().Paragraph(3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := todo.NewTodo(
				tt.args.postNumber,
				tt.args.date,
				tt.args.id,
				tt.args.title,
				tt.args.content,
			)

			want := &todo.Todo{
				MongoID:    tt.args.id,
				PostNumber: tt.args.postNumber,
				Date:       tt.args.date,
				Title:      tt.args.title,
				Content:    tt.args.content,
				Progress: &todo.Progress{
					Todo:       true,
					InProgress: false,
					Finished:   false,
				},
			}

			if !reflect.DeepEqual(got, want) {
				byteArray, _ := json.MarshalIndent(got, "", "\t")
				got := string(byteArray)
				t.Errorf("NewTodo() = \n%s,\n want: %#v\n",
					got,
					want,
				)
			}
		})
	}
}

// Apparently this is a performant way of generating a random string
// acording to https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randStringBytes(n int) string {
	// this is to be used in the test case where a random string is mocked like
	// the mongoDb _id field
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
