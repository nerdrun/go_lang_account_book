package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	domains "account.com/test/domains"
	services "account.com/test/services"
)

func FileOptionPrompt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose option")
	fmt.Println("l - Load a file")
	fmt.Println("c - Create a file")
	fmt.Println("d - Delete a file")
	opt, _ := GetInput("---> ", reader)
	switch opt {
	case "l":
		files, err := services.LookupFolder()
		if err != nil {
			fmt.Println(err)
			FileOptionPrompt()
		}
		index := selectFile(files, reader)
		selectedFile := (*files)[index]
		f, err := services.LoadFile(selectedFile.Name())
		if err != nil {
			panic(err)
		}
		buf, err := io.ReadAll(f)

		if err != nil {
			panic(err)
		}

		var account domains.Account
		err = json.Unmarshal(buf, &account)
		if err != nil {
			panic(err)
		}
		opt := selectMenu(reader)
		displayItems(opt, &account)
	case "c":
		createFile(reader)
		FileOptionPrompt()
	case "d":
		services.DeleteFile()
	default:
		FileOptionPrompt()
	}
}

func itemFormat(name string, value float32) {
	fmt.Printf("%-25s%.2f\n", name+":", value)
}

func displayInputItems(input *domains.Input) {
	income := &input.Income
	fmt.Printf("Your Input\n\n")
	fmt.Printf("Income %s\n", strings.Repeat("=", 22))
	itemFormat(income.Salary.Name, income.Salary.Value)
	itemFormat(income.Tip.Name, income.Tip.Value)
	itemFormat(income.Bonus.Name, income.Bonus.Value)
	itemFormat(income.Commission.Name, income.Commission.Value)
	itemFormat(income.Other.Name, income.Other.Value)

	fmt.Println()
	other := &input.Other
	fmt.Printf("Other %s\n", strings.Repeat("=", 23))
	itemFormat(other.Transferred.Name, other.Transferred.Value)
	itemFormat(other.Interest.Name, other.Interest.Value)
	itemFormat(other.Dividend.Name, other.Dividend.Value)
	itemFormat(other.Gift.Name, other.Gift.Value)
	itemFormat(other.Refund.Name, other.Refund.Value)
	itemFormat(other.Installment.Name, other.Installment.Value)
	itemFormat(other.Balance.Name, other.Balance.Value)
}

func displayOutputItems(output *domains.Output) {
	fmt.Println("Your Output")
	fmt.Println("Outcome")
}

func displayItems(opt string, account *domains.Account) {
	// output := account.Output
	switch opt {
	case "1":
		displayInputItems(&account.Input)
	case "2":
		fmt.Println("Show Output")
	}
}

func selectMenu(r *bufio.Reader) string {
	fmt.Println("What do you want to see?")
	fmt.Println("1) input")
	fmt.Println("2) output")
	opt, _ := GetInput(": ", r)
	return opt
}

func selectFile(files *[]fs.DirEntry, r *bufio.Reader) int {
	opt, _ := GetInput("Please select a file: ", r)
	index, err := services.SelectFile(opt, files, r)
	if err != nil {
		fmt.Println(err)
		return selectFile(files, r)
	}
	return index
}

func createFile(r *bufio.Reader) {
	opt, _ := GetInput("Please type a file name you want to create : ", r)
	err := services.CreateFile(opt, r)
	if err != nil {
		fmt.Println(err)
		createFile(r)
	}
}
