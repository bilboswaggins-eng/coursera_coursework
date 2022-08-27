package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Initialize Empty Slice len(3)
	var rawInput string
	sli := make([]int, 3)

	//User Prompt
	fmt.Print("Hello friend, give me an Integer to add to the Slice or type X to exit the program.")
	fmt.Print("\n")

	// For loop
	for {
		fmt.Scan(&rawInput)
		intInput, err := strconv.Atoi(rawInput)
		if err == nil {
			sli = append(sli, int(intInput))
			sort.Sort(sort.IntSlice(sli))
			fmt.Print(sli, "\n")
		} else if err != nil && strings.ToLower(rawInput) == "x" {
			fmt.Print("You entered X. Goodbye")
			break
		}
	}
}
