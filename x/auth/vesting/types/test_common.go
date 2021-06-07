package types

import (
	"github.com/mydexchain/chain-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/mydexchain/chain-sdk/crypto/types"
	"github.com/mydexchain/chain-sdk/testutil/testdata"

	sdk "github.com/mydexchain/chain-sdk/types"
)

// NewTestMsg generates a test message
func NewTestMsg(addrs ...sdk.AccAddress) *testdata.TestMsg {
	return testdata.NewTestMsg(addrs...)
}

// NewTestCoins coins to more than cover the fee
func NewTestCoins() sdk.Coins {
	return sdk.Coins{
		sdk.NewInt64Coin("atom", 10000000),
	}
}

// KeyTestPubAddr generates a test key pair
func KeyTestPubAddr() (cryptotypes.PrivKey, cryptotypes.PubKey, sdk.AccAddress) {
	key := secp256k1.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return key, pub, addr
}