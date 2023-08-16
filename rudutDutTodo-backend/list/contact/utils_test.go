package contact_test

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/contact"
	"os"
	"os/exec"
	"strings"
	"sync"
	"testing"
)

type mongoHandlesJsonFile struct {
	DBPath string `json:"dbPath"`
	DBs    []struct {
		Info struct {
			URI     string `json:"uri"`
			Handles []struct {
				Name        string   `json:"name"`
				Collections []string `json:"collections"`
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

func spinUpMongoDB() error {
	mHJ, err := localJsonToStruct()
	if err != nil {
		return err
	}

	var lineNumbers int
	var wg sync.WaitGroup

	wg.Add(1)
	spinUpErr := make(chan struct{})
	go func() {
		defer wg.Done()
		fmt.Printf("\t**spinning up instance on path: %s**\n", mHJ.DBPath)
		cmd := exec.Command("mongod", "--dbpath", mHJ.DBPath)

		stdOut, _ := cmd.StdoutPipe()
		cmd.Start()

		scanner := bufio.NewScanner(stdOut)
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
			return errors.New("cannot find strings on `mongod` output")
		}
	default:
	}

	return nil
}

func setupEnvVar(t *testing.T, val string) {
	err := os.Setenv(contact.MONGO_URI_ENV_VAR, val)
	if err != nil {
		t.Fatalf("having trouble setting up env var %q\n", err)
	}
}
