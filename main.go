package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Category struct {
	Good int `json:"Good"`
}

func GetInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

// func displayItems(items []Item) {
// 	fmt.Println("Your account list")
// 	for index, item := range items {
// 		fmt.Printf("%d) %s%10d\n", index, item.Name, item.Amount)
// 	}
// }

func prompt() {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := GetInput("Choose option (1 - Input, 2 - Output, 3 - Save, 4 - Exit): ", reader)

	switch opt {
	case "1":
		s := fmt.Sprintf("Item%6d", 30)
		fmt.Println(s)
		// input, _ := getInput("Your input: ", reader)
		// added, _ := strconv.Atoi(input)
		// items += added
	case "2":
		// output, _ := getInput("Your output : ", reader)
		// subtracted, _ := strconv.Atoi(output)
		// items -= subtracted
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
	}
}

func main() {
	fmt.Println("Welcome, this is an account book in terminal")
	FileOptionPrompt()
	// loadFiles()
	// prompt(load())
}
