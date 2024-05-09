package ascii

import (
	"fmt"
	"strings"
)

func PrintAscii(str string, contentSlice []string, index int) {
	if index == 8 {
		return
	}

	for _, char := range str {
		character := contentSlice[int(char)-32]
		character = strings.ReplaceAll(character, "\r\n", "\n")
		lines := strings.Split(character, "\n")
		fmt.Printf(lines[index])
	}
	fmt.Println()
	PrintAscii(str, contentSlice, index+1)
}
