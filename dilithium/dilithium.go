package dilithium

import (
    "crypto/rand"
    "errors"

    dilithium "github.com/kudelskisecurity/crystals-go" // Assume there's a Go implementation or binding
)

type DilithiumKey struct {
    PublicKey  []byte
    PrivateKey []byte
}

func GenerateKey() (*DilithiumKey, error) {
	d := Dilithium2() //Creates a Dilithium instance with recommended security level
	pubKey, privKey := d.KeyGen()
    // pubKey, privKey, err := dilithium.GenerateKey(rand.Reader)
    // if err != nil {
    //     return nil, err
    // }
    return &DilithiumKey{
        PublicKey:  pubKey,
        PrivateKey: privKey,
    }, nil
}

func (key *DilithiumKey) Sign(message []byte) ([]byte, error) {
	return Dilithium2.Sign(key.privKey, message)
    // return dilithium.Sign(rand.Reader, key.PrivateKey, message)
}

func (key *DilithiumKey) Verify(message, sig []byte) bool {
	d := Dilithium2()
	return d.Verify(key.pubKey, sig, message)
    // return dilithium.Verify(key.PublicKey, message, sig)
}
