package ascii

import (
	"fmt"
	"os"
	"strings"
)

func IsPrintableAscii(str string) (string, error) {
	var nonPritables string
	var foundEscapes string
	var result string
	errMessage := ": Not within the printable ascii range"
	for index, char := range str {
		escapes := "avrf"
		backspace := char == '\b'
		var previous byte
		var next byte

		if index > 0 {
			previous = str[index-1]
		}
		if index < len(str)-1 {
			next = str[index+1]
		}

		NextIsAnEscapeLetter := strings.ContainsAny(string(next), escapes)
		isAnEscape := (char == '\\' && NextIsAnEscapeLetter && previous != '\\')
		isNonPrintable := ((char < ' ' || char > '~') && char != '\n')

		if isAnEscape {
			foundEscapes += string(next)
		} else if backspace {
			result = result[:len(result)-1]
		} else if isNonPrintable {
			nonPritables += string(char)
		} else {
			result += string(char)
		}
	}

	if foundEscapes != "" {
		escSlash := ""
		for _, es := range foundEscapes {
			escSlash += "\\" + string(es)
		}
		return "", fmt.Errorf("%s%s", escSlash, errMessage)
	} else if nonPritables != "" {
		return "", fmt.Errorf("%s%s", nonPritables, errMessage)
	}

	return result, nil
}

func CheckFileValidity(fileName string) error {
	path := "./banner"
	openPath, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer openPath.Close()
	filenames, err := openPath.Readdirnames(0)
	if err != nil {
		return fmt.Errorf("error name")
	}
	fileNameString := strings.Join(filenames, " ")
	if !strings.Contains(fileNameString, fileName) {
		return fmt.Errorf("%s is not a valid banner style\n"+
			"Try \"standard\", \"shadow\", or \"thinkertoy\"",
			fileName[:len(fileName)-4])
	}
	return nil
}

func CheckFileTamper(fileName string, content []byte) error {
	errMessage := "is tampered"
	lengthContent := len(content)

	if fileName == "standard.txt" && lengthContent != 6623 {
		return fmt.Errorf("%s%s", fileName, errMessage)
	} else if fileName == "thinkertoy.txt" && lengthContent != 5558 {
		return fmt.Errorf("%s%s", fileName, errMessage)
	} else if fileName == "shadow.txt" && lengthContent != 7465 {
		return fmt.Errorf("%s%s", fileName, errMessage)
	}

	return nil
}
