package contact_test

import (
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/contact"
	"os"
	"testing"
)

func TestGetMongoUriFromEnv(t *testing.T) {
	mongoURIs, err := setupURIs()
	if err != nil {
		t.Fatalf("having trouble setting up mongoURIs: %q", err)
	}
	testDBURI := contact.MongoURI(mongoURIs.TestDBURI)
	todoListDBURI := contact.MongoURI(mongoURIs.TodoListDBURI)

	tests := []struct {
		name        string
		setupEnvVar bool
		envVarValue string
		want        *contact.MongoURI
		wantErr     bool
		expectedErr error
	}{
		{
			name:        "env var is not declared",
			setupEnvVar: false,
			envVarValue: "",
			want:        nil,
			wantErr:     true,
			expectedErr: contact.MongoEnvVarNotDeclaredErr,
		},
		{
			name:        "env var is local testDB",
			setupEnvVar: true,
			envVarValue: mongoURIs.TestDBURI,
			want:        &testDBURI,
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name:        "env var is local todoListDB",
			setupEnvVar: true,
			envVarValue: mongoURIs.TodoListDBURI,
			want:        &todoListDBURI,
			wantErr:     false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupEnvVar {
				err := os.Setenv(contact.MONGO_URI_ENV_VAR, tt.envVarValue)
				if err != nil {
					t.Fatalf("having trouble setting up env var %q\n", err)
				}
			}
			got, err := contact.GetMongoUriFromEnv()
			if tt.wantErr {
				if err == nil {
					t.Fatalf(
						"wanted err: %q, but didn't get one\n",
						tt.expectedErr,
					)
				}

				if err != tt.expectedErr {
					t.Fatalf(
						"expected error = %q, but got different error %q",
						tt.expectedErr,
						err,
					)
				}
				return
			}
			if *got != *tt.want {
				t.Errorf("GetMongoUriFromEnv() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMongoURI_ToString(t *testing.T) {
	tests := []struct {
		name string
		m    contact.MongoURI
		want string
	}{
		{
			name: "initial",
			m:    contact.MongoURI("hello world"),
			want: "hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ToString(); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
