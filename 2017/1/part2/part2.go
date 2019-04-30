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

	jump, index, length, sum := len(num)/2, 0, len(num), 0

	for i := 0; i < length-1; i++ {
		if i+jump >= length {
			index = (jump - (length - 1 - i)) % length
		} else {
			index = i + jump
		}
		if num[i] == num[index] {
			sum += int(num[i] - '0')
		}
	}

	log.Println(sum)
}
