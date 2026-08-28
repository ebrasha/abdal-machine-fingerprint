/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : normalize.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Deterministic normalization for platform machine identifiers.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package normalize

import (
	"strings"

	"github.com/ebrasha/abdal-machine-fingerprint/core/platform"
)

// ErrEmpty is returned when normalization produces an empty machine ID.
var ErrEmpty = platform.ErrNotFound

// MachineID trims surrounding whitespace and line endings from a raw machine ID.
func MachineID(raw string) (string, error) {
	normalized := strings.TrimSpace(raw)
	normalized = strings.Trim(normalized, "\r\n")
	normalized = strings.TrimSpace(normalized)
	if normalized == "" {
		return "", ErrEmpty
	}
	return normalized, nil
}
