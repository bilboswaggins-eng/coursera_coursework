package main

import (
	"fmt"
	"log"
)

// Animal Interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}
type Bird struct {
	name string
}
type Snake struct {
	name string
}

// Cow Methods
func (c *Cow) Eat() {
	println("grass")
}
func (c *Cow) Move() {
	println("walk")
}
func (c *Cow) Speak() {
	println("moo")
}

// Bird Methods
func (b *Bird) Eat() {
	println("worms")
}
func (b *Bird) Move() {
	println("fly")
}
func (b *Bird) Speak() {
	println("peep")
}

// Snake Methods
func (s *Snake) Eat() {
	println("mice")
}
func (s *Snake) Move() {
	println("slither")
}
func (s *Snake) Speak() {
	println("hsss")
}

func main() {
	animals := make(map[string]Animal)
	//Intro Promt
	fmt.Print("Wanna play a little game? Following are the commands:\n")
	for {
		fmt.Print("Create new animal (all strings should be separated by spaces): newanimal <name of you animal> <animal type cow/bird/snake>\n")
		fmt.Print("Query information on animals: query <name of you animal> <information on animal eat/move/speak>\n")
		var iCommand, iName, iOption string
		_, err := fmt.Scanln(&iCommand, &iName, &iOption)
		if err != nil {
			fmt.Print("Exiting program due to error.")
			log.Fatal(err)
		}
		//Iterate through options
		switch iCommand {
		case "newanimal":
			switch iOption {
			case "cow":
				{
					animals[iName] = &Cow{name: iName}
					/*fmt.Printf("Your created a %s called %s!!\n", iOption, iName)*/
				}
			case "bird":
				{
					animals[iName] = &Bird{name: iName}
					fmt.Printf("Your created a %s called %s!!\n", iOption, iName)
				}
			case "snake":
				{
					animals[iName] = &Snake{name: iName}
					fmt.Printf("Your created a %s called %s!!\n", iOption, iName)
				}
			default:
				{
					fmt.Print("Error, try agian.")
				}
			}
		case "query":
			animal, _ := animals[iName]
			switch iOption {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		default:
			fmt.Print("Request not understood.")
		}
	}
}
