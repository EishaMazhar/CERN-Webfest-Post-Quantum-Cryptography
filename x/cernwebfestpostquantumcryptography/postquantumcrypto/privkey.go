package postquantumcrypto

import (
	"crypto/rand"
	"crypto/sha256"

	"github.com/tendermint/tendermint/crypto"
)

// PostQuantumPrivKey defines your custom private key type
type PostQuantumPrivKey struct {
	Key []byte
}

// GenerateKey generates a new key pair
func GenerateKey() (PostQuantumPubKey, PostQuantumPrivKey, error) {
	privKey := make([]byte, 32)
	_, err := rand.Read(privKey)
	if err != nil {
		return PostQuantumPubKey{}, PostQuantumPrivKey{}, err
	}
	pubKey := sha256.Sum256(privKey)
	return PostQuantumPubKey{Key: pubKey[:]}, PostQuantumPrivKey{Key: privKey}, nil
}

// Sign signs a message using the custom private key
func (privKey PostQuantumPrivKey) Sign(msg []byte) ([]byte, error) {
	signature := sha256.Sum256(append(privKey.Key, msg...))
	return signature[:], nil
}

// PubKey returns the associated public key
func (privKey PostQuantumPrivKey) PubKey() crypto.PubKey {
	pubKey := sha256.Sum256(privKey.Key)
	return PostQuantumPubKey{Key: pubKey[:]}
}

// Bytes returns the byte representation of the private key
func (privKey PostQuantumPrivKey) Bytes() []byte {
	return privKey.Key
}

// Equals checks equality between two private keys
func (privKey PostQuantumPrivKey) Equals(other crypto.PrivKey) bool {
	return string(privKey.Key) == string(other.Bytes())
}

// Type returns the type of the private key (as a string)
func (privKey PostQuantumPrivKey) Type() string {
	return "PostQuantumPrivKey"
}
