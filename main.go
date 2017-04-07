package main

import (
	"os"
	"fmt"
	//"text/scanner"
	"bufio"
)

var divider string = "------------"
var inputs []string
var text string


func main() {

		/*
			Adding some features: ls, cd, converter, calculator , reader 

		 */

		input := bufio.NewScanner(os.Stdin) //makes a scanner, returns a type *Scanner

		fmt.Println("Welcome")
		fmt.Println(divider)
		fmt.Println("Type -help to get started")

		for text != "-quit" {
			fmt.Print("> ")
			input.Scan()
			text = input.Text()
			if text == "-help"{
				help()
			}

		}
/*
		numArgs := len(inputs) //get the length of args

		fmt.Println("You inputted", numArgs, "arguments.")

		for i := range inputs {

			if inputs[i] == "-help" {
				help()
			}

			if inputs[i] == "-quit"{
				fmt.Println("You just quit")
			}
		}
		fmt.Println(">>>")
		fmt.Println("Path starts here::", inputs)
*/


}

func help() {
	fmt.Println("Help")
	fmt.Println(divider)
	fmt.Println("-Converter")
	fmt.Println("Takes an integer and converts the value and converts it to binary, hexidecimal, or unicode.")
	fmt.Println("Quit")
	fmt.Println("-Quits the program")
}
