package postquantumcrypto

import (
	"crypto/sha256"
	"fmt"

	"github.com/tendermint/tendermint/crypto"
)

// PostQuantumPubKey defines your custom public key type
type PostQuantumPubKey struct {
	Key []byte
}

// Address returns the public key's address
func (pubKey PostQuantumPubKey) Address() crypto.Address {
	return crypto.AddressHash(pubKey.Key)
}

// Bytes returns the byte representation of the public key
func (pubKey PostQuantumPubKey) Bytes() []byte {
	return pubKey.Key
}

// VerifySignature verifies a signature against the provided data
func (pubKey PostQuantumPubKey) VerifySignature(msg []byte, sig []byte) bool {
	fmt.Println("PostQuantumPubKey: VerifySignature called")
	fmt.Printf("PostQuantumPubKey: Message: %x\n", msg)
	fmt.Printf("PostQuantumPubKey: Signature: %x\n", sig)
	fmt.Printf("PostQuantumPubKey: PublicKey: %x\n", pubKey.Bytes())

	expectedSig := sha256.Sum256(append(pubKey.Key, msg...))
	fmt.Printf("PostQuantumPubKey: Expected Signature: %x\n", expectedSig[:])

	verified := string(expectedSig[:]) == string(sig)
	fmt.Printf("PostQuantumPubKey: Verification result: %v\n", verified)

	return verified
}

// Equals checks equality between two public keys
func (pubKey PostQuantumPubKey) Equals(other crypto.PubKey) bool {
	return string(pubKey.Bytes()) == string(other.Bytes())
}

// Type returns the type of the public key (as a string)
func (pubKey PostQuantumPubKey) Type() string {
	return "PostQuantumPubKey"
}
