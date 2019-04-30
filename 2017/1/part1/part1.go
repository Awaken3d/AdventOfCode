package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	f, _ := os.Open("input")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	num := ""
	for scanner.Scan() {
		num = scanner.Text()
	}

	last := num[len(num)-1]
	num = string(last) + num

	sum := 0
	for i := 0; i < len(num)-1; i++ {
		if num[i] == num[i+1] {
			sum += int(num[i] - '0')
		}
	}

	log.Println(sum)
}
