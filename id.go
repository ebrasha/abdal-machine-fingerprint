/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : id.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Public API for retrieving the raw normalized machine ID.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package abdalmf

import (
	"github.com/ebrasha/abdal-machine-fingerprint/internal/machine"
)

// AbdalID returns the normalized raw machine identifier for the current host.
// The value is confidential and should be handled carefully by the caller.
func AbdalID() (string, error) {
	id, err := machine.NormalizedID()
	if err != nil {
		return "", mapMachineIDError(err)
	}
	return id, nil
}

func mapMachineIDError(err error) error {
	switch {
	case machine.IsUnsupportedPlatform(err):
		return ErrUnsupportedPlatform
	case machine.IsNotFound(err):
		return ErrMachineIDNotFound
	default:
		return ErrMachineIDNotFound
	}
}
