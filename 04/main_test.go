package main

import "testing"

type test struct {
	password       string
	expectedResult bool
}

func TestValidatePassword(t *testing.T) {
	tests := []test{
		test{
			password:       "111111",
			expectedResult: true,
		},
		test{
			password:       "223450",
			expectedResult: false,
		},
		test{
			password:       "123789",
			expectedResult: false,
		},
	}

	for _, tst := range tests {
		t.Run(tst.password, testValidatePasswordFunc(tst.password, tst.expectedResult))
	}
}

func testValidatePasswordFunc(password string, expectedResult bool) func(*testing.T) {
	return func(t *testing.T) {
		calculated := ValidatePassword(password)
		if calculated != expectedResult {
			t.Errorf("Expected validate to be %t, got %t instead", expectedResult, calculated)
		}
	}
}
