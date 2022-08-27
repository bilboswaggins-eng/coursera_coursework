package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Ahoy matey! Give me a string and I'll tell you if I can find Ian!")

	var userInput string
	fmt.Scan(&userInput)
	userInput = strings.ToLower(userInput)

	if strings.HasPrefix(userInput, "i") && strings.Contains(userInput, "a") && strings.HasSuffix(userInput, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
