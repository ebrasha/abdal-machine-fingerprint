/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : encode_test.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Unit tests for fingerprint output encodings.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package encode

import (
	"errors"
	"testing"
)

func TestParseEncoding(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Encoding
		wantErr error
	}{
		{name: "hex lower", input: "hex", want: EncodingHex},
		{name: "hex upper", input: "HEX", want: EncodingHex},
		{name: "hex upper dashed", input: "hex-upper", want: EncodingHexUpper},
		{name: "base64", input: "BASE64", want: EncodingBase64},
		{name: "base64url", input: "base64url", want: EncodingBase64URL},
		{name: "invalid", input: "binary", wantErr: ErrUnsupportedEncoding},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseEncoding(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("ParseEncoding() error = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseEncoding() unexpected error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("ParseEncoding() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBytesEncodings(t *testing.T) {
	data := []byte{0x90, 0x4f, 0x89, 0x5a}

	hexValue, err := Bytes(data, EncodingHex)
	if err != nil {
		t.Fatalf("Bytes(HEX) error = %v", err)
	}
	if hexValue != "904f895a" {
		t.Fatalf("Bytes(HEX) = %q, want %q", hexValue, "904f895a")
	}

	hexUpper, err := Bytes(data, EncodingHexUpper)
	if err != nil {
		t.Fatalf("Bytes(HEX-UPPER) error = %v", err)
	}
	if hexUpper != "904F895A" {
		t.Fatalf("Bytes(HEX-UPPER) = %q, want %q", hexUpper, "904F895A")
	}

	base64Value, err := Bytes(data, EncodingBase64)
	if err != nil {
		t.Fatalf("Bytes(BASE64) error = %v", err)
	}
	if base64Value != "kE+JWg==" {
		t.Fatalf("Bytes(BASE64) = %q, want %q", base64Value, "kE+JWg==")
	}

	base64URL, err := Bytes(data, EncodingBase64URL)
	if err != nil {
		t.Fatalf("Bytes(BASE64URL) error = %v", err)
	}
	if base64URL != "kE-JWg" {
		t.Fatalf("Bytes(BASE64URL) = %q, want %q", base64URL, "kE-JWg")
	}
}
