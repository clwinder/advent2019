package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const puzzleInput = "240920-789857"

func main() {
	passwordLimits := strings.Split(puzzleInput, "-")
	if len(passwordLimits) != 2 {
		log.Fatalf("Expected there to be only 2 limits, got %d instead", len(passwordLimits))
	}
	passwordStart, err := strconv.Atoi(passwordLimits[0])
	if err != nil {
		log.Fatalf("Failed to convert %s to an int", passwordLimits[0])
	}
	passwordEnd, err := strconv.Atoi(passwordLimits[1])
	if err != nil {
		log.Fatalf("Failed to convert %s to an int", passwordLimits[1])
	}

	numValid := 0
	for p := passwordStart; p <= passwordEnd; p++ {
		valid := ValidatePassword(strconv.Itoa(p))
		if valid {
			numValid = numValid + 1
		}
	}

	fmt.Println(numValid)
}

// ValidatePassword validates that a password meets all of the requirements.
func ValidatePassword(password string) bool {
	pwSplit := strings.Split(password, "")
	if len(pwSplit) != 6 {
		return false
	}

	for i := 0; i < len(pwSplit); i++ {
		if i == 0 {
			continue
		}
		if pwSplit[i] < pwSplit[i-1] {
			return false
		}
	}

	var hasPair bool
	i := 0
	for {
		if i >= len(pwSplit)-1 {
			break
		}
		if pwSplit[i] == pwSplit[i+1] {
			if i == len(pwSplit)-2 {
				hasPair = true
				break
			}
			if pwSplit[i] == pwSplit[i+2] {
				i += 3
				if i >= len(pwSplit)-1 {
					break
				}
				if pwSplit[i] == pwSplit[i-3] {
					i++
					if i >= len(pwSplit)-1 {
						break
					}
					if pwSplit[i] == pwSplit[i-3] {
						i++
					}
				}
			} else {
				hasPair = true
				break
			}
		} else {
			i++
		}
	}

	return hasPair
}
