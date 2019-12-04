package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Input is the input to an intCode program.
type Input struct {
	noun int
	verb int
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	intCode, err := bytesToInts(content, ",")
	if err != nil {
		log.Fatalf("Failed to convert bytes to ints: %s", err)
	}

	part1, err := Part1(intCode)
	if err != nil {
		log.Fatalf("Failed to perform part 1: %s", err)
	}
	fmt.Println("Part 1 result: ", part1)

	part2, err := Part2(intCode)
	if err != nil {
		log.Fatalf("Failed to perform part 2: %s", err)
	}
	fmt.Println("Part 2 result: ", part2)

}

// Part1 returns the answer for the first part of day 2.
func Part1(program []int) (int, error) {
	memory := make([]int, len(program))
	copy(memory, program)

	// before running the program, replace position 1 with the value 12 and replace position 2 with the value 2
	memory[1] = 12
	memory[2] = 2

	result, err := RunProgram(memory)
	if err != nil {
		return 0, err
	}

	return result[0], nil
}

// Part2 returns the answer for the second part of day 2.
func Part2(program []int) (int, error) {
	inputs, err := FindInputs(program, 19690720)
	if err != nil {
		return 0, err
	}

	result := (100 * inputs.noun) + inputs.verb

	return result, nil
}

// FindInputs finds the noun and verb inputs for a given output.
func FindInputs(program []int, output int) (Input, error) {
	memory := make([]int, len(program))
	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			input := Input{
				noun: n,
				verb: v,
			}

			copy(memory, program)
			memory[1] = input.noun
			memory[2] = input.verb

			result, err := RunProgram(memory)
			if err != nil {
				return Input{}, err
			}

			if result[0] == output {
				return input, nil
			}
		}
	}
	return Input{}, fmt.Errorf("Failed to find inputs for given output %d", output)
}

// RunProgram runs the program, processing the opcodes, and returns the result.
func RunProgram(program []int) ([]int, error) {
	for i := 0; i < len(program); i += 4 {
		if program[i] == 1 {
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
		} else if program[i] == 2 {
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
		} else if program[i] == 99 {
			break
		} else {
			return nil, fmt.Errorf("Unknown opcode %d", program[i])
		}
	}

	return program, nil
}

func bytesToInts(b []byte, delimiter string) ([]int, error) {
	strs := strings.Split(string(b), delimiter)
	var ints []int
	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}
