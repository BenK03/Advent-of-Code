package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func main() {

	// set dial tracker and count tracker
	tracker := 50
	count := 0

	// Open the input file and save it to file variable
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	// For each line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
		// Get the operation (L or R)
		operation := line[0]

		// Get the distance and make it into an integer
		dist, err := strconv.Atoi(strings.TrimSpace(line[1:]))
		if err != nil {
			panic(err)
		}

		// Create a second distance variable to manipulate without losing original
		dist2 := dist

		// operation is L
		if operation == 'L' {
			count += (dist / 100) // account for all passes of 0
			dist2 = dist % 100

			// if true then we pass 0 once more
			if tracker != 0 && dist2 >= tracker {
				count++
			}

		// operation is R
		} else {
				count += (dist/100) // account for all passes of 0
				dist2 = dist % 100

				// if true then we pass 0 once more
				if tracker != 0 && dist2 + tracker >= 100 {
					count++
				}
			}

		// Update tracker
		// if L
		if operation == 'L' {
			tracker = (tracker - dist) % 100
			if tracker < 0 {
				tracker += 100
			}

		// if R
		} else {
			tracker = (tracker + dist) % 100
		}


	}

	// Print Answer
	fmt.Println(count)

}