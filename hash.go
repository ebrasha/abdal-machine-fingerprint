/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : hash.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Public API for hashing the normalized machine identifier.
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

	"github.com/ebrasha/abdal-machine-fingerprint/internal/encode"
	"github.com/ebrasha/abdal-machine-fingerprint/internal/hash"
	"github.com/ebrasha/abdal-machine-fingerprint/internal/machine"
)

// AbdalHash returns a lowercase hexadecimal hash of the normalized machine ID.
func AbdalHash(algorithm string) (string, error) {
	algo, err := hash.ParseAlgorithm(algorithm)
	if err != nil {
		return "", ErrUnsupportedAlgorithm
	}

	id, err := machine.NormalizedID()
	if err != nil {
		return "", mapMachineIDError(err)
	}

	digest, err := hash.Digest([]byte(id), algo)
	if err != nil {
		if errors.Is(err, hash.ErrUnsupportedAlgorithm) {
			return "", ErrUnsupportedAlgorithm
		}
		return "", err
	}

	encoded, err := encode.Bytes(digest, encode.Encoding(DefaultEncoding))
	if err != nil {
		return "", ErrUnsupportedEncoding
	}
	return encoded, nil
}
