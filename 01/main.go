package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	values := strings.Split(string(content), "\n")

	var totalFuel int
	for _, value := range values {
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Failed to convert string value to int: %s", err)
		}
		totalFuel += TotalFuel(v)
	}

	fmt.Printf("Total fuel required = %d\n", totalFuel)
}

// TotalFuel calculates the total fuel required for a module of a given mass,
// including any fuel required for the extra fuel.
func TotalFuel(mass int) int {
	m := mass
	var tot int
	for {
		f := FuelRequired(m)
		if f <= 0 {
			break
		}
		m = f
		tot += m
	}
	return tot
}

// FuelRequired calculates the fuel required for specific mass.
func FuelRequired(mass int) int {
	return (mass / 3) - 2
}
