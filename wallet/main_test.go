package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test if wallet generates correctly
func TestWalletGenerate(t *testing.T) {
	w, err := Generate("test", ChainConfig{Id: 1, Name: "test", RPC: "test"})
	assert.Nil(t, err)
	assert.Equal(t, "test", w.Nickname)
	assert.NotNil(t, w.PrivateKey)
	assert.Equal(t, 1, w.Chain.Id)
	assert.Equal(t, "test", w.Chain.Name)
	assert.Equal(t, "test", w.Chain.RPC)
	assert.Equal(t, 0, len(w.Addresses))
}
