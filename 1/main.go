package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func getListsDistance(list1 []int, list2 []int) int {
	totalDistance := 0.0

	for i := 0; i < len(list1); i++ {
		totalDistance += math.Abs(float64(list1[i] - list2[i]))
	}

	return int(totalDistance)
}

func getListsSimilarityScore(list1 []int, list2 []int) int {
	similarityScore := 0
	matches := make(map[int]int)
	list_len := len(list1)

	for i := 0; i < list_len; i++ {
		if matches[list1[i]] != 0 {
			similarityScore += matches[list1[i]]
		} else {
			for j := 0; j < list_len; j++ {
				if list1[i] == list2[j] {
					matches[list1[i]] += 1
				}
			}
			similarityScore += list1[i] * matches[list1[i]]
		}
	}

	return similarityScore
}

func main() {

	list1 := []int{}
	list2 := []int{}

	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		num1, num2 := 0, 0
		fmt.Sscanf(string(line), "%d   %d", &num1, &num2)
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := getListsDistance(list1, list2)
	fmt.Printf("Total distance: %d\n", totalDistance)

	similarityScore := getListsSimilarityScore(list1, list2)
	fmt.Printf("Similarity score: %d\n", similarityScore)

}
