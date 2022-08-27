/*Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//Create Map
	var user map[string]string
	user = make(map[string]string)

	var userName string
	var userAddr string

	//User Prompt & Input into map
	fmt.Print("Hi there what is your name?\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userName = scanner.Text()

	fmt.Print("And what is your address?\n")
	scanner.Scan()
	userAddr = scanner.Text()
	user[userName] = userAddr

	//Conver map into JSON & print
	userJson, _ := json.Marshal(user)
	fmt.Print(string(userJson))
}
