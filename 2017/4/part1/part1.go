package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	lock  sync.Mutex
	wrong int
)

const workers = 10

func worker(pass chan []string) {

	for p := range pass {
		words := make(map[string]bool)

		for _, word := range p {
			if ok := words[word]; ok {
				lock.Lock()
				wrong++
				lock.Unlock()
				break
			}
			words[word] = true
		}
	}

}

func main() {
	f, _ := os.Open("input")
	scan := bufio.NewScanner(f)
	pass := make(chan []string)
	count := 0
	for i := 0; i < workers; i++ {
		go worker(pass)
	}

	for scan.Scan() {
		line := scan.Text()
		count++
		phrase := strings.Fields(line)
		pass <- phrase
	}

	close(pass)

	log.Println(count - wrong)
}
