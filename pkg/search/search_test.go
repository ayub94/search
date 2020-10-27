package search

import (
	"reflect"
	"testing"
)

func TestFindAnyMatchTextInFile(t *testing.T) {
	type args struct {
		phrase   string
		filetext string
	}
	tests := []struct {
		name    string
		args    args
		wantRes Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := FindAnyMatchTextInFile(tt.args.phrase, tt.args.filetext); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("FindAnyMatchTextInFile() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
