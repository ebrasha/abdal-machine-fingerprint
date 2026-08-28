/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : protected_test.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Integration tests for protected and HMAC public APIs.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package abdalmf

import (
	"errors"
	"testing"
)

func TestAbdalProtectedEmptyApplicationID(t *testing.T) {
	_, err := AbdalProtected("", "SHA256")
	if !errors.Is(err, ErrEmptyApplicationID) {
		t.Fatalf("AbdalProtected() error = %v, want %v", err, ErrEmptyApplicationID)
	}
}

func TestAbdalHMACEmptyApplicationID(t *testing.T) {
	_, err := AbdalHMAC("", "SHA256")
	if !errors.Is(err, ErrEmptyApplicationID) {
		t.Fatalf("AbdalHMAC() error = %v, want %v", err, ErrEmptyApplicationID)
	}
}

func TestAbdalProtectedUnsupportedAlgorithm(t *testing.T) {
	_, err := AbdalProtected("abdal-security-tools", "SHA999")
	if !errors.Is(err, ErrUnsupportedAlgorithm) {
		t.Fatalf("AbdalProtected() error = %v, want %v", err, ErrUnsupportedAlgorithm)
	}
}

func TestAbdalEncodeUnsupportedEncoding(t *testing.T) {
	_, err := AbdalEncode("abdal-security-tools", "SHA256", "BINARY")
	if !errors.Is(err, ErrUnsupportedEncoding) {
		t.Fatalf("AbdalEncode() error = %v, want %v", err, ErrUnsupportedEncoding)
	}
}

func TestAbdalHashUnsupportedAlgorithm(t *testing.T) {
	_, err := AbdalHash("SHA999")
	if !errors.Is(err, ErrUnsupportedAlgorithm) {
		t.Fatalf("AbdalHash() error = %v, want %v", err, ErrUnsupportedAlgorithm)
	}
}
