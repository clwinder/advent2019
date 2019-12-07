package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

// Day03 represents the answers for the day 3 problem.
type Day03 struct {
	part1 int
	part2 int
}

// CartesianCoord represents the cartesian coordinate system.
type CartesianCoord struct {
	x int
	y int
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	values := strings.Split(string(content), "\n")
	path1 := strings.Split(values[0], ",")
	path2 := strings.Split(values[1], ",")

	answers, err := NearestDistance(path1, path2)
	if err != nil {
		log.Fatalf("Failed to find nearest distance: %s", err)
	}
	fmt.Println("Part 1: ", answers.part1)
	fmt.Println("Part 2: ", answers.part2)
}

// NearestDistance calculates the Manhattan distance to the nearest point to the origin
// where the two wires cross, and the shortest path to an intersection.
func NearestDistance(path1, path2 []string) (Day03, error) {
	log.Println("Started looking for nearest intersect")
	coords1, err := PathToCoords(path1)
	if err != nil {
		return Day03{}, err
	}
	coords2, err := PathToCoords(path2)
	if err != nil {
		return Day03{}, err
	}

	var distances []int
	var pathLengths []int
	for i, c1 := range coords1 {
		for j, c2 := range coords2 {
			if (c1.x == c2.x) && (c1.y == c2.y) {
				dist := manhattanDist(c1.x, c1.y)
				if dist != 0 {
					distances = append(distances, dist)
				}
				pathLength := i + j
				if pathLength != 0 {
					pathLengths = append(pathLengths, pathLength)
				}
			}
		}
	}

	log.Println("Finished looking for nearest intersect")

	return Day03{
		part1: minInt(distances),
		part2: minInt(pathLengths),
	}, nil
}

// PathToCoords converts the path from instructions to cartiesian coordinates.
func PathToCoords(path []string) ([]CartesianCoord, error) {
	coords := []CartesianCoord{
		CartesianCoord{
			x: 0,
			y: 0,
		},
	}
	for _, p := range path {
		direction := string(p[0])
		distance, err := strconv.Atoi(p[1:])
		if err != nil {
			return nil, err
		}

		switch direction {
		case "R":
			for j := 0; j < distance; j++ {
				coords = append(coords, CartesianCoord{
					x: coords[len(coords)-1].x + 1,
					y: coords[len(coords)-1].y,
				})
			}
		case "L":
			for j := 0; j < distance; j++ {
				coords = append(coords, CartesianCoord{
					x: coords[len(coords)-1].x - 1,
					y: coords[len(coords)-1].y,
				})
			}
		case "U":
			for j := 0; j < distance; j++ {
				coords = append(coords, CartesianCoord{
					x: coords[len(coords)-1].x,
					y: coords[len(coords)-1].y + 1,
				})
			}
		case "D":
			for j := 0; j < distance; j++ {
				coords = append(coords, CartesianCoord{
					x: coords[len(coords)-1].x,
					y: coords[len(coords)-1].y - 1,
				})
			}
		default:
			return nil, fmt.Errorf("Unknown direction %s", direction)
		}
	}

	return coords, nil
}

func manhattanDist(x, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func minInt(intList []int) int {
	var min int
	for i, l := range intList {
		if i == 0 {
			min = l
		} else {
			if l < min {
				min = l
			}
		}
	}
	return min
}
