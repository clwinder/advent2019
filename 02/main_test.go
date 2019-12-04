package main

import "testing"

type test struct {
	inputProgram   []int
	expectedOutput []int
}

func TestRunProgram(t *testing.T) {
	tests := []test{
		test{
			inputProgram:   []int{1, 0, 0, 0, 99},
			expectedOutput: []int{2, 0, 0, 0, 99},
		},
		test{
			inputProgram:   []int{2, 3, 0, 3, 99},
			expectedOutput: []int{2, 3, 0, 6, 99},
		},
		test{
			inputProgram:   []int{2, 4, 4, 5, 99, 0},
			expectedOutput: []int{2, 4, 4, 5, 99, 9801},
		},
		test{
			inputProgram:   []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			expectedOutput: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for i, tst := range tests {
		t.Run(string(i), testRunProgramFunc(tst.inputProgram, tst.expectedOutput))
	}
}

func testRunProgramFunc(inputProgram []int, expectedOutput []int) func(t *testing.T) {
	return func(t *testing.T) {
		calculated, err := RunProgram(inputProgram)
		if err != nil {
			t.Errorf("Expected error to be nil, got %s instead", err)
		}
		for i, c := range calculated {
			if c != expectedOutput[i] {
				t.Errorf("Expected the output to be %v, got %v instead", expectedOutput, calculated)
			}
		}
	}
}
