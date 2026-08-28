/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : normalize_test.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Unit tests for machine ID normalization behavior.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package normalize

import (
	"errors"
	"testing"
)

func TestMachineID(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{
			name:  "plain value",
			input: "abc123",
			want:  "abc123",
		},
		{
			name:  "leading trailing spaces",
			input: "  abc123  ",
			want:  "abc123",
		},
		{
			name:  "trailing lf",
			input: "abc123\n",
			want:  "abc123",
		},
		{
			name:  "trailing crlf",
			input: "abc123\r\n",
			want:  "abc123",
		},
		{
			name:  "combined whitespace",
			input: " \r\nabc123\n ",
			want:  "abc123",
		},
		{
			name:    "empty",
			input:   "",
			wantErr: ErrEmpty,
		},
		{
			name:    "whitespace only",
			input:   " \r\n ",
			wantErr: ErrEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MachineID(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("MachineID() error = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("MachineID() unexpected error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("MachineID() = %q, want %q", got, tt.want)
			}
		})
	}
}
