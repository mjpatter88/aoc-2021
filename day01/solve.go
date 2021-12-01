package main

import (
	"bufio"
	"fmt"
	"log"
  "strconv"
	"os"
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
}
