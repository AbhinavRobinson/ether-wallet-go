package wallet

import (
	"crypto/ecdsa"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

/*
	Config for creating a simple ECDSA wallet.
*/
type SimpleWalletConfig struct {
	Nickname      string
	PrivateKey    *ecdsa.PrivateKey
	Chain         ChainConfig
	Address       string // is keccak of pubkey
	NativeBalance string
	Tokens        []ERC20Token
}

/*
	Config for creating a hierarchical deterministic wallet.
*/
type HDWalletConfig struct {
	Nickname  string
	Wallet    *hdwallet.Wallet
	Chain     ChainConfig
	Addresses []HDAddress // generated addresses for each path
}

/*
	Config for the generated simple wallets from the derived paths of the HD wallet.
*/
type HDAddress struct {
	PrivateKey    *ecdsa.PrivateKey // is derived from MasterPrivateKey
	Address       string            // is keccak of pubkey
	Path          string
	NativeBalance string
	Tokens        []ERC20Token
}

/*
	Chain config on which the wallet is created.
*/
type ChainConfig struct {
	Id   int
	Name string
	RPC  string
}

/*
	ERC20 Token config, this will be tracked for each wallet address.
*/
type ERC20Token struct {
	Name     string
	Symbol   string
	Address  string
	Decimals uint8
	Balance  string
}
