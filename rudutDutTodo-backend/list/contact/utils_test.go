package contact_test

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type mongouriJsonFile struct {
	DBPath        string `json:"dbPath"`
	TestDBURI     string `json:"testDB"`
	TodoListDBURI string `json:"todoListDB"`
}

func localJsonToStruct() (mongouriJsonFile, error) {
	mongoURIJsonFileContents, err := os.ReadFile("./mongouri.json")
	var mURIJ mongouriJsonFile

	if err != nil {
		return mURIJ,
			fmt.Errorf("having trouble opening json mongoURIJsonFileContents: %q\n", err)
	}

	err = json.Unmarshal(mongoURIJsonFileContents, &mURIJ)
	if err != nil {
		return mURIJ,
			fmt.Errorf("having trouble unmarshalling json mongoURIJsonFileContents: %q\n", err)
	}

	return mURIJ, nil
}

func spinUpMongoDB() error {
	mURIJ, err := localJsonToStruct()
	if err != nil {
		return err
	}

	var lineNumbers int
	var wg sync.WaitGroup

	wg.Add(1)
	spinUpErr := make(chan struct{})
	go func() {
		defer wg.Done()
		fmt.Printf("\t**spinning up instance on path: %s**\n", mURIJ.DBPath)
		cmd := exec.Command("mongod", "--dbpath", mURIJ.DBPath)

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
