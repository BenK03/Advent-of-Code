package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	count := 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	matrix := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimRight(line, "\r")
		matrix = append(matrix, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
		return
	}

	rows := len(matrix)

	for i := 0; i < rows; i++ {
		cols := len(matrix[i])
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '.' {
				continue
			} else {
				valid := 0

				// check north
				if i-1 >= 0 && matrix[i-1][j] == '@' {
					valid++
				}

				// check east
				if j+1 <= cols-1 && matrix[i][j+1] == '@' {
					valid++
				}

				// check south
				if i+1 <= len(matrix)-1 && matrix[i+1][j] == '@' {
					valid++
				}

				// check west
				if j-1 >= 0 && matrix[i][j-1] == '@' {
					valid++
				}

				// check northwest
				if j-1 >= 0 && i-1 >= 0 && matrix[i-1][j-1] == '@' {
					valid++
				}

				// check northeast
				if i-1 >= 0 && j+1 <= cols-1 && matrix[i-1][j+1] == '@' {
					valid++
				}

				// check southeast
				if i+1 <= len(matrix)-1 && j+1 <= cols-1 && matrix[i+1][j+1] == '@' {
					valid++
				}

				// check southwest
				if i+1 <= len(matrix)-1 && j-1 >= 0 && matrix[i+1][j-1] == '@' {
					valid++
				}

				if valid < 4 {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}