package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	numbers := readFileAsNumbers("input.txt")

	previousValue := 0
	increments := 0
	decrements := 0

	for _, number := range numbers {
		if previousValue != 0 {
			if number > previousValue {
				increments++
			}
			if number < previousValue {
				decrements++
			}
		}
		previousValue = number
	}

	fmt.Printf("Day 1 first solution is: Increments: %d and decrements: %d\n", increments, decrements)

	previousValue = 0
	increments = 0
	decrements = 0

	for i, number := range numbers {
		secondSum := getSumNumber(numbers, i, 1)
		thirdSum := getSumNumber(numbers, i, 2)

		sum := number + secondSum + thirdSum
		if previousValue != 0 {
			if sum > previousValue {
				increments++
			}
			if sum < previousValue {
				decrements++
			}
		}

		previousValue = sum
	}

	fmt.Printf("Day 1 second solution is: Increments: %d and decrements: %d\n", increments, decrements)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSumNumber(numbers []int, index int, offset int) int {
	value := 0
	if index+offset < len(numbers) {
		value = numbers[index+offset]
	}

	return value
}

func readFileAsNumbers(path string) []int {
	file, e := os.Open(path)

	check(e)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var numbers []int

	for scanner.Scan() {
		numbers = append(numbers, convertLineToNumber(*scanner))
	}

	file.Close()

	return numbers
}

func convertLineToNumber(scanner bufio.Scanner) int {
	line := scanner.Text()
	number, e := strconv.Atoi(line)
	check(e)

	return number
}
