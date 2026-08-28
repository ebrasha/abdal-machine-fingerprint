/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : errors.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Platform-layer sentinel errors for machine ID retrieval.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package platform

import "errors"

var (
	// ErrNotFound indicates that no machine ID could be read from the host OS.
	ErrNotFound = errors.New("platform: machine id not found")

	// ErrUnsupported indicates that the current operating system is not supported.
	ErrUnsupported = errors.New("platform: unsupported platform")
)
