package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var number_of_ones [12]int
	total := len(lines)
	for _, line := range lines {
		for i, digit := range line {
			if digit == '1' {
				number_of_ones[i]++
			}
		}
	}
	var gamma string
	var epsilon string
	for _, ones := range number_of_ones {
		if ones > total/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gamma_number, e := strconv.ParseInt(gamma, 2, 64)
	check(e)
	epsilon_number, e := strconv.ParseInt(epsilon, 2, 64)
	check(e)
	fmt.Printf("Day 3 first solution. Gamma is %v and  epsilon is %v, so power consumption is %v", gamma_number, epsilon_number, gamma_number*epsilon_number)
}

func doSecondDay(lines []string) {

}
