package errxit_test

import (
	"errors"
	"testing"

	"git.sr.ht/~jamesponddotco/errxit-go"
)

var errGeneric = errors.New("generic error")

func TestError_New(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		giveError error
		giveCode  int
		wantErr   string
		wantCode  int
	}{
		{
			name:      "Valid error and code",
			giveError: errGeneric,
			giveCode:  1,
			wantErr:   "1: generic error",
			wantCode:  1,
		},
		{
			name:      "Nil error and valid code",
			giveError: nil,
			giveCode:  2,
			wantErr:   "undefined error",
			wantCode:  2,
		},
		{
			name:      "Zero values",
			giveError: nil,
			giveCode:  0,
			wantErr:   "undefined error",
			wantCode:  0,
		},
		{
			name:      "Negative code",
			giveError: errGeneric,
			giveCode:  -1,
			wantErr:   "-1: generic error",
			wantCode:  -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := errxit.New(tt.giveError, tt.giveCode)

			if got.Error() != tt.wantErr {
				t.Errorf("New(%v, %d).Error() = %q, want %q", tt.giveError, tt.giveCode, got.Error(), tt.wantErr)
			}

			if got.Code() != tt.wantCode {
				t.Errorf("New(%v, %d).Code() = %d, want %d", tt.giveError, tt.giveCode, got.Code(), tt.wantCode)
			}
		})
	}
}

func TestError_Unwrap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		giveError error
		giveCode  int
		wantErr   error
	}{
		{
			name:      "With underlying error",
			giveError: errGeneric,
			giveCode:  1,
			wantErr:   errGeneric,
		},
		{
			name:      "With nil underlying error",
			giveError: nil,
			giveCode:  2,
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var (
				err       = errxit.New(tt.giveError, tt.giveCode)
				unwrapped = err.Unwrap()
			)

			if !errors.Is(unwrapped, tt.wantErr) {
				t.Fatalf("Unwrap() = %v, want %v", unwrapped, tt.wantErr)
			}
		})
	}
}
