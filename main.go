package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func getUserInput(msg string) []int {
	var input string
	fmt.Println(msg)

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
		panic("Error reading")
	}

	inputArr := strings.Split(input, ".")
	parsedInputArr := make([]int, len(inputArr))

	for i, num := range inputArr {
		parsedInputArr[i], err = strconv.Atoi(num)
		if err != nil {
			panic("Error converting input to integer")
		}
	}

	if len(parsedInputArr) != 4 {
		panic("Length of your input doesn't match the number of places")
	}

	return parsedInputArr
}

func isInArray(arr []int, toFind int) bool {
	for _, v := range arr {
		if v == toFind {
			return true
		}
	}

	return false
}

func calculateCorrect(user, sequence []int) (int, int) {
	correctPlace := 0
	correctColor := 0
	var matchedColors []int

	for i, u := range user {
		if u == sequence[i] {
			correctPlace++
		} else {
			if isInArray(sequence, u) && !isInArray(matchedColors, u) {
				correctColor++
				matchedColors = append(matchedColors, u)
			}
		}
	}

	return correctPlace, correctColor
}

func main() {
	// Nasleduji snip generuje automaticke sequence
	// Odkomentovat pro hru
	//
	// colorCount := 6
	// generating sequence
	// sequence := make([]int, colorCount)
	// for i := 0; i < colorCount; i++ {
	//	sequence[i] = rand.Intn(colorCount) + 1
	// }
	sequence := []int{1, 1, 3, 4}
	correct := false

	for !correct {
		input := getUserInput("Napis svuj guess jako cisla oddelene teckou")

		if slices.Equal(input, sequence) {
			correct = true
		}

		fmt.Println(calculateCorrect(input, sequence))
	}

	fmt.Println("Uhodnuto")
	fmt.Println(sequence)
}
