package crud_test

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/contact"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/crud"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/todo"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestBsonMToTodo(t *testing.T) {
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
		t.Errorf("having trouble parsing date %q\n", dateFromTestObj)
	}

	var testObjAsBsonM bson.M
	err = bson.UnmarshalExtJSON([]byte(testObj), false, &testObjAsBsonM)
	if err != nil {
		t.Fatalf("bson.UnmarshalExtJSON(): %v\n", err)
	}

	type args struct {
		bsonM bson.M
	}
	tests := []struct {
		name        string
		args        args
		want        todo.Todo
		wantErr     bool
		expectedErr error
	}{
		{
			name: "initial",
			args: args{
				bsonM: testObjAsBsonM,
			},
			want: todo.Todo{
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
			wantErr:     false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := crud.BsonMToTodo(tt.args.bsonM)
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
				t.Errorf("BsonMToTodo() got = %v, want %v", got, tt.want)
			}

		})
	}
}

type mongoHandlesJsonFile struct {
	DBPath string `json:"dbPath"`
	DBs    []struct {
		Info struct {
			URI     string `json:"uri"`
			Handles []struct {
				Name        string `json:"name"`
				Collections []struct {
					CName string `json:"cName"`
					Items []struct {
						ID struct {
							OID string `json:"$oid"`
						} `json:"_id"`
						Content string `json:"content"`
						Date    struct {
							JSDateObject string `json:"$date"`
						} `json:"date"`
						PostID   int `json:"postId"`
						Progress struct {
							Todo       bool `json:"todo"`
							InProgress bool `json:"inProgress"`
							Finished   bool `json:"finished"`
						} `json:"progress"`
						Title string `json:"title"`
					} `json:"items"`
				} `json:"collections"`
			} `json:"handles"`
		} `json:"dbInfo"`
	} `json:"dbs"`
}

func localJsonToStruct() (mongoHandlesJsonFile, error) {
	mongoURIJsonFileContents, err := os.ReadFile("./mongohandles.json")
	var mHJ mongoHandlesJsonFile

	if err != nil {
		return mHJ,
			fmt.Errorf("having trouble opening json mongoURIJsonFileContents: %q\n", err)
	}

	err = json.Unmarshal(mongoURIJsonFileContents, &mHJ)
	if err != nil {
		return mHJ,
			fmt.Errorf("having trouble unmarshalling json mongoURIJsonFileContents: %q\n", err)
	}

	return mHJ, nil
}

func spinUpMongoDB() (*os.Process, error) {
	var mongoProcess *os.Process

	mHJ, err := localJsonToStruct()
	if err != nil {
		return &os.Process{}, err
	}

	var lineNumbers int
	var wg sync.WaitGroup

	wg.Add(1)
	spinUpErr := make(chan struct{})
	go func() {
		defer wg.Done()

		// wait for a possible concurrent test is using the db to finish
		//{{
		for {
			pidOfCmd := exec.Command("pidof", "mongod")
			err = pidOfCmd.Run()
			if err != nil {
				// this would probably mean that a mongod process isn't running
				// in which case we should break so this goroutine can't spin it up
				break
			}
		}
		//}}

		// check if a mongoDB Process is up
		pGrepCmd := exec.Command("pgrep", "mongod")
		stdErr, _ := pGrepCmd.StderrPipe()
		pGrepCmd.Start()
		scanner := bufio.NewScanner(stdErr)
		if scanner.Scan() {
			time.Sleep(2 * time.Second)
		}

		fmt.Printf("\t**spinning up instance on path: %s**\n", mHJ.DBPath)
		cmd := exec.Command("mongod", "--dbpath", mHJ.DBPath)

		stdOut, _ := cmd.StdoutPipe()
		cmd.Start()

		mongoProcess = cmd.Process

		scanner = bufio.NewScanner(stdOut)
		for scanner.Scan() {
			lineNumbers++
			if lineNumbers > 30 {
				close(spinUpErr)
				return
			}
			m := scanner.Text()
			if strings.Contains(m, "ctx") && strings.Contains(m, "Listening on") {
				// TODO: implement a struct for hardcoded strings
				return
			}
		}
	}()
	wg.Wait()

	select {
	case _, ok := <-spinUpErr:
		if !ok {
			return mongoProcess, errors.New("cannot find strings on `mongod` output")
		}
	default:
	}

	return mongoProcess, nil
}

func setupEnvVar(t *testing.T, val string) {
	err := os.Setenv(contact.MONGO_URI_ENV_VAR, val)
	if err != nil {
		t.Fatalf("having trouble setting up env var %q\n", err)
	}
}
