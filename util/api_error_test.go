package util

import (
	"testing"
)

func TestNewApiError(t *testing.T) {
	cases := []struct{
		label    string
		wantErr  bool // the test if err is nil
		wantMsg  string
		wantCode int
	}{
		{"ShouldReturnApiError", false, "Some error Message", 400},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			err := NewApiError(tc.wantMsg, 400)
			isNil := err == nil
			if isNil != tc.wantErr {
				t.Fatalf("Wrong err returned. Want: %v got: %v\n", tc.wantErr, isNil)
			}

			if err.StatusCode != tc.wantCode {
				t.Fatalf("Wrong code returned. Want: %v got: %v\n", tc.wantCode, err.StatusCode)
			}

			if err.Message != tc.wantMsg {
				t.Fatalf("Wrong msg returned. Want: %v got: %v\n", tc.wantMsg, err.Message)
			}
		})
	}
}
