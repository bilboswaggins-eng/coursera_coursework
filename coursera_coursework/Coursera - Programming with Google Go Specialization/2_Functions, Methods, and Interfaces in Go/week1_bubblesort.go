package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Print("Type in a sequence of 10 integers, line by line (press enter after each number).")

	//User input to slice
	var userInput string
	intSlice := make([]int, 0)
	sequencePosition := 0
	sequenceLen := 10
	for sequencePosition < sequenceLen {
		fmt.Scan(&userInput)
		userInt, err := strconv.Atoi(userInput)
		if err == nil {
			sequencePosition = sequencePosition + 1
			fmt.Printf(" %d / 10 integers entered.\n", sequencePosition)
			intSlice = append(intSlice, userInt)
		} else if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Print("\nYou entered the following numbers: ", intSlice)
	BubbleSort(intSlice)
}

// The BubbleSort() repeats swap for the same amount of the slices length
func BubbleSort(intSlice []int) {
	fmt.Print("\n...\nstarting BubbleSort\n...\n")
	var index int = 0
	var sliceLen int = len(intSlice)
	var rounds int = 0

	for rounds < sliceLen {
		index = 0
		for index < sliceLen-1 {
			if intSlice[index] > intSlice[index+1] {
				Swap(intSlice, index)
			}
			index = index + 1
		}
		rounds = rounds + 1
	}
	fmt.Print("\n... and BubbleSort gives you a organized slice!:", intSlice)
}

// Swap function
func Swap(intSlice []int, index int) {
	var firstNumber int = intSlice[index]
	var secondNumber int = intSlice[index+1]
	intSlice[index] = secondNumber
	intSlice[index+1] = firstNumber

}
