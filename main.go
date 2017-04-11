package main

import (
	"os"
	"fmt"
	//"text/scanner"
	"bufio"
	"GoShell/converter"
	"log"
)

var divider string = "------------"
var inputs []string
var text string


func main() {

		/*
			Adding some features: ls(*readdirnames), cd,  calculator , reader

			 */

		n,err := os.Create("test.txt")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(n)

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

			if text =="-converter"{
				converter.Conv()
			}

			if text =="-mkdir"{
				//TODO: Make this take in a second parameter to work.. i.e -mkdir dirName
				var userInput string
				fmt.Println("Directory name?")
				fmt.Scanln(&userInput)
				mkdir(userInput)

			}

		}
/*
		numArgs := len(inputs) //get the length of args
		fmt.Println("You inputted", numArgs, "arguments.")
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

func mkdir(name string) {
	err := os.Mkdir(name, os.ModePerm) //gives a full r/w permission directory
	if err != nil {
		log.Fatalln(err)
	}
}
