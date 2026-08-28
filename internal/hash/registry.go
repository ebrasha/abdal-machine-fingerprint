/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : registry.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Hash factory registry for extensible algorithm support.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
)

type hashFactory func() hash.Hash

var digestFactories = map[Algorithm]hashFactory{
	AlgoMD5:        func() hash.Hash { return md5.New() },
	AlgoSHA1:       func() hash.Hash { return sha1.New() },
	AlgoSHA224:     func() hash.Hash { return sha256.New224() },
	AlgoSHA256:     func() hash.Hash { return sha256.New() },
	AlgoSHA384:     func() hash.Hash { return sha512.New384() },
	AlgoSHA512:     func() hash.Hash { return sha512.New() },
	AlgoSHA512224:  func() hash.Hash { return sha512.New512_224() },
	AlgoSHA512256:  func() hash.Hash { return sha512.New512_256() },
	AlgoSHA3224:    func() hash.Hash { return sha3.New224() },
	AlgoSHA3256:    func() hash.Hash { return sha3.New256() },
	AlgoSHA3384:    func() hash.Hash { return sha3.New384() },
	AlgoSHA3512:    func() hash.Hash { return sha3.New512() },
	AlgoBLAKE2B256: func() hash.Hash { h, _ := blake2b.New256(nil); return h },
	AlgoBLAKE2B384: func() hash.Hash { h, _ := blake2b.New384(nil); return h },
	AlgoBLAKE2B512: func() hash.Hash { h, _ := blake2b.New512(nil); return h },
	AlgoBLAKE2S256: func() hash.Hash { h, _ := blake2s.New256(nil); return h },
}

func newDigestHasher(algo Algorithm) (hash.Hash, bool) {
	factory, ok := digestFactories[algo]
	if !ok {
		return nil, false
	}
	return factory(), true
}

func newHMACHasher(algo Algorithm) (hashFactory, bool) {
	factory, ok := digestFactories[algo]
	if !ok || !SupportsStandardHMAC(algo) {
		return nil, false
	}
	return factory, true
}
