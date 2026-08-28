/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : digest_test.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Unit tests for plain digest operations.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"testing"
)

func TestDigestSHA256KnownVector(t *testing.T) {
	input := []byte("abdalmf-test-input")
	expected := sha256.Sum256(input)

	digest, err := Digest(input, AlgoSHA256)
	if err != nil {
		t.Fatalf("Digest() error = %v", err)
	}
	if hex.EncodeToString(digest) != hex.EncodeToString(expected[:]) {
		t.Fatalf("Digest() mismatch for SHA256")
	}
}

func TestDigestUnsupportedAlgorithm(t *testing.T) {
	_, err := Digest([]byte("data"), Algorithm("SHA999"))
	if !errors.Is(err, ErrUnsupportedAlgorithm) {
		t.Fatalf("Digest() error = %v, want %v", err, ErrUnsupportedAlgorithm)
	}
}

func TestDigestAllAlgorithms(t *testing.T) {
	algorithms := []Algorithm{
		AlgoMD5,
		AlgoSHA1,
		AlgoSHA224,
		AlgoSHA256,
		AlgoSHA384,
		AlgoSHA512,
		AlgoSHA512224,
		AlgoSHA512256,
		AlgoSHA3224,
		AlgoSHA3256,
		AlgoSHA3384,
		AlgoSHA3512,
		AlgoBLAKE2B256,
		AlgoBLAKE2B384,
		AlgoBLAKE2B512,
		AlgoBLAKE2S256,
	}

	for _, algo := range algorithms {
		t.Run(string(algo), func(t *testing.T) {
			digest, err := Digest([]byte("machine-id"), algo)
			if err != nil {
				t.Fatalf("Digest() error = %v", err)
			}
			if len(digest) == 0 {
				t.Fatal("Digest() returned empty output")
			}
		})
	}
}
