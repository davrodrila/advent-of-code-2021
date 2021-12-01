package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	numbers := readFileAsNumbers("input.txt")

	doFirstDay(numbers)
	doSecondDay(numbers)
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertLineToNumber(scanner bufio.Scanner) int {
	line := scanner.Text()
	number, e := strconv.Atoi(line)
	check(e)

	return number
}

func doFirstDay(numbers []int) {

	previousValue := 0
	increments := 0
	decrements := 0

	for _, number := range numbers {
		increments, decrements = checkValueChanges(previousValue, number, increments, decrements)
		previousValue = number
	}

	fmt.Printf("Day 1 first solution is: Increments: %d and decrements: %d\n", increments, decrements)

}

func doSecondDay(numbers []int) {
	previousValue := 0
	increments := 0
	decrements := 0

	for i, number := range numbers {
		secondSum := getSumNumber(numbers, i, 1)
		thirdSum := getSumNumber(numbers, i, 2)

		currentValue := number + secondSum + thirdSum

		increments, decrements = checkValueChanges(previousValue, currentValue, increments, decrements)

		previousValue = currentValue
	}

	fmt.Printf("Day 1 second solution is: Increments: %d and decrements: %d\n", increments, decrements)
}

func getSumNumber(numbers []int, index int, offset int) int {
	value := 0
	if index+offset < len(numbers) {
		value = numbers[index+offset]
	}

	return value
}

func checkValueChanges(previousValue int, currentValue int, increments int, decrements int) (int, int) {
	if previousValue != 0 {
		if currentValue > previousValue {
			increments++
		}
		if currentValue < previousValue {
			decrements++
		}
	}

	return increments, decrements
}
