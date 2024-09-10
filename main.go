package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Name   string
	Amount int
}

type Category struct {
	Good int `json:"Good"`
}

func load() []Item {
	fmt.Println("Opening a file")
	file, err := os.OpenFile("./total.txt", os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("Opening file error : %v\n", err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, _ := file.Read(buf)
	itemSlices := strings.Fields(string(buf[:n]))

	var items []Item
	for i := 0; i < len(itemSlices); i += 2 {
		amount, _ := strconv.Atoi(itemSlices[i+1])
		items = append(items, Item{Name: itemSlices[i], Amount: amount})
	}

	return items
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func displayItems(items []Item) {
	fmt.Println("Your account list")
	for index, item := range items {
		fmt.Printf("%d) %s%10d\n", index, item.Name, item.Amount)
	}
}

func fileOptionPrompt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose option")
	fmt.Println("l - Load a file")
	fmt.Println("c - Create a file")
	fmt.Println("d - Delete a file")
	opt, _ := getInput("---> ", reader)
	switch opt {
	case "l":
		files, err := lookupFolder()
		if err != nil {
			fmt.Println(err)
			fileOptionPrompt()
		}
		index := selectFile(files, reader)
		selectedFile := (*files)[index]
		loadFile(selectedFile.Name())
	case "c":
		createFile()
	case "d":
		deleteFile()
	default:
		fileOptionPrompt()
	}
}

func prompt(items []Item) {
	reader := bufio.NewReader(os.Stdin)
	displayItems(items)
	opt, _ := getInput("Choose option (1 - Input, 2 - Output, 3 - Save, 4 - Exit): ", reader)

	switch opt {
	case "1":
		s := fmt.Sprintf("Item%6d", 30)
		fmt.Println(s)
		// input, _ := getInput("Your input: ", reader)
		// added, _ := strconv.Atoi(input)
		// items += added
		fmt.Printf("Total : %d\n", items)
		prompt(items)
	case "2":
		// output, _ := getInput("Your output : ", reader)
		// subtracted, _ := strconv.Atoi(output)
		// items -= subtracted
		fmt.Printf("Total : %d\n", items)
		prompt(items)
	case "3":
		// err := os.WriteFile("./total.txt", []byte(strconv.Itoa(items)), fs.ModePerm)
		// if err != nil {
		// 	fmt.Printf("Writing file failed : %v", err)
		// }
		fmt.Println("Your file is saved successfully.")
	case "4":
		fmt.Println("Thank you")
	default:
		fmt.Println("Nothing in the list")
		prompt(items)
	}
}

func main() {
	fmt.Println("Welcome, this is an account book in terminal")
	fileOptionPrompt()
	// loadFiles()
	// prompt(load())
}
