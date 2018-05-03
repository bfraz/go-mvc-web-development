package model

import "testing"


func TestRetreivalofLPWithIDUsesExpectedID(t *testing.T) {
	testDB := new(mockDB)
	testDB.returnedRows = &mockRows{}
	db = testDB

	livestockID := 2
	_, err := GetLivestockProductsWithLivestockID(livestockID)
    if err != nil{
        t.Errorf("err should be nil")
    }
	if testDB.lastArgs[0] != livestockID {
		t.Errorf("LivestockID should be as expected")
	}
}
