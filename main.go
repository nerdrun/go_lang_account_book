package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	name   string
	amount int
}

func load() int {
	fmt.Println("Opening a file")
	file, err := os.OpenFile("./total.txt", os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Opening file error : %v", err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	total, err := strconv.Atoi(string(buf[:n]))
	if err != nil {
		fmt.Println("Converting Error %v", err)
	}
	file2, err := os.OpenFile("./test.txt", os.O_APPEND|os.O_CREATE, os.ModePerm)

	buf2 := make([]byte, 1024)

	for {
		n2, _ := file2.Read(buf2)
		fmt.Println(string(buf[:n2]))
	}

	return total
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func prompt(total int) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (1 - Input, 2 - Output, 3 - Save, 4 - Exit): ", reader)

	switch opt {
	case "1":
		s := fmt.Sprintf("Item%6d", 30)
		fmt.Println(s)
		input, _ := getInput("Your input: ", reader)
		added, _ := strconv.Atoi(input)
		total += added
		fmt.Printf("Total : %d\n", total)
		prompt(total)
	case "2":
		output, _ := getInput("Your output : ", reader)
		subtracted, _ := strconv.Atoi(output)
		total -= subtracted
		fmt.Printf("Total : %d\n", total)
		prompt(total)
	case "3":
		err := os.WriteFile("./total.txt", []byte(strconv.Itoa(total)), fs.ModePerm)
		if err != nil {
			fmt.Printf("Writing file failed : %v", err)
		}
		fmt.Println("Your file is saved successfully.")
	case "4":
		fmt.Println("Thank you")
	default:
		fmt.Println("Nothing in the list")
		prompt(total)
	}
}

func main() {
	fmt.Println("Welcome, this is an account book in terminal")
	prompt(load())
}
