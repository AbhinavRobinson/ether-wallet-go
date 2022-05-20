package db

import (
	badger "github.com/dgraph-io/badger/v3"
)

var Db *badger.DB

// Global Functions

func OpenDBConnection() error {
	// Open the Badger database located in the root of the project directory.
	// It will be created if it doesn't exist.
	opts := badger.DefaultOptions("../private")
	opts.IndexCacheSize = 100 << 20 // 100 mb, based on the amount of data
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

// Atomic Functions

func WriteOnDB(key string, value string) error {
	return Db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
}

func GetFromDB(key string) (string, error) {
	var result string
	err := Db.View(func(txn *badger.Txn) error {
		value, err := txn.Get([]byte(key))
		if err != nil {
			result = value.String()
		}
		return err
	})
	return result, err
}

func EraseFromDB(key string) error {
	return Db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}
