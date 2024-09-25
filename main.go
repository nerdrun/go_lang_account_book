package main

import (
	"bufio"
	"fmt"
	"os"

	pkg "account.com/test/pkg"
)

func prompt() {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := pkg.GetInput("Choose option (1 - Input, 2 - Output, 3 - Save, 4 - Exit): ", reader)

	switch opt {
	case "1":
		s := fmt.Sprintf("Item%6d", 30)
		fmt.Println(s)
	case "2":
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

func displayWelcome() {
	fmt.Println("Welcome ACB, this is a terminal-UI accounting book")
}

func main() {
	displayWelcome()

	buffer := bufio.NewReader(os.Stdin)
	pkg.GetInput("yo", buffer)
	// FileOptionPrompt()
	// loadFiles()
	// prompt(load())
}
