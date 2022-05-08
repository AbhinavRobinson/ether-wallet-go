package db

import (
	badger "github.com/dgraph-io/badger/v3"
)

var Db *badger.DB

func OpenDBConnection() error {
	// Open the Badger database located in the root of the project directory.
	// It will be created if it doesn't exist.
	opts := badger.DefaultOptions("../private")
	opts.IndexCacheSize = 100 << 20 // 100 mb or some other size based on the amount of data
	db, err := badger.Open(opts)
	if err != nil {
		return err
	}
	Db = db
	return nil
}

func CloseDBConnection() error {
	return Db.Close()
}

func DropDBData() error {
	return Db.DropAll()
}
