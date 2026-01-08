package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	res := 0

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
			continue
		}

		k := len(line) - 12

		var stack []byte

		for i := 0; i < len(line); i++ {
			c := line[i]

			for k > 0 && len(stack) > 0 && stack[len(stack)-1] < c {
				stack = stack[:len(stack)-1]
				k--
			}

			stack = append(stack, c)
		}

		if k > 0 {
			stack = stack[:len(stack)-k]
		}

		first12 := string(stack[:12])

		val, err := strconv.Atoi(first12)
		if err != nil {
			panic(err)
			return
		}

		res += val
	}

	fmt.Println(res)
}