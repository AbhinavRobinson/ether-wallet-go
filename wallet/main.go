package wallet

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func Generate(nickname string, chain ChainConfig) (*WalletConfig, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return &WalletConfig{
		Nickname:   nickname,
		PrivateKey: privateKey,
		Chain:      chain,
		Addresses:  []WalletAddress{},
	}, nil
}
