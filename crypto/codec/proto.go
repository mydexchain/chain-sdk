package codec

import (
	codectypes "github.com/mydexchain/chain-sdk/codec/types"
	"github.com/mydexchain/chain-sdk/crypto/keys/ed25519"
	"github.com/mydexchain/chain-sdk/crypto/keys/multisig"
	"github.com/mydexchain/chain-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/mydexchain/chain-sdk/crypto/types"
)

// RegisterInterfaces registers the sdk.Tx interface.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface("dex.crypto.PubKey", (*cryptotypes.PubKey)(nil))
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ed25519.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &secp256k1.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &multisig.LegacyAminoPubKey{})
}
