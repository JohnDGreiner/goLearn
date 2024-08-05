package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[:]
		sums = append(sums, Sum(tail))
	}
	return sums
}

func main() {
	fmt.Println(Sum([]int{1, 2}))
	fmt.Println(SumAll([]int{3, 5}, []int{6, 7}))
	fmt.Println(SumAllTails([]int{1, 2, 5}, []int{0, 9, 1}))
}
