/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : encode.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Output encoding helpers for fingerprint strings.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package encode

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"
)

// Encoding identifies a supported output encoding format.
type Encoding string

const (
	EncodingHex       Encoding = "HEX"
	EncodingHexUpper  Encoding = "HEX-UPPER"
	EncodingBase64    Encoding = "BASE64"
	EncodingBase64URL Encoding = "BASE64URL"
)

// ErrUnsupportedEncoding indicates that the requested encoding is unknown.
var ErrUnsupportedEncoding = errors.New("encode: unsupported encoding")

// ParseEncoding normalizes and resolves a user-provided encoding string.
func ParseEncoding(name string) (Encoding, error) {
	key := normalizeEncodingName(name)
	if key == "" {
		return "", ErrUnsupportedEncoding
	}

	encoding, ok := encodingAliases[key]
	if !ok {
		return "", ErrUnsupportedEncoding
	}
	return encoding, nil
}

func normalizeEncodingName(name string) string {
	normalized := strings.TrimSpace(name)
	normalized = strings.ToUpper(normalized)
	normalized = strings.NewReplacer("-", "", "_", "", " ", "").Replace(normalized)
	return normalized
}

var encodingAliases = map[string]Encoding{
	"HEX":       EncodingHex,
	"HEXUPPER":  EncodingHexUpper,
	"BASE64":    EncodingBase64,
	"BASE64URL": EncodingBase64URL,
}

// Bytes encodes digest bytes using the requested output format.
func Bytes(data []byte, encoding Encoding) (string, error) {
	switch encoding {
	case EncodingHex:
		return hex.EncodeToString(data), nil
	case EncodingHexUpper:
		return strings.ToUpper(hex.EncodeToString(data)), nil
	case EncodingBase64:
		return base64.StdEncoding.EncodeToString(data), nil
	case EncodingBase64URL:
		return base64.RawURLEncoding.EncodeToString(data), nil
	default:
		return "", ErrUnsupportedEncoding
	}
}
