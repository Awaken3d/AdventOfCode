package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("readFileWithReadString")

	// file, err := os.Open("input2")
	// defer file.Close()
	var file *os.File
	// if err != nil {
	// 	log.Fatalf("could not open file: %v", err)
	// }

	freq := 0
	freqHolder := make(map[int]bool)
	var line string
	var err error
	for {
		file, _ = os.Open("input2")
		reader := bufio.NewReader(file)
		b := false
		for !b {
			line, err = reader.ReadString('\n')

			if err == io.EOF {
				b = true
				continue
			}
			line = strings.TrimRight(line, "\n")
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Printf("error during conversion %v", err)
			}
			log.Printf("num read is %d \n", num)
			freq += num

			if _, ok := freqHolder[freq]; ok {
				log.Printf("result is %d", freq)
				os.Exit(0)
			}

			freqHolder[freq] = true
		}
	}

	return
}
