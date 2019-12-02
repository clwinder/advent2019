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
		totalFuel += FuelRequired(v)
	}

	fmt.Printf("Total fuel required = %d\n", totalFuel)
}

// FuelRequired calculates the fuel required for a module of a specific mass.
func FuelRequired(mass int) int {
	return (mass / 3) - 2
}
