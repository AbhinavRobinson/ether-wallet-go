package wallet

import (
	"crypto/ecdsa"
)

/*
	Config for creating a simple ECDSA wallet.
*/
type SimpleWalletConfig struct {
	Nickname      string
	PrivateKey    *ecdsa.PrivateKey
	Chain         ChainConfig
	Address       string // is keccak of pubkey
	PublicAddress string
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
