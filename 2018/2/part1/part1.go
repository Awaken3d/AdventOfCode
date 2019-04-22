package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

var (
	wg           sync.Mutex
	twos, threes int
)

func work(ids chan string, maps chan map[rune]int, done chan bool) {

	for id := range ids {
		chars := make(map[rune]int)
		for _, c := range id {
			//log.Println("rune found ", c)
			chars[c]++
		}
		maps <- chars

	}

	done <- true

}

func counter(maps chan map[rune]int, done chan bool) {
	for m := range maps {
		two, three := false, false
		for _, value := range m {

			if two && three {
				break
			}
			if value == 2 && !two {
				wg.Lock()
				twos++
				wg.Unlock()
				two = true
			}
			if value == 3 && !three {
				wg.Lock()
				threes++
				wg.Unlock()
				three = true
			}
		}
	}
	done <- true
}

func main() {
	doneW, doneC := make(chan bool), make(chan bool)
	ids := make(chan string)
	maps := make(chan map[rune]int)
	workers := 10
	counters := 10
	for i := 0; i < workers; i++ {
		go work(ids, maps, doneW)
	}
	for i := 0; i < counters; i++ {
		go counter(maps, doneC)
	}
	// Open the file.
	f, _ := os.Open("input1")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.

	for scanner.Scan() {
		id := scanner.Text()
		ids <- id
	}
	close(ids)

	for i := 0; i < workers; i++ {
		<-doneW
	}
	close(maps)
	close(doneW)
	for i := 0; i < counters; i++ {
		<-doneC
	}
	close(doneC)
	log.Printf("twos is %d and threes is %d", twos, threes)
	log.Printf("checksum is %d", twos*threes)

}
