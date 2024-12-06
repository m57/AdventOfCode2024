package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func safetyCheckIncreasing(list []int, index int) (bool, int) {
	if index == len(list)-1 {
		return true, -1
	}

	// Good, first check passed.
	if list[index] < list[index+1] {
		// Check if the distance is safe also, this reduces the need to loop through again.
		if safetyCheckDistance(list[index], list[index+1]) {
			return safetyCheckIncreasing(list, index+1)
		} else {
			return false, index
		}
	}

	return false, index
}

func safetyCheckDecreasing(list []int, index int) (bool, int) {

	if index == len(list)-1 {
		return true, -1
	}

	if list[index] > list[index+1] {
		if safetyCheckDistance(list[index], list[index+1]) {
			return safetyCheckDecreasing(list, index+1)
		} else {
			return false, index
		}
	}

	return false, index
}

func safetyCheckDistance(num1 int, num2 int) bool {
	distance := math.Abs(float64(num1 - num2))
	// fmt.Printf("Distance: %f\n", distance)
	return distance >= 1 && distance <= 3
}

func problemDampener(list []int, indexIncrease int, indexDecrease int) bool {
	// fmt.Printf("List: %v\n", list)
	// fmt.Printf("Index Increase: %d\n", indexIncrease)
	// fmt.Printf("Index Decrease: %d\n", indexDecrease)

	for modifier := 0; modifier+indexIncrease < (len(list)); modifier++ {
		listCopy1 := append([]int(nil), list...)
		newListInc := append(listCopy1[:indexIncrease+modifier], listCopy1[indexIncrease+modifier+1:]...)
		// fmt.Printf("New List after Increase: %v\n", newListInc)
		resi, _ := safetyCheckIncreasing(newListInc, 0)
		if resi {
			return true
		}
	}

	for modifier := 0; modifier+indexDecrease < (len(list)); modifier++ {
		listCopy2 := append([]int(nil), list...)
		newListDec := append(listCopy2[:indexDecrease+modifier], listCopy2[indexDecrease+modifier+1:]...)
		// fmt.Printf("New List after Decrease: %v\n", newListDec)
		resd, _ := safetyCheckDecreasing(newListDec, 0)
		if resd {
			return true
		}
	}

	return false
}

type UnsafeList struct {
	list  []int
	index int
}

func main() {
	safeReports := 0

	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {

		nums := strings.Split(string(line), " ")
		numsLen := len(nums)
		workingList := make([]int, numsLen)

		for i := 0; i < numsLen; i++ {
			workingList[i], _ = strconv.Atoi(nums[i])
		}

		resIncrease, indexIncrease := safetyCheckIncreasing(workingList, 0)
		resDecrease, indexDecrease := safetyCheckDecreasing(workingList, 0)

		if resIncrease || resDecrease {
			safeReports++
			fmt.Printf("\033[1;32m[Safe]\033[0m %s\n", line)
		} else {
			if problemDampener(workingList, indexIncrease, indexDecrease) {
				safeReports++
				fmt.Printf("\033[1;33m[Safe]\033[0m %s\n", line)
			} else {
				fmt.Printf("\033[1;31m[Unsafe]\033[0m %s\n", line)
			}
		}
	}

	fmt.Printf("Safe Reports: %d\n", safeReports)

}
