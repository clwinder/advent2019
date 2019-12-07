package main

import "testing"

type test struct {
	path1          []string
	path2          []string
	expectedResult Day03
}

func TestNearestDistance(t *testing.T) {
	tests := []test{
		test{
			path1: []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			path2: []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			expectedResult: Day03{
				part1: 159,
				part2: 610,
			},
		},
		test{
			path1: []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			path2: []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			expectedResult: Day03{
				part1: 135,
				part2: 410,
			},
		},
	}
	for i, tst := range tests {
		t.Run(string(i), testNearestDistanceFunc(tst.path1, tst.path2, tst.expectedResult))
	}
}

func testNearestDistanceFunc(path1, path2 []string, expectedResult Day03) func(t *testing.T) {
	return func(t *testing.T) {
		calculated, err := NearestDistance(path1, path2)
		if err != nil {
			t.Errorf("Expected error to be nil, got %s instead", err)
		}
		if expectedResult.part1 != calculated.part1 {
			t.Errorf("Expected nearest to be %d for part 1, got %d instead", expectedResult.part1, calculated.part1)
		}
		if expectedResult.part2 != calculated.part2 {
			t.Errorf("Expected shortest to be %d for part 2, got %d instead", expectedResult.part2, calculated.part2)
		}
	}
}
