/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : hmac.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : HMAC and keyed BLAKE2 protected fingerprint operations.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package hash

import (
	"crypto/hmac"
	"errors"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
)

// ErrUnsupportedHMACAlgorithm indicates that protected fingerprinting is unavailable for the algorithm.
var ErrUnsupportedHMACAlgorithm = errors.New("hash: unsupported hmac algorithm")

// Protected computes an application-specific fingerprint keyed by the machine ID.
// Standard hash algorithms use HMAC. BLAKE2 algorithms use keyed BLAKE2 mode.
func Protected(appID string, machineID string, algo Algorithm) ([]byte, error) {
	if SupportsStandardHMAC(algo) {
		return standardHMAC([]byte(appID), []byte(machineID), algo)
	}
	if SupportsKeyedBLAKE2(algo) {
		return keyedBLAKE2([]byte(appID), []byte(machineID), algo)
	}
	return nil, ErrUnsupportedHMACAlgorithm
}

func standardHMAC(message, key []byte, algo Algorithm) ([]byte, error) {
	factory, ok := newHMACHasher(algo)
	if !ok {
		return nil, ErrUnsupportedHMACAlgorithm
	}

	mac := hmac.New(factory, key)
	if _, err := mac.Write(message); err != nil {
		return nil, err
	}
	return mac.Sum(nil), nil
}

func keyedBLAKE2(message, key []byte, algo Algorithm) ([]byte, error) {
	var (
		hasher interface {
			Write([]byte) (int, error)
			Sum([]byte) []byte
		}
		err error
	)

	switch algo {
	case AlgoBLAKE2B256:
		hasher, err = blake2b.New256(key)
	case AlgoBLAKE2B384:
		hasher, err = blake2b.New384(key)
	case AlgoBLAKE2B512:
		hasher, err = blake2b.New512(key)
	case AlgoBLAKE2S256:
		hasher, err = blake2s.New256(key)
	default:
		return nil, ErrUnsupportedHMACAlgorithm
	}
	if err != nil {
		return nil, ErrUnsupportedHMACAlgorithm
	}

	if _, err := hasher.Write(message); err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}
