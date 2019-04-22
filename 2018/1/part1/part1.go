package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	freq := 0
	// Open the file.
	f, _ := os.Open("input")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		freq += num
		//fmt.Println(line)
	}
	fmt.Println(freq)
}
