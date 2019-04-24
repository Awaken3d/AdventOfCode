package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	f, _ := os.Open("input")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
	}

	found := true
	for {
		if !found {
			break
		}
		found = false
		nline := line
		for i := 0; i < len(line)-1; i++ {
			if ((unicode.IsLower(rune(line[i])) && unicode.IsUpper(rune(line[i+1]))) || (unicode.IsUpper(rune(line[i])) && unicode.IsLower(rune(line[i+1])))) && unicode.ToLower(rune(line[i])) == unicode.ToLower(rune(line[i+1])) {
				found = true
				nline = strings.ReplaceAll(nline, string(line[i])+string(line[i+1]), "")
			}
		}
		line = nline
	}

	fmt.Println(len(line))

}
