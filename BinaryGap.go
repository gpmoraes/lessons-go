package main

import (
	"fmt"
	"strconv"
)

func main() {
	ret := Solution(506)
	fmt.Printf("RESPONSE: %d\n", ret)
}

func Solution(N int) int {
	count := 0
	total := 0
	// turn int in to binary
	binaryStr := strconv.FormatInt(int64(N), 2)
	fmt.Printf("BINARY: %s\n", binaryStr)
	// iterate the binary string
	for i := 0; i < len(binaryStr); i++ {
		if binaryStr[i] == '0' {
			count++
		} else {
			if total < count {
				total = count
			}
			if count > 0 {
				count = 0
			}
		}
	}
	// pint the result
	return total
}
