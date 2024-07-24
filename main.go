package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome, this is an account book in terminal")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What do you want to do?")
	fmt.Println("1. Add input")
	fmt.Println("2. Add output")
	fmt.Println("3. Save")
	fmt.Println("4. Exit")
	fmt.Print("Select: ")

	command, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	switch strings.TrimSuffix(command, "\n") {
	case "1":
		fmt.Println("You selected 1")
	case "2":
		fmt.Println("You selected 2")
	case "3":
		fmt.Println("You selected 3")
	case "4":
		fmt.Println("You selected 4")
	default:
		fmt.Println("Nothing in the list")
	}

	fmt.Println("Thank you")
}
