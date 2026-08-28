/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : id_test.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Integration tests for AbdalID on supported platforms.
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
	"runtime"
	"testing"
)

func TestAbdalIDSupportedPlatform(t *testing.T) {
	switch runtime.GOOS {
	case "windows", "linux", "darwin":
	default:
		t.Skip("platform-specific machine ID test skipped")
	}

	id, err := AbdalID()
	if err != nil {
		t.Fatalf("AbdalID() error = %v", err)
	}
	if id == "" {
		t.Fatal("AbdalID() returned empty value")
	}
}

func TestAbdalHashSupportedPlatform(t *testing.T) {
	switch runtime.GOOS {
	case "windows", "linux", "darwin":
	default:
		t.Skip("platform-specific hash test skipped")
	}

	fingerprint, err := AbdalHash("SHA256")
	if err != nil {
		t.Fatalf("AbdalHash() error = %v", err)
	}
	if len(fingerprint) != 64 {
		t.Fatalf("AbdalHash() length = %d, want 64", len(fingerprint))
	}
}

func TestAbdalProtectedSupportedPlatform(t *testing.T) {
	switch runtime.GOOS {
	case "windows", "linux", "darwin":
	default:
		t.Skip("platform-specific protected test skipped")
	}

	fingerprint, err := AbdalProtected("abdal-security-tools", "SHA256")
	if err != nil {
		t.Fatalf("AbdalProtected() error = %v", err)
	}
	if len(fingerprint) != 64 {
		t.Fatalf("AbdalProtected() length = %d, want 64", len(fingerprint))
	}
}

func TestAbdalHMACEqualsProtected(t *testing.T) {
	switch runtime.GOOS {
	case "windows", "linux", "darwin":
	default:
		t.Skip("platform-specific HMAC test skipped")
	}

	protected, err := AbdalProtected("abdal-security-tools", "SHA512")
	if err != nil {
		t.Fatalf("AbdalProtected() error = %v", err)
	}

	hmacValue, err := AbdalHMAC("abdal-security-tools", "SHA512")
	if err != nil {
		t.Fatalf("AbdalHMAC() error = %v", err)
	}
	if protected != hmacValue {
		t.Fatalf("AbdalHMAC() and AbdalProtected() mismatch")
	}
}

func TestErrorsAreSentinelValues(t *testing.T) {
	if ErrMachineIDNotFound.Error() == "" {
		t.Fatal("ErrMachineIDNotFound has empty message")
	}
	if !errors.Is(ErrUnsupportedAlgorithm, ErrUnsupportedAlgorithm) {
		t.Fatal("ErrUnsupportedAlgorithm is not comparable with errors.Is")
	}
}
