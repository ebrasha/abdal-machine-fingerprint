/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : errors.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Sentinel errors for the abdalmf public API.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package abdalmf

import "errors"

var (
	// ErrMachineIDNotFound is returned when the platform machine ID cannot be read or is empty.
	ErrMachineIDNotFound = errors.New("abdalmf: machine id not found")

	// ErrUnsupportedPlatform is returned on operating systems without a machine ID provider.
	ErrUnsupportedPlatform = errors.New("abdalmf: unsupported platform")

	// ErrUnsupportedAlgorithm is returned when the requested hash algorithm is unknown.
	ErrUnsupportedAlgorithm = errors.New("abdalmf: unsupported algorithm")

	// ErrUnsupportedEncoding is returned when the requested output encoding is unknown.
	ErrUnsupportedEncoding = errors.New("abdalmf: unsupported encoding")

	// ErrEmptyApplicationID is returned when a protected or HMAC call receives an empty application ID.
	ErrEmptyApplicationID = errors.New("abdalmf: empty application id")

	// ErrUnsupportedHMACAlgorithm is returned when HMAC cannot be computed with the requested algorithm.
	ErrUnsupportedHMACAlgorithm = errors.New("abdalmf: unsupported hmac algorithm")
)
