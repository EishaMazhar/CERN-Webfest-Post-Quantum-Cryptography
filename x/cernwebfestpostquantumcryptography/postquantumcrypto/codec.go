package postquantumcrypto

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers the custom key types with the codec
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(PostQuantumPubKey{}, "postquantumcrypto/PubKey", nil)
	cdc.RegisterConcrete(PostQuantumPrivKey{}, "postquantumcrypto/PrivKey", nil)
}

func init() {
	RegisterCodec(codec.NewLegacyAmino())
}
