/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : hmac_test.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Unit tests for protected HMAC and keyed BLAKE2 fingerprints.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"testing"
)

func TestProtectedSHA256ReferenceCompatibility(t *testing.T) {
	appID := "ms.azur.appX"
	machineID := "1a1238d601ad430cbea7efb0d1f3d92d"

	digest, err := Protected(appID, machineID, AlgoSHA256)
	if err != nil {
		t.Fatalf("Protected() error = %v", err)
	}

	mac := hmac.New(sha256.New, []byte(machineID))
	mac.Write([]byte(appID))
	expected := mac.Sum(nil)

	if !hmac.Equal(digest, expected) {
		t.Fatalf("Protected() digest mismatch, got %s want %s", hex.EncodeToString(digest), hex.EncodeToString(expected))
	}
}

func TestProtectedKeyedBLAKE2(t *testing.T) {
	digest, err := Protected("abdal-security-tools", "machine-id", AlgoBLAKE2B256)
	if err != nil {
		t.Fatalf("Protected() error = %v", err)
	}
	if len(digest) != 32 {
		t.Fatalf("Protected() BLAKE2B-256 length = %d, want 32", len(digest))
	}
}

func TestProtectedUnsupportedAlgorithm(t *testing.T) {
	_, err := Protected("app", "machine", Algorithm("SHA999"))
	if !errors.Is(err, ErrUnsupportedHMACAlgorithm) {
		t.Fatalf("Protected() error = %v, want %v", err, ErrUnsupportedHMACAlgorithm)
	}
}

func TestProtectedDeterministic(t *testing.T) {
	first, err := Protected("abdal-security-tools", "machine-id", AlgoSHA512)
	if err != nil {
		t.Fatalf("Protected() first call error = %v", err)
	}
	second, err := Protected("abdal-security-tools", "machine-id", AlgoSHA512)
	if err != nil {
		t.Fatalf("Protected() second call error = %v", err)
	}
	if hex.EncodeToString(first) != hex.EncodeToString(second) {
		t.Fatal("Protected() is not deterministic")
	}
}
