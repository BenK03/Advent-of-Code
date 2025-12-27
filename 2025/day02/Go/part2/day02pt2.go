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
			n := len(str)

			for j := 1; j <= n / 2; j++ {

				pattern := str[:j]
				repeats := n / j

				if strings.Repeat(pattern, repeats) == str {
					ans += i
					break
				}
				
			}
		}

	}
	
	fmt.Println(ans)

}