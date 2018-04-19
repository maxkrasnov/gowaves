package GoWaves

import (
	"crypto/sha256"
	"github.com/steakknife/keccak"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/curve25519"
)

// Generate address for public key
func generateAddress(publicKey []byte, chainID string) []byte {
	var address []byte
	hash := secureHash(publicKey[:])

	address = append(address, byte(1))
	address = append(address, []byte(chainID)...)
	address = append(address, hash[0:20]...)

	sum := secureHash(address)
	address = append(address, sum[0:4]...)

	return address
}

// Generate Private key for seed
func generatePrivateKey(seed string, nonce int) [32]byte {
	var b []byte
	b = append(b[:], []byte{byte(nonce), byte(nonce), byte(nonce), byte(nonce)}...)
	b = append(b[:], []byte(seed)...)

	h := secureHash(b)

	privateKey := sha256.Sum256(h)
	privateKey[0] &= 248
	privateKey[31] &= 127
	privateKey[31] |= 64

	return privateKey
}

// Generate public key for private key
func generatePublicKey(privateKey [32]byte) []byte {
	var publicKey [32]byte

	curve25519.ScalarBaseMult(&publicKey, &privateKey)

	return publicKey[:]
}

// Generate hash by blake2 and keccak
func secureHash(msg []byte) []byte {
	blake256 := blake2b.Sum256(msg)
	hashKec256 := keccak.New256()
	hashKec256.Write(blake256[:])

	return hashKec256.Sum(nil)
}
