package model

type mockDB struct {
	lastQuery   string
	lastArgs    []interface{}
	returnedRows Rows
}

func (db *mockDB) Query(query string, args ...interface{}) (Rows, error) {
	db.lastQuery = query
	db.lastArgs = args
	return db.returnedRows, nil
}

type mockRows struct{}

func (m mockRows) Scan(dest ...interface{}) error {
	return nil
}
func (m mockRows) Next() bool{
  return false;
}

func (m mockRows) Close() error {
  return nil
}
