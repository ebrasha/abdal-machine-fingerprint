/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : protected.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Public API for protected and HMAC machine fingerprints.
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
	"strings"

	"github.com/ebrasha/abdal-machine-fingerprint/core/encode"
	"github.com/ebrasha/abdal-machine-fingerprint/core/hash"
	"github.com/ebrasha/abdal-machine-fingerprint/core/machine"
)

// AbdalProtected returns a lowercase hexadecimal protected fingerprint for the application.
func AbdalProtected(applicationID, algorithm string) (string, error) {
	return abdalProtectedWithEncoding(applicationID, algorithm, encode.EncodingHex)
}

// AbdalHMAC returns a lowercase hexadecimal HMAC-style protected fingerprint.
func AbdalHMAC(applicationID, algorithm string) (string, error) {
	return abdalProtectedWithEncoding(applicationID, algorithm, encode.EncodingHex)
}

func abdalProtectedWithEncoding(applicationID, algorithm string, outputEncoding encode.Encoding) (string, error) {
	if strings.TrimSpace(applicationID) == "" {
		return "", ErrEmptyApplicationID
	}

	algo, err := hash.ParseAlgorithm(algorithm)
	if err != nil {
		return "", ErrUnsupportedAlgorithm
	}
	if !hash.SupportsProtected(algo) {
		return "", ErrUnsupportedHMACAlgorithm
	}

	id, err := machine.NormalizedID()
	if err != nil {
		return "", mapMachineIDError(err)
	}

	digest, err := hash.Protected(applicationID, id, algo)
	if err != nil {
		if errors.Is(err, hash.ErrUnsupportedHMACAlgorithm) {
			return "", ErrUnsupportedHMACAlgorithm
		}
		return "", err
	}

	encoded, err := encode.Bytes(digest, outputEncoding)
	if err != nil {
		return "", ErrUnsupportedEncoding
	}
	return encoded, nil
}
