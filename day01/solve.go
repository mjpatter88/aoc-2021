package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Solve(input []int) int {
	// The only trick is that the first item should never count as increasing.
	// To accomplish this, we just set the initial "lastMeasurement" value to
	// be the first input value.
	if len(input) == 0 {
		return 0
	}
	numIncreasing := 0
	lastMeasurement := input[0]

	for _, value := range input {
		if value > lastMeasurement {
			numIncreasing += 1
		}
		lastMeasurement = value
	}
	return numIncreasing
}

func Solve2(input []int) int {
	// In this case we need > 3 elements to have at least two windows to compare
	if len(input) <= 3 {
		return 0
	}
	numIncreasing := 0
	newSum := 0
	lastSum, err := CalcWindowSum(input, 0)
	index := 1

	for err == nil {
		newSum, err = CalcWindowSum(input, index)
		if err == nil && newSum > lastSum {
			numIncreasing += 1
		}
		lastSum = newSum
		index += 1
	}
	return numIncreasing
}

// Returns the sum of a three-measurement window starting at the given index.
// Returns an error if there are not enough input elements.
func CalcWindowSum(input []int, startingIndex int) (int, error) {
	maxStartingIndex := len(input) - 3
	if startingIndex > maxStartingIndex {
		return 0, errors.New("Not enough elements to calculate three-measurement window.")
	}
	return input[startingIndex] + input[startingIndex+1] + input[startingIndex+2], nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numString := scanner.Text()
		num, err := strconv.Atoi(numString)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	answer := Solve(input)
	fmt.Println(answer)

	answer2 := Solve2(input)
	fmt.Println(answer2)
}
