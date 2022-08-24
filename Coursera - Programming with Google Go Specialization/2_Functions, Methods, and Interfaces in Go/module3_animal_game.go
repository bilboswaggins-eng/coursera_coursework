package main

import (
	"fmt"
)

// Animal Class
type Animal struct {
	name       string
	food       string
	locomotion string
	sound      string
}

// Constructor Method
func newAnimal(name string, food string, locomotion string, sound string) Animal {
	a := Animal{name: name, food: food, locomotion: locomotion, sound: sound}
	return a
}

// Animal Actions
func (myanimal *Animal) Eat() string {
	return myanimal.food
}
func (myanimal *Animal) Move() string {
	return myanimal.locomotion
}
func (myanimal *Animal) Speak() string {
	return myanimal.sound
}

func main() {
	//Getting and verifying user input
	var userAnimal string
	var animalAction string
	fmt.Print("Welcome to the animal game!!! Give me 2 strings in one line separated by a space. The first part should be an animal (cow/bird/snake) and the second part of the string should be the information requested on the animal (eat/move/speak).\n")
	fmt.Print(">")
	fmt.Scanf("%s %s", &userAnimal, &animalAction)
	fmt.Printf("\nYou specified animal: %s and action: %s.\n\n", userAnimal, animalAction)

	//Creating Animals
	cow := newAnimal("cow", "grass", "walk", "moo")
	bird := newAnimal("bird", "worms", "fly", "peep")
	snake := newAnimal("snake", "mice", "slither", "hsss")
	animals := []Animal{cow, bird, snake}

	//Iterate through Structs to find appropriate response
	if animalAction == "eat" {
		for _, a := range animals {
			if a.name == userAnimal {
				fmt.Print(a.Eat())
			}
		}
	} else if animalAction == "move" {
		for _, a := range animals {
			if a.name == userAnimal {
				fmt.Print(a.Move())
			}
		}
	} else if animalAction == "speak" {
		for _, a := range animals {
			if a.name == userAnimal {
				fmt.Print(a.Speak())
			}
		}
	}
}
