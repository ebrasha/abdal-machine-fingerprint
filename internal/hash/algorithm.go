/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : algorithm.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Algorithm name normalization and parsing helpers.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package hash

import (
	"errors"
	"strings"
)

// Algorithm identifies a supported hash algorithm by its canonical name.
type Algorithm string

// Canonical algorithm identifiers. Values mirror abdalmf/abdalconstants.go.
const (
	AlgoMD5        Algorithm = "MD5"
	AlgoSHA1       Algorithm = "SHA1"
	AlgoSHA224     Algorithm = "SHA224"
	AlgoSHA256     Algorithm = "SHA256"
	AlgoSHA384     Algorithm = "SHA384"
	AlgoSHA512     Algorithm = "SHA512"
	AlgoSHA512224  Algorithm = "SHA512-224"
	AlgoSHA512256  Algorithm = "SHA512-256"
	AlgoSHA3224    Algorithm = "SHA3-224"
	AlgoSHA3256    Algorithm = "SHA3-256"
	AlgoSHA3384    Algorithm = "SHA3-384"
	AlgoSHA3512    Algorithm = "SHA3-512"
	AlgoBLAKE2B256 Algorithm = "BLAKE2B-256"
	AlgoBLAKE2B384 Algorithm = "BLAKE2B-384"
	AlgoBLAKE2B512 Algorithm = "BLAKE2B-512"
	AlgoBLAKE2S256 Algorithm = "BLAKE2S-256"
)

// ErrUnsupportedAlgorithm indicates that the requested algorithm is unknown.
var ErrUnsupportedAlgorithm = errors.New("hash: unsupported algorithm")

// ParseAlgorithm normalizes and resolves a user-provided algorithm string.
func ParseAlgorithm(name string) (Algorithm, error) {
	key := normalizeAlgorithmName(name)
	if key == "" {
		return "", ErrUnsupportedAlgorithm
	}

	algo, ok := algorithmAliases[key]
	if !ok {
		return "", ErrUnsupportedAlgorithm
	}
	return algo, nil
}

func normalizeAlgorithmName(name string) string {
	normalized := strings.TrimSpace(name)
	normalized = strings.ToUpper(normalized)
	normalized = strings.NewReplacer("-", "", "_", "", " ", "").Replace(normalized)
	return normalized
}

var algorithmAliases = map[string]Algorithm{
	normalizeAlgorithmName(string(AlgoMD5)):        AlgoMD5,
	normalizeAlgorithmName(string(AlgoSHA1)):       AlgoSHA1,
	normalizeAlgorithmName(string(AlgoSHA224)):     AlgoSHA224,
	normalizeAlgorithmName(string(AlgoSHA256)):     AlgoSHA256,
	normalizeAlgorithmName(string(AlgoSHA384)):     AlgoSHA384,
	normalizeAlgorithmName(string(AlgoSHA512)):     AlgoSHA512,
	normalizeAlgorithmName(string(AlgoSHA512224)):  AlgoSHA512224,
	normalizeAlgorithmName(string(AlgoSHA512256)):  AlgoSHA512256,
	normalizeAlgorithmName(string(AlgoSHA3224)):    AlgoSHA3224,
	normalizeAlgorithmName(string(AlgoSHA3256)):    AlgoSHA3256,
	normalizeAlgorithmName(string(AlgoSHA3384)):    AlgoSHA3384,
	normalizeAlgorithmName(string(AlgoSHA3512)):    AlgoSHA3512,
	normalizeAlgorithmName(string(AlgoBLAKE2B256)): AlgoBLAKE2B256,
	normalizeAlgorithmName(string(AlgoBLAKE2B384)): AlgoBLAKE2B384,
	normalizeAlgorithmName(string(AlgoBLAKE2B512)): AlgoBLAKE2B512,
	normalizeAlgorithmName(string(AlgoBLAKE2S256)): AlgoBLAKE2S256,
}

// SupportsStandardHMAC reports whether the algorithm can be used with crypto/hmac.
func SupportsStandardHMAC(algo Algorithm) bool {
	switch algo {
	case AlgoMD5, AlgoSHA1, AlgoSHA224, AlgoSHA256, AlgoSHA384, AlgoSHA512,
		AlgoSHA512224, AlgoSHA512256, AlgoSHA3224, AlgoSHA3256, AlgoSHA3384, AlgoSHA3512:
		return true
	default:
		return false
	}
}

// SupportsKeyedBLAKE2 reports whether the algorithm uses keyed BLAKE2 for protected fingerprints.
func SupportsKeyedBLAKE2(algo Algorithm) bool {
	switch algo {
	case AlgoBLAKE2B256, AlgoBLAKE2B384, AlgoBLAKE2B512, AlgoBLAKE2S256:
		return true
	default:
		return false
	}
}

// SupportsProtected reports whether the algorithm can be used for protected fingerprints.
func SupportsProtected(algo Algorithm) bool {
	return SupportsStandardHMAC(algo) || SupportsKeyedBLAKE2(algo)
}
