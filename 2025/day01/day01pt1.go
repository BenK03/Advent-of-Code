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

		// If operation is L
		if operation == 'L' {
			tracker = (tracker - dist) % 100
			if tracker < 0 {
				tracker += 100
			}
			// if tracker is 0, increment count
			if tracker == 0 {
				count++
			}
		
		// If operation is R
		} else {
			tracker = (tracker + dist) % 100
			// if tracker is 0, increment count
			if tracker == 0 {
				count++
			}
		}

	}

	// Print Answer
	fmt.Println(count)

}