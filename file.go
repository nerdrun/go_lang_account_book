package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strconv"
)

func lookupFolder() (*[]fs.DirEntry, error) {
	entries, err := os.ReadDir("./")
	if err != nil {
		return nil, fmt.Errorf("can't read the directory : %v", err)
	}
	var accbooks []fs.DirEntry
	for _, entry := range entries {
		extension := path.Ext(entry.Name())
		if extension == ".accbook" {
			accbooks = append(accbooks, entry)
		}
	}
	if len(accbooks) == 0 {
		fmt.Println("Please create a file first")
		fileOptionPrompt()
	}

	for index, acc := range accbooks {
		fmt.Printf("%d) %s\n", index, acc.Name())
	}

	return &accbooks, nil
}

func selectFile(files *[]fs.DirEntry, r *bufio.Reader) int {
	opt, _ := getInput("Please select a file: ", r)
	index, err := strconv.Atoi(opt)
	if err != nil {
		fmt.Println("Please enter only digit")
		return selectFile(files, r)
	}
	if index > len(*files) || index < 0 {
		fmt.Println("Please select a correct file number")
		return selectFile(files, r)
	}
	return index
}

func loadFile(path string) (*os.File, error) {
	if path == "" {
		return nil, fmt.Errorf("the path of a file must not be empty")
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err != nil {
		return nil, fmt.Errorf("opening file errror : %w", err)
	}

	return f, nil
}

// func loadFile(file *fs.DirEntry) {
// 	f, err := os.OpenFile((*file).Name(), os.O_APPEND|os.O_CREATE, os.ModePerm)
// 	if err != nil {
// 		fmt.Printf("Opening file error : %v\n", err)
// 	}
// 	defer f.Close()

// 	buf := make([]byte, 1024)
// 	n, _ := f.Read(buf)
// 	var data Category
// 	// if err := json.Unmarshal(buf, &data); err != nil {
// 	// 	panic(err)
// 	// }
// 	fmt.Println("comparison")
// 	fmt.Println(buf)
// 	fmt.Println(buf[:n])
// 	err = json.Unmarshal(buf[:n], &data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(data)
// 	// fmt.Println(data)
// }

func createFile() {
	f, err := loadFile("input.init")
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Name())
	fmt.Println(f)

	defer f.Close()

	buf := make([]byte, 30000)
	n, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(string(buf[:n]))
}

func deleteFile() {
}
