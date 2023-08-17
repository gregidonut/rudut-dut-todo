package contact_test

import (
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/contact"
	"strings"
	"testing"
)

func TestPingMongo(t *testing.T) {
	mongoHandles, err := localJsonToStruct()
	if err != nil {
		t.Fatalf("having trouble setting up mongoHandles: %q", err)
	}

	type args struct {
		uri contact.MongoURI
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		expectedErr error
	}{
		{
			name: "ping testDB",
			args: args{
				uri: contact.MongoURI(mongoHandles.DBs[0].Info.URI),
			},
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name: "ping todolistDB",
			args: args{
				uri: contact.MongoURI(mongoHandles.DBs[1].Info.URI),
			},
			wantErr:     false,
			expectedErr: nil,
		},
		{
			// just to trigger an error, implicitly declaring an incorrect uri
			name: "wrong uri",
			args: args{
				uri: contact.MongoURI("hello"),
			},
			wantErr:     true,
			expectedErr: contact.MongoPingErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// client return val here is unnecessary since ping just returns it and
			// does nothing to it anyway
			_, err := contact.PingMongo(tt.args.uri)
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
		})
	}
}
