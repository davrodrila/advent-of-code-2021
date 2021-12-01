package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, e := os.Open("input.txt")

	check(e)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	previousValue := 0
	increments := 0
	decrements := 0

	for _, line := range txtlines {
		number, e := strconv.Atoi(line)
		check(e)
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
	fmt.Printf("Day 1 first solution is: Increments: %d and decrements: %d", increments, decrements)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
