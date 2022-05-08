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
	assert.NotEqual(t, nil, len(w.Addresses))
}
