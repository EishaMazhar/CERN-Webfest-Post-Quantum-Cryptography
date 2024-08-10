package test

import (
    "testing"
    "github.com/cosmos/cosmos-sdk/crypto/dilithium"
)

func TestDilithiumKeyGeneration(t *testing.T) {
    key, err := dilithium.GenerateKey()
    if err != nil {
        t.Fatalf("Key generation failed: %v", err)
    }
    if len(key.PublicKey) == 0 || len(key.PrivateKey) == 0 {
        t.Fatalf("Key generation produced empty keys")
    }
}

func TestDilithiumSigning(t *testing.T) {
    key, _ := dilithium.GenerateKey()
    message := []byte("test message")
    signature, err := key.Sign(message)
    if err != nil {
        t.Fatalf("Signing failed: %v", err)
    }
    if !key.Verify(message, signature) {
        t.Fatalf("Signature verification failed")
    }
}
