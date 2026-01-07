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
	defer file.Close()

	var nums [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)


		nums = append(nums, split)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
		return
	}

	operations := nums[len(nums)-1]
	nums = nums[:len(nums)-1]

	cols := len(operations)
	res := 0

	for col := 0; col < cols; col++ {
		colVals := make([]int, 0, len(nums))

		for row := 0; row < len(nums); row++ {

			val, err := strconv.Atoi(nums[row][col])
			if err != nil {
				panic(err)
				return
			}

			colVals = append(colVals, val)
		}

		if operations[col] == "+" {
			sum := 0
			for i := 0; i < len(colVals); i++ {
				sum += colVals[i]
			}
			res += sum
		} else {
			holder := 1
			for i := 0; i < len(colVals); i++ {
				holder *= colVals[i]
			}
			res += holder
		}
	}

	fmt.Println(res)
}