/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : algorithm_test.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Unit tests for algorithm name normalization and parsing.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package hash

import (
	"errors"
	"testing"
)

func TestParseAlgorithm(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Algorithm
		wantErr error
	}{
		{name: "sha256", input: "SHA256", want: AlgoSHA256},
		{name: "sha256 lower", input: "sha256", want: AlgoSHA256},
		{name: "sha256 dashed", input: "SHA-256", want: AlgoSHA256},
		{name: "sha256 underscored", input: "SHA_256", want: AlgoSHA256},
		{name: "sha3 dashed", input: "SHA3-512", want: AlgoSHA3512},
		{name: "sha3 compact", input: "SHA3512", want: AlgoSHA3512},
		{name: "sha3 underscored", input: "sha3_512", want: AlgoSHA3512},
		{name: "blake2b", input: "BLAKE2B-512", want: AlgoBLAKE2B512},
		{name: "invalid", input: "SHA999", wantErr: ErrUnsupportedAlgorithm},
		{name: "empty", input: "   ", wantErr: ErrUnsupportedAlgorithm},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAlgorithm(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("ParseAlgorithm() error = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseAlgorithm() unexpected error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("ParseAlgorithm() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSupportsProtected(t *testing.T) {
	if !SupportsProtected(AlgoSHA256) {
		t.Fatal("SHA256 should support protected fingerprints")
	}
	if !SupportsProtected(AlgoBLAKE2B512) {
		t.Fatal("BLAKE2B-512 should support protected fingerprints")
	}
}
