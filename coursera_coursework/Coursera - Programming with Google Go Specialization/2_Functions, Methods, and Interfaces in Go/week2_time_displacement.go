/*
Description: Coursera Function, Methods Module
written by: Cuban Sherman
date: 2022-08-22
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	//Prompt User for Input of variables
	fmt.Print("We are going to calculate the displacement of an object with parameters your provide(acceleration, velocity, initial displacement & time.\n")
	fmt.Print("Provice the acceleration of the object: ")
	var acceleration float64
	fmt.Scan(&acceleration)
	fmt.Print("Provice the initial velocity of the object: ")
	var initVelocity float64
	fmt.Scan(&initVelocity)
	fmt.Print("Provice the initial displacement of the object: ")
	var initDisplacement float64
	fmt.Scan(&initDisplacement)
	fmt.Print("Provice the amount seconds that the object will be moving: ")
	var time float64
	fmt.Scan(&time)
	fmt.Print("\n")

	//Process and print request
	MathRequest := GenDisplaceFn(acceleration, initVelocity, initDisplacement)
	result := MathRequest(time)
	fmt.Print(result)
}

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return 1/2*a*math.Pow(time, 2) + v*time + s
	}
	return fn
}
