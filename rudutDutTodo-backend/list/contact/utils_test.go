package contact_test

import (
	"encoding/json"
	"fmt"
	"os"
)

type mongoURIJson struct {
	TestDBURI     string `json:"testDB"`
	TodoListDBURI string `json:"todoListDB"`
}

func setupURIs() (mongoURIJson, error) {
	mongoURIJsonFileContents, err := os.ReadFile("./mongouri.json")
	var mURIJ mongoURIJson

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
