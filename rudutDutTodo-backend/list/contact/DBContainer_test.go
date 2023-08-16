package contact_test

import (
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/contact"
	"reflect"
	"testing"
)

func TestNewDBContainer(t *testing.T) {
	type args struct {
		DBName         string
		collectionName string
	}
	tests := []struct {
		name        string
		args        args
		want        *contact.DBContainer
		wantErr     bool
		expectedErr error
	}{
		{
			name: "initial",
			args: args{
				DBName:         "test",
				collectionName: "test-date-docs",
			},
			want: &contact.DBContainer{
				DBName:         "test",
				CollectionName: "test-date-docs",
			},
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name: "dbname is misssing",
			args: args{
				DBName:         "",
				collectionName: "test-date-docs",
			},
			want: &contact.DBContainer{
				DBName:         "test",
				CollectionName: "test-date-docs",
			},
			wantErr:     true,
			expectedErr: contact.MissingDBInfoErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := contact.NewDBContainer(tt.args.DBName, tt.args.collectionName)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected err: %q, but didn't get one", tt.expectedErr)
				}

				if err != tt.expectedErr {
					t.Fatalf("expected err: %q, but got different err; %q", tt.expectedErr, err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDBContainer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
