package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var ranges [][]int

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		front, _ := strconv.Atoi(parts[0])
		rear, _ := strconv.Atoi(parts[1])

		ranges = append(ranges, []int{front, rear})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	var fixedRanges [][]int

	for _, r := range ranges {
		front := r[0]
		rear := r[1]

		if len(fixedRanges) == 0 {
			fixedRanges = append(fixedRanges, []int{front, rear})
		} else {
			last := fixedRanges[len(fixedRanges)-1]
			lastEnd := last[1]

			if front <= lastEnd+1 {
				if rear > lastEnd {
					last[1] = rear
				}
			} else {
				fixedRanges = append(fixedRanges, []int{front, rear})
			}
		}
	}

	total := 0
	for _, r := range fixedRanges {
		total += (r[1] - r[0] + 1)
	}

	fmt.Println(total)
}