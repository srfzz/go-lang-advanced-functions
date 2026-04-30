package main

import "fmt"

func main() {
	numbers := []int{1, 5, 6, 7, 9}
	doubles := tranformNumbers(&numbers, double)
	triple := tranformNumbers(&numbers, triple)
	fmt.Println(doubles)
	fmt.Println(triple)
}

func tranformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}
