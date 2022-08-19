package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	// 	fmt.Println("Hello, World")
	// 	fmt.Print("Hello there")
	// 	fmt.Print("Bye There")
	// sayHelloWorld("Go is fun")
	// var whattosay string

	reader := bufio.NewReader(os.Stdin)

	whattosay := doctor.Intro()

	fmt.Println(whattosay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		// fmt.Println(userInput)

		// response := doctor.Response(userInput)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}

	}
}
