package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	ans := []int{} // new array to store result of each line

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		possible := []int{}
		for i := 0; i < len(line) - 1; i++ {
			for j := i + 1; j < len(line); j++ {
				
				a := int(line[i] - '0')
				b := int(line[j] - '0')
				possible = append(possible, a*10+b)
			}
		}

		mx := possible[0]
		for _, v := range possible[1:] {
			if v > mx {
				mx = v
			}
		}

		ans = append(ans, mx)
	}

	sum := 0
	for _, v := range ans {
		sum += v
	}
	fmt.Println(sum)
}