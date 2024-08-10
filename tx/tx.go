package tx

import (
    "errors"
    "github.com/cosmos/cosmos-sdk/crypto/dilithium"
)

type Transaction struct {
    Message   []byte
    Signature []byte
    PubKey    []byte
}

func (tx *Transaction) Sign(privKey dilithium.DilithiumKey) error {
    sig, err := privKey.Sign(tx.Message)
    if err != nil {
        return err
    }
    tx.Signature = sig
    tx.PubKey = privKey.PublicKey
    return nil
}

func (tx *Transaction) Verify() bool {
    pubKey := dilithium.DilithiumKey{PublicKey: tx.PubKey}
    return pubKey.Verify(tx.Message, tx.Signature)
}
