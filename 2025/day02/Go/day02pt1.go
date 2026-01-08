package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	ans := 0

	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	s := string(file)

	for _, part := range strings.Split(s, ",") {
		bounds := strings.Split(part, "-")

		// convert to int
		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			panic(err)
		}

		// convert to int
		end, err := strconv.Atoi(bounds[1])
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)

			// odd number of digits can not be invalid
			if len(str)%2 != 0 {
				continue
			}

			mid := len(str) / 2

			// check if first half == last half
			if str[:mid] == str[mid:] {
				ans += i
			}
		}
	}

	fmt.Println(ans)
}