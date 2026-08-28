/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : abdalconstants_test.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 19:52:22
 * Description  : Tests for root package constants and version metadata.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package abdalmf

import (
	"testing"

	"github.com/ebrasha/abdal-machine-fingerprint/internal/encode"
	"github.com/ebrasha/abdal-machine-fingerprint/internal/hash"
)

func TestVersion(t *testing.T) {
	if Version != "1.0.0" {
		t.Fatalf("Version = %q, want %q", Version, "1.0.0")
	}
}

func TestDefaultConstants(t *testing.T) {
	if DefaultHashAlgorithm != AlgorithmSHA256 {
		t.Fatalf("DefaultHashAlgorithm = %q, want %q", DefaultHashAlgorithm, AlgorithmSHA256)
	}
	if DefaultProtectedAlgorithm != AlgorithmSHA256 {
		t.Fatalf("DefaultProtectedAlgorithm = %q, want %q", DefaultProtectedAlgorithm, AlgorithmSHA256)
	}
	if DefaultEncoding != EncodingHex {
		t.Fatalf("DefaultEncoding = %q, want %q", DefaultEncoding, EncodingHex)
	}
}

func TestAlgorithmConstantsMatchInternalParser(t *testing.T) {
	algo, err := hash.ParseAlgorithm(AlgorithmSHA256)
	if err != nil {
		t.Fatalf("ParseAlgorithm() error = %v", err)
	}
	if string(algo) != AlgorithmSHA256 {
		t.Fatalf("ParseAlgorithm() = %q, want %q", algo, AlgorithmSHA256)
	}
}

func TestEncodingConstantsMatchInternalParser(t *testing.T) {
	enc, err := encode.ParseEncoding(EncodingHex)
	if err != nil {
		t.Fatalf("ParseEncoding() error = %v", err)
	}
	if string(enc) != EncodingHex {
		t.Fatalf("ParseEncoding() = %q, want %q", enc, EncodingHex)
	}
}
