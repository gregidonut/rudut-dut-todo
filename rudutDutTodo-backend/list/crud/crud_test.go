package crud_test

import (
	"log"
	"os"
	"testing"
)

// package scope setup and teardown
func TestMain(m *testing.M) {
	mongoProcess, err := spinUpMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	exitVal := m.Run()

	mongoProcess.Kill()

	os.Exit(exitVal)
}
