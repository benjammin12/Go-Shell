package main

import ("os"
	"fmt"
)

func main() {
	x := os.Args //slice of args
	//str := make([]string, 5)

	numArgs := len(x)

	fmt.Println("You inputted" , numArgs, "arguments.")

	for i := 0 ; i < numArgs; i++ {
		fmt.Println("You inputted" ,x[i])
	}
}
