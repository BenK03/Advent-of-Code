package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
		return
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	var ranges [][]int
	var ids []int

	toggle := true

	for _, line := range lines {
		if line == "" {
			toggle = false
			continue
		}

		if toggle {
			parts := strings.Split(line, "-")
			front, _ := strconv.Atoi(parts[0])
			rear, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, []int{front, rear})
		} else {
			value, _ := strconv.Atoi(line)
			ids = append(ids, value)
		}
	}

	res := 0
	for _, id := range ids {
		for _, r := range ranges {
			if r[0] <= id && id <= r[1] {
				res++
				break // do not overcount
			}
		}
	}

	fmt.Println(res)
}