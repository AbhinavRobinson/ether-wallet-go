package wallet

import (
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
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

// Generates and returns a HD wallet with nickname and chain config.
func GenerateHDWallet(nickname string, chain ChainConfig) (*HDWalletConfig, error) {
	mnemonic, err := hdwallet.NewMnemonic(128)
	if err != nil {
		return nil, err
	}
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	return &HDWalletConfig{
		Nickname:  nickname,
		Wallet:    wallet,
		Chain:     chain,
		Addresses: make([]HDAddress, 0),
	}, nil
}

// Generates new HDWallet Address from path and wallet.
func (hdw *HDWalletConfig) GenerateHDWalletAddress(path string) error {
	derivedPath := hdwallet.MustParseDerivationPath(path)
	account, err := hdw.Wallet.Derive(derivedPath, false)
	if err != nil {
		return err
	}
	privateKey, err := hdw.Wallet.PrivateKey(account)
	if err != nil {
		return err
	}
	hdw.Addresses = append(hdw.Addresses, HDAddress{
		PrivateKey:    privateKey,
		Address:       account.Address.Hex(),
		Path:          derivedPath.String(),
		NativeBalance: "",
		Tokens:        make([]ERC20Token, 0),
	})
	return nil
}
