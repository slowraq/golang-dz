package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var operation string
	var numbers string

	fmt.Scan(&operation, &numbers)

	arr := parseNumbers(numbers)

	switch operation {
	case "SUM":
		fmt.Println(sum(arr))
	case "AVG":
		fmt.Println(avg(arr))
	case "MED":
		fmt.Println(median(arr))
	default:
		fmt.Println("Unknown operation")
	}
}

func parseNumbers(input string) []int {
	parts := strings.Split(input, ",")
	arr := make([]int, 0, len(parts))

	for _, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err == nil {
			arr = append(arr, n)
		}
	}

	return arr
}

func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

func avg(arr []int) float64 {
	return float64(sum(arr)) / float64(len(arr))
}

func median(arr []int) float64 {
	sort.Ints(arr)

	n := len(arr)

	if n%2 == 1 {
		return float64(arr[n/2])
	}

	return float64(arr[n/2-1]+arr[n/2]) / 2
}