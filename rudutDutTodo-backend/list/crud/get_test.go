package crud_test

import (
	"encoding/json"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/contact"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/crud"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/todo"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestGetList(t *testing.T) {
	mongoHandles, err := localJsonToStruct()
	if err != nil {
		t.Fatalf("having trouble setting up mongoHandles: %q", err)
	}

	testTodoListCol := mongoHandles.DBs[0].Info.Handles[0].Collections[1]

	type DBCollectionItemsID struct {
		OID string
	}

	type DBCollectionItemsDate struct {
		JSDateObject string
	}

	type DBCollectionItemsProgress struct {
		Todo       bool
		InProgress bool
		Finished   bool
	}

	type DBCollectionItems struct {
		ID       DBCollectionItemsID
		Content  string
		Date     DBCollectionItemsDate
		PostID   int
		Progress DBCollectionItemsProgress
		Title    string
	}

	type DBCollection struct {
		CName string
		Items DBCollectionItems
	}

	type dbContArgs struct {
		DBName       string
		DBCollection DBCollection
	}

	tests := []struct {
		name        string
		setupEnvVar bool
		envVarValue string
		dbContArgs  dbContArgs
		wantErr     bool
		expectedErr error
	}{
		{
			name:        "get from testDB.test-todo-list.first-item",
			setupEnvVar: true,
			envVarValue: mongoHandles.DBs[0].Info.URI,
			dbContArgs: dbContArgs{
				DBName: mongoHandles.DBs[0].Info.Handles[0].Name,
				DBCollection: DBCollection{
					CName: testTodoListCol.CName,
					Items: DBCollectionItems{
						ID: DBCollectionItemsID{
							OID: testTodoListCol.Items[0].ID.OID,
						},
						Content: testTodoListCol.Items[0].Content,
						Date: DBCollectionItemsDate{
							JSDateObject: testTodoListCol.Items[0].Date.JSDateObject,
						},
						PostID: testTodoListCol.Items[0].PostID,
						Progress: DBCollectionItemsProgress{
							Todo:       testTodoListCol.Items[0].Progress.Todo,
							InProgress: testTodoListCol.Items[0].Progress.InProgress,
							Finished:   testTodoListCol.Items[0].Progress.Finished,
						},
						Title: testTodoListCol.Items[0].Title,
					},
				},
			},
			wantErr:     false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			setupEnvVar(t, tt.envVarValue)

			dbCont, err := contact.NewDBContainer(
				tt.dbContArgs.DBName,
				tt.dbContArgs.DBCollection.CName,
			)

			got, err := crud.GetList(*dbCont)
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
				t.Fatalf("error pinging db: %q\n", err)
			}

			mockTodoItem := tt.dbContArgs.DBCollection.Items
			mockTime, err := time.Parse(time.RFC3339, mockTodoItem.Date.JSDateObject)
			mockTodoItemStruct := todo.Todo{
				MongoID:    mockTodoItem.ID.OID,
				PostNumber: 0,
				Date:       mockTime,
				Title:      mockTodoItem.Title,
				Content:    mockTodoItem.Content,
				Progress: &todo.Progress{
					Todo:       mockTodoItem.Progress.Todo,
					InProgress: mockTodoItem.Progress.InProgress,
					Finished:   mockTodoItem.Progress.Finished,
				},
			}
			if err != nil {
				t.Fatalf("error Marhsalling mockCollection: %q\n", err)
			}

			want := []todo.Todo{mockTodoItemStruct}
			if !reflect.DeepEqual(got, want) {
				gotMarshal, err := json.MarshalIndent(got, "", "    ")
				if err != nil {
					t.Fatalf("error marshalling got variable: %q", err)
				}

				wantMarshal, err := json.MarshalIndent(want, "", "    ")
				if err != nil {
					t.Fatalf("error marshalling want variable: %q", err)
				}

				t.Fatalf("got: %s != want: %s\n", gotMarshal, wantMarshal)
			}

		})
	}
}
