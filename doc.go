/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : doc.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Package documentation and privacy notes for abdalmf.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

// Package abdalmf provides a cross-platform API for stable machine identity
// and cryptographically derived machine fingerprints.
//
// AbdalID returns the raw operating-system machine identifier. Treat it as
// confidential; storage, display, and transmission are the developer's
// responsibility.
//
// AbdalHash and AbdalProtected offer safer alternatives when the raw machine
// ID must not be stored or shared directly. AbdalProtected and AbdalHMAC derive
// an application-specific fingerprint using HMAC or keyed BLAKE2, keyed by the
// normalized machine ID.
//
// The library is fully offline. It performs no network requests, telemetry,
// analytics, or automatic logging.
//
// Supported platforms: Windows, Linux, macOS.
package abdalmf
