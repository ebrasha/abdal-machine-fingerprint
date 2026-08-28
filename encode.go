/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : encode.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Public API for protected fingerprints with custom encodings.
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
)

// AbdalEncode returns a protected fingerprint using the requested output encoding.
func AbdalEncode(applicationID, algorithm, encoding string) (string, error) {
	parsedEncoding, err := encode.ParseEncoding(encoding)
	if err != nil {
		if errors.Is(err, encode.ErrUnsupportedEncoding) {
			return "", ErrUnsupportedEncoding
		}
		return "", err
	}
	return abdalProtectedWithEncoding(applicationID, algorithm, parsedEncoding)
}
