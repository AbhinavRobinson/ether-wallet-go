package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Setup suite for opening and closing database.
func setup(tb testing.TB) func(tb testing.TB) {
	err := OpenDBConnection()
	assert.Nil(tb, err)
	// Return a function to teardown the test
	return func(tb testing.TB) {
		defer CloseDBConnection()
	}
}

// test if connection opens and closes.
func TestOpenDBConnection(t *testing.T) {
	teardown := setup(t)
	teardown(t)
}

// test if data is dropped.
func TestClearDBData(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	err := DropDBData()
	assert.Nil(t, err)
}

// test if data is writen and fetched, deleted and not fetched
func TestAtomic(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	err := WriteOnDB("A", "1")
	assert.Nil(t, err)

	result, err := GetFromDB("A")
	assert.Nil(t, err)
	assert.Equal(t, "1", result)

	err = EraseFromDB("A")
	assert.Nil(t, err)

	assert.Panics(t, func() { GetFromDB("A") })
}
