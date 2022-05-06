package wallet

import (
	"github.com/ethereum/go-ethereum/crypto"
)

// Generates and returns a simple ECDSA wallet with nickname and chain config.
func GenerateSimpleWallet(nickname string, chain ChainConfig) (*SimpleWalletConfig, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return &SimpleWalletConfig{
		Nickname:   nickname,
		PrivateKey: privateKey,
		Chain:      chain,
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey).Hex(),
	}, nil
}
