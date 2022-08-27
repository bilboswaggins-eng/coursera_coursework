/*Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file,
your program should iterate through your slice of structs and print the first and last names found in each struct.*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	//Prompt the user for the name of the text file
	var localFile string

	fmt.Print("What is the path/filename of the file you want to read?\n")
	fmt.Scan(&localFile)

	//Read each line of the text file and create a struct which contains the first and last names
	file, err := os.Open(localFile)
	if err != nil {
		log.Fatal("Failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	type Person struct {
		fname string
		lname string
	}
	var sli = []*Person{}
	for scanner.Scan() {
		p1 := new(Person)
		r, _ := regexp.Compile("([A-Z])\\w+")
		p1.fname = r.FindString(scanner.Text())
		r, _ = regexp.Compile("\\s(.*)")
		p1.lname = r.FindString(scanner.Text())

		//Struct created will be added to a slice
		sli = append(sli, p1)
	}
	file.Close()

	//After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct
	for index, entry := range sli {
		fmt.Printf("%d%+v\n", index, entry)
	}
	fmt.Print(("\nEnd of Program.\n"))
}
