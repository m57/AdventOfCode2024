package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func doMulOnString(data []byte) int {

	totalSum := 0

	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	for _, match := range r.FindAllSubmatch(data, -1) {
		num1, _ := strconv.Atoi(string(match[1]))
		num2, _ := strconv.Atoi(string(match[2]))
		totalSum += num1 * num2
		fmt.Printf("\033[1;33m[+]\033[0m mul(%d, %d) = %d\n", num1, num2, num1*num2)
	}
	return totalSum
}

func partA(data []byte) {

	totalSum := 0
	doMulOnString(data)
	fmt.Printf("[+] Total sum: %d\n", totalSum)
}

func partB(data []byte) {

	totalSum := 0

	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	rDont, _ := regexp.Compile(`(don\'t\(\)){1}`)
	blocks := strings.Split(string(data), "do()")

	for _, block := range blocks {

		skip := false
		fmt.Printf("\n\033[1;32m[+]\033[0m Block: %s\n", block)

		muls := r.FindAllStringIndex(block, -1)
		donts := rDont.FindAllStringIndex(block, -1)
		if len(donts) == 0 {
			donts = [][]int{{0, 0}}
			skip = true
		}

		for _, mul := range muls {
			if mul[0] < donts[0][0] || skip {
				mulStr := []byte(block[mul[0]:mul[1]])
				totalSum += doMulOnString(mulStr)
			}

		}
	}
	fmt.Printf("\n[+] Total sum: %d\n", totalSum)
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[+] Data size: %d\n", len(data))

	partA(data)
	partB(data)

}
