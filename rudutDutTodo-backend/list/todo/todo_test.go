package todo_test

import (
	"encoding/json"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/todo"
	"reflect"
	"strings"
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
	testObj := `{
	"_id": "64ddbdaae24251218972e72f",
	"content": "a random string of todo Content",
	"date": "2023-08-17T06:26:50.497Z",
	"postId": 0,
	"progress": {
		"finished": false,
		"inProgress": false,
		"todo": true
	},
	"title": "this is a test todo item"
}`
	dateFromTestObj, err := time.Parse(time.RFC3339, "2023-08-17T06:26:50.497Z")
	if err != nil {
		t.Errorf("having throuble parsing date %q\n", dateFromTestObj)
	}

	type args struct {
		object json.RawMessage
	}
	tests := []struct {
		name        string
		args        args
		want        *todo.Todo
		wantErr     bool
		expectedErr error
	}{
		{
			name: "initial",
			args: args{
				object: json.RawMessage(testObj),
			},
			want: &todo.Todo{
				MongoID:    "64ddbdaae24251218972e72f",
				PostNumber: 0,
				Date:       dateFromTestObj,
				Title:      "this is a test todo item",
				Content:    "a random string of todo Content",
				Progress: &todo.Progress{
					Todo:       true,
					InProgress: false,
					Finished:   false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := todo.NewTodo(tt.args.object)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected err: %q, but didn't get one\n", tt.expectedErr)
				}
				if !strings.Contains(err.Error(), tt.expectedErr.Error()) {
					t.Fatalf("expected err: %q, to contain: %q, but did not\n", err, tt.expectedErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %q\n", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
