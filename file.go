package main

import (
	"bufio"
	"fmt"
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
		services.LoadFile(selectedFile.Name())
	case "c":
		createFile(reader)
	case "d":
		services.DeleteFile()
	default:
		FileOptionPrompt()
	}
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
