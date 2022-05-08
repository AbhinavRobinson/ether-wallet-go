package db

import "testing"

func setup(tb testing.TB) func(tb testing.TB) {
	err := OpenDBConnection()
	if err != nil {
		panic(err)
	}
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
	if err != nil {
		panic(err)
	}
}
