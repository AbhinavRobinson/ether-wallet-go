package db

import "testing"

// test if connection opens and closes.
func TestOpenDBConnection(t *testing.T) {
	err := OpenDBConnection()
	if err != nil {
		panic(err)
	}
	err = CloseDBConnection()
	if err != nil {
		panic(err)
	}
}
