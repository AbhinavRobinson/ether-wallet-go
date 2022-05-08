package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test if wallet generates correctly
func TestGenerate(t *testing.T) {
	w, err := GenerateSimpleWallet("test", ChainConfig{Id: 1, Name: "test", RPC: "test"})
	assert.Nil(t, err)
	assert.NotNil(t, w.PrivateKey)
	assert.Equal(t, "test", w.Nickname)
	assert.Equal(t, 1, w.Chain.Id)
	assert.Equal(t, "test", w.Chain.Name)
	assert.Equal(t, "test", w.Chain.RPC)
	assert.NotEqual(t, "", w.Address)
}

// Test if HDWallet generates correctly
func TestGenerateHDWallet(t *testing.T) {
	w, err := GenerateHDWallet("test", ChainConfig{Id: 1, Name: "test", RPC: "test"})
	assert.Nil(t, err)
	assert.NotNil(t, w.Wallet)
	assert.Equal(t, "test", w.Nickname)
	assert.Equal(t, 1, w.Chain.Id)
	assert.Equal(t, "test", w.Chain.Name)
	assert.Equal(t, "test", w.Chain.RPC)
	assert.NotNil(t, len(w.Addresses))
}

// Test if HDWallet address generates correctly
func TestGenerateHDWalletAddress(t *testing.T) {
	w, err := GenerateHDWallet("test", ChainConfig{Id: 1, Name: "test", RPC: "test"})
	assert.Nil(t, err)
	err = w.GenerateHDWalletAddress("m/44'/60'/0'/0/0/0")
	assert.Nil(t, err)
	assert.NotEqual(t, nil, len(w.Addresses))
	assert.NotNil(t, w.Addresses[0].PrivateKey)
	assert.NotEqual(t, "", w.Addresses[0].Address)
	assert.Equal(t, "m/44'/60'/0'/0/0/0", w.Addresses[0].Path)
	assert.Equal(t, "", w.Addresses[0].NativeBalance)
	assert.NotNil(t, w.Addresses[0].Tokens)
}
