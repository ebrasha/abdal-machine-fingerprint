/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : abdalconstants.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 19:52:22
 * Description  : Single source of truth for root public package constants.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package abdalmf

// Version is the current runtime version of the abdalmf package (SemVer).
// Git release tags MUST match this value with a v prefix (e.g. v1.0.0).
const Version = "1.0.0"

// Default hash and encoding values used by the public API.
const (
	DefaultHashAlgorithm      = AlgorithmSHA256
	DefaultProtectedAlgorithm = AlgorithmSHA256
	DefaultEncoding           = EncodingHex
)

// Supported hash algorithm identifiers for use with AbdalHash, AbdalProtected, and AbdalHMAC.
const (
	AlgorithmMD5        = "MD5"
	AlgorithmSHA1       = "SHA1"
	AlgorithmSHA224     = "SHA224"
	AlgorithmSHA256     = "SHA256"
	AlgorithmSHA384     = "SHA384"
	AlgorithmSHA512     = "SHA512"
	AlgorithmSHA512224  = "SHA512-224"
	AlgorithmSHA512256  = "SHA512-256"
	AlgorithmSHA3224    = "SHA3-224"
	AlgorithmSHA3256    = "SHA3-256"
	AlgorithmSHA3384    = "SHA3-384"
	AlgorithmSHA3512    = "SHA3-512"
	AlgorithmBLAKE2B256 = "BLAKE2B-256"
	AlgorithmBLAKE2B384 = "BLAKE2B-384"
	AlgorithmBLAKE2B512 = "BLAKE2B-512"
	AlgorithmBLAKE2S256 = "BLAKE2S-256"
)

// Supported output encoding identifiers for use with AbdalEncode.
const (
	EncodingHex       = "HEX"
	EncodingHexUpper  = "HEX-UPPER"
	EncodingBase64    = "BASE64"
	EncodingBase64URL = "BASE64URL"
)
