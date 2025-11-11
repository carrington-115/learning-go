package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	fmt.Println("Working with files")

	// Data elements
	path := "./content.txt"
	content := "Hello World\nI am Mark\nI am current learning GoLang\nIt is fun"
	FileOperation := Files{
		Content:    content,
		Path:       path,
		ThrowError: throwError,
	}

	// Writing the file
	file := FileOperation.CreateFile(FileOperation.Path)
	length, err := io.WriteString(file, FileOperation.Content)
	FileOperation.ThrowError(err)
	if reflect.TypeOf(length).Kind() != reflect.Int {
		fmt.Println("The file was not correctly created")
		file.Close()
	}

	// Reading the file
	fileData := FileOperation.ReadFile(path)
	fmt.Println("File data: ")
	fmt.Println(fileData)
	file.Close()
}

type Files struct {
	Content    string
	Path       string
	ThrowError func(err error)
}

func (f Files) CreateFile(path string) *os.File {
	file, err := os.Create(path)
	f.ThrowError(err)
	return file
}

func (f Files) ReadFile(path string) string {
	filebytes, err := os.ReadFile(path)
	f.ThrowError(err)
	return string(filebytes)
}

func throwError(err error) {
	if err != nil {
		panic(err)
	}
}
