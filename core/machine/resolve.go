/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : resolve.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Shared machine ID resolution for public API functions.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package machine

import (
	"errors"

	"github.com/ebrasha/abdal-machine-fingerprint/core/normalize"
	"github.com/ebrasha/abdal-machine-fingerprint/core/platform"
)

// NormalizedID reads and normalizes the host machine identifier.
func NormalizedID() (string, error) {
	raw, err := platform.ReadRawMachineID()
	if err != nil {
		return "", err
	}
	return normalize.MachineID(raw)
}

// IsNotFound reports whether an error indicates a missing machine ID.
func IsNotFound(err error) bool {
	return errors.Is(err, platform.ErrNotFound) || errors.Is(err, normalize.ErrEmpty)
}

// IsUnsupportedPlatform reports whether the host OS is unsupported.
func IsUnsupportedPlatform(err error) bool {
	return errors.Is(err, platform.ErrUnsupported)
}
