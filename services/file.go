package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strconv"

	domains "account.com/test/domains"
)

func LookupFolder() (*[]fs.DirEntry, error) {
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
		return nil, fmt.Errorf("please create a file first")
	}

	for index, acc := range accbooks {
		fmt.Printf("%d) %s\n", index, acc.Name())
	}

	return &accbooks, nil
}

func SelectFile(opt string, files *[]fs.DirEntry, r *bufio.Reader) (int, error) {
	index, err := strconv.Atoi(opt)
	if err != nil {
		return 0, fmt.Errorf("please enter only digit")
	}
	if index > len(*files) || index < 0 {
		return 0, fmt.Errorf("please select a correct file number")
	}
	return index, nil
}

func LoadFile(path string) (*os.File, error) {
	if path == "" {
		return nil, fmt.Errorf("the path of a file must not be empty")
	}

	f, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("opening file errror : %w", err)
	}

	return f, nil
}

func CreateFile(opt string, r *bufio.Reader) error {
	fileName := fmt.Sprintf("%s.accbook", opt)
	_, err := LoadFile(fileName)
	if err == nil {
		return fmt.Errorf("the file already exists")
	}
	account := domains.GenerateAccount()

	buf, err := json.Marshal(account)

	if err != nil {
		return fmt.Errorf("an error got occurred while converting input json to byte")
	}

	err = os.WriteFile(fileName, buf, os.ModePerm)

	if err != nil {
		return fmt.Errorf("can't write a file")
	}
	return nil
}

func DeleteFile() {
}
