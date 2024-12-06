package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[+] Data size: %d\n", len(data))

	totalSum := 0
	totalMatches := 0

	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	for _, match := range r.FindAllSubmatch(data, -1) {
		num1, _ := strconv.Atoi(string(match[1]))
		num2, _ := strconv.Atoi(string(match[2]))
		totalSum += num1 * num2
		totalMatches++
	}

	fmt.Printf("[+] Total matches: %d\n", totalMatches)
	fmt.Printf("[+] Total sum: %d\n", totalSum)
}
