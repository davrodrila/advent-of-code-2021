package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	lines := readFileAsLines("input.txt")

	doFirstDay(lines)
	doSecondDay(lines)

}

func readFileAsLines(path string) []string {
	file, e := os.Open(path)

	check(e)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var orders []string

	for scanner.Scan() {
		orders = append(orders, scanner.Text())
	}

	file.Close()

	return orders
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func doFirstDay(lines []string) {
	depth, horizontalPosition := 0, 0
	for _, line := range lines {
		tokenized := strings.Split(line, " ")
		if tokenized[0] == "forward" {
			horizontalPosition += convertToNumber(tokenized[1])
		}
		if tokenized[0] == "up" {
			depth -= convertToNumber(tokenized[1])
		}
		if tokenized[0] == "down" {
			depth += convertToNumber(tokenized[1])
		}
	}
	fmt.Printf("Day 2 first solution Horizontal Position: %d Depth: %d Result %d\n", horizontalPosition, depth, depth*horizontalPosition)
}

func doSecondDay(lines []string) {
	aim, depth, horizontalPosition := 0, 0, 0
	for _, line := range lines {
		tokenized := strings.Split(line, " ")
		if tokenized[0] == "forward" {
			number := convertToNumber(tokenized[1])
			horizontalPosition += number
			depth += aim * number
		}
		if tokenized[0] == "up" {
			aim -= convertToNumber(tokenized[1])
		}
		if tokenized[0] == "down" {
			aim += convertToNumber(tokenized[1])
		}
	}
	fmt.Printf("Day 2 second solution Horizontal Position: %d Depth: %d Result %d\n", horizontalPosition, depth, depth*horizontalPosition)
}

func convertToNumber(line string) int {
	number, e := strconv.Atoi(line)
	check(e)

	return number
}
