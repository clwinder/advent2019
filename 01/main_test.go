package main

import "testing"

func TestTotalFuel(t *testing.T) {
	t.Run("mass 12", testTotalFuelFunc(12, 2))
	t.Run("mass 14", testTotalFuelFunc(14, 2))
	t.Run("mass 1969", testTotalFuelFunc(1969, 966))
	t.Run("mass 100756", testTotalFuelFunc(100756, 50346))
}

func testTotalFuelFunc(mass, expectedFuel int) func(t *testing.T) {
	return func(t *testing.T) {
		calculated := TotalFuel(mass)
		if calculated != expectedFuel {
			t.Errorf("Expected fuel %d for mass %d, got %d instead", calculated, mass, expectedFuel)
		}
	}
}

func TestFuelRequired(t *testing.T) {
	t.Run("mass 12", testFuelRequiredFunc(12, 2))
	t.Run("mass 14", testFuelRequiredFunc(14, 2))
	t.Run("mass 1969", testFuelRequiredFunc(1969, 654))
	t.Run("mass 100756", testFuelRequiredFunc(100756, 33583))
}

func testFuelRequiredFunc(mass, expectedFuel int) func(t *testing.T) {
	return func(t *testing.T) {
		calculated := FuelRequired(mass)
		if calculated != expectedFuel {
			t.Errorf("Expected fuel %d for mass %d, got %d instead", calculated, mass, expectedFuel)
		}
	}
}
