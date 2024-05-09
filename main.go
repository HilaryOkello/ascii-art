package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"ascii-art/ascii"
)

func main() {
	lenArgs := len(os.Args)
	if lenArgs < 2 || lenArgs > 3 {
		fmt.Printf("Incorrect no. of arguments.\n" +
			"Expects: \"go run . <string> | cat -e\"\n" +
			"or\n" +
			"\"go run . <string> <banner name> | cat -e\"\n")
		return
	}

	str := os.Args[1]
	fmt.Println([]byte(str))
	str = strings.ReplaceAll(str, "\\b", "\b")
	str = strings.ReplaceAll(str, "\\t", "    ")
	str = strings.ReplaceAll(str, "\n", "\\n")
	str, err := ascii.IsPrintableAscii(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := "standard.txt"
	if lenArgs == 3 {
		fileName = os.Args[2] + ".txt"
	}
	errFile := ascii.CheckFileValidity(fileName)
	if errFile != nil {
		fmt.Println(errFile)
		return
	}

	filePath := os.DirFS("./banner")
	contentByte, err := fs.ReadFile(filePath, fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(contentByte) == 0 {
		fmt.Println("Banner file is empty")
		return
	}
	er := ascii.CheckFileTamper(fileName, contentByte)
	if er != nil {
		fmt.Println(er)
		return
	}

	contentString := string(contentByte[1:])
	if fileName == "thinkertoy.txt" {
		contentString = strings.ReplaceAll(string(contentByte[2:]), "\r\n", "\n")
	}
	contentSlice := strings.Split(contentString, "\n\n")
	words := strings.Split(str, "\\n")
	count := 0
	for _, str := range words {
		if str == "" {
			count++
			if count < len(words) {
				fmt.Println()
			}
		} else {
			ascii.PrintAscii(str, contentSlice, 0)
		}
	}
}
