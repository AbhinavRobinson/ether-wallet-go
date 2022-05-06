package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test if wallet generates correctly
func TestGenerate(t *testing.T) {
	w, err := Generate("test", ChainConfig{Id: 1, Name: "test", RPC: "test"})
	assert.Nil(t, err)
	assert.NotNil(t, w.PrivateKey)
	assert.Equal(t, "test", w.Nickname)
	assert.Equal(t, 1, w.Chain.Id)
	assert.Equal(t, "test", w.Chain.Name)
	assert.Equal(t, "test", w.Chain.RPC)
	assert.Equal(t, "", w.Address)
}

// Should be able to generate address
func TestGenerateAddress(t *testing.T) {
	w, err := Generate("test", ChainConfig{Id: 1, Name: "test", RPC: "test"})
	assert.Nil(t, err)
	assert.Equal(t, "", w.Address)
	w.GenerateAddress()
	assert.NotEqual(t, "", w.Address)
}
