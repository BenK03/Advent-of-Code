package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	totalRemoved := 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	matrix := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		matrix = append(matrix, row)
	}

	rows := len(matrix)

	// repeat until no more can be removed
	for {
		toRemove := [][2]int{}

		for i := 0; i < rows; i++ {
			cols := len(matrix[i])
			for j := 0; j < cols; j++ {

				if matrix[i][j] == '.' {
					continue
				}

				valid := 0

				// north
				if i-1 >= 0 && matrix[i-1][j] == '@' {
					valid++
				}

				// east
				if j+1 <= cols-1 && matrix[i][j+1] == '@' {
					valid++
				}

				// south
				if i+1 <= rows-1 && matrix[i+1][j] == '@' {
					valid++
				}

				// west
				if j-1 >= 0 && matrix[i][j-1] == '@' {
					valid++
				}

				// northwest
				if i-1 >= 0 && j-1 >= 0 && matrix[i-1][j-1] == '@' {
					valid++
				}

				// northeast
				if i-1 >= 0 && j+1 <= cols-1 && matrix[i-1][j+1] == '@' {
					valid++
				}

				// southeast
				if i+1 <= rows-1 && j+1 <= cols-1 && matrix[i+1][j+1] == '@' {
					valid++
				}

				// southwest
				if i+1 <= rows-1 && j-1 >= 0 && matrix[i+1][j-1] == '@' {
					valid++
				}

				if valid < 4 {
					toRemove = append(toRemove, [2]int{i, j})
				}
			}
		}

		// nothing removable → stop
		if len(toRemove) == 0 {
			break
		}

		// remove all at once
		for _, pos := range toRemove {
			r := pos[0]
			c := pos[1]
			matrix[r][c] = '.'
		}

		totalRemoved += len(toRemove)
	}

	fmt.Println(totalRemoved)
}