package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

func load() int {
	fmt.Println("Opening a file")
	file, err := os.OpenFile("./total.txt", os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Opening file error : %v", err)
	}
	defer file.Close()

	// TODO: No total... correctly
	buf := make([]byte, 10)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
	fmt.Println(buf)
	// var a byte
	// binary.Read(bytes.NewReader(buf), binary.BigEndian, &a)
	// total := int(a)
	// fmt.Printf("Total : %v \n", total)
	return 0
}

func main() {
	fmt.Println("Welcome, this is an account book in terminal")
	total := load()

	keepGoing := true

	for keepGoing {
		reader := bufio.NewReader(os.Stdin)
		displayMainMenu()

		command, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		switch strings.TrimSuffix(command, "\n") {
		case "1":
			// TODO: refactoring
			keepGoing := true
			for keepGoing {
				fmt.Println("How much do you want to add?")
				command, err := reader.ReadString('\n')
				if err != nil {
					panic(err)
				}
				newInput := strings.TrimSuffix(command, "\n")
				added, _ := strconv.Atoi(newInput)
				total += added
				fmt.Printf("total : %d \n\n", total)

				keepGoing = toContinue("Do you want to add more?", reader)
			}
		case "2":
			keepGoing := true
			for keepGoing {
				fmt.Println("How much do you want to subtract?")
				command, err := reader.ReadString('\n')
				if err != nil {
					panic(err)
				}
				newInput := strings.TrimSuffix(command, "\n")
				subtracted, _ := strconv.Atoi(newInput)
				total -= subtracted
				fmt.Printf("total : %d \n\n", total)

				keepGoing = toContinue("Do you want to subtract more?", reader)
			}
		case "3":
			err := os.WriteFile("./total.txt", []byte(strconv.Itoa(total)), fs.ModePerm)
			if err != nil {
				fmt.Printf("Writing file failed : %v", err)
			}
			fmt.Println("You selected 3")
		case "4":
			fmt.Println("You selected 4")
			keepGoing = false
		default:
			fmt.Println("Nothing in the list")
		}
	}

	fmt.Println("Thank you")
}

func displayMainMenu() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Add input")
	fmt.Println("2. Add output")
	fmt.Println("3. Save")
	fmt.Println("4. Exit")
	fmt.Print("Select: ")
}

func toContinue(title string, reader *bufio.Reader) bool {
	fmt.Println(title)
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	command, _ := reader.ReadString('\n')
	return strings.TrimSuffix(command, "\n") != "2"
}
