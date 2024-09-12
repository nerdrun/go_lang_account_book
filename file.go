package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"

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

		var account services.Account
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

func displayItems(opt string, account *services.Account) {
	switch opt {
	case "1":
		fmt.Println("Show Input")
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
