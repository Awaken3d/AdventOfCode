package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	lock  sync.Mutex
	wg    sync.WaitGroup
	wrong int
)

const workers = 10

func worker(pass chan []string) {
	defer wg.Done()
	for p := range pass {
		words := make(map[string]bool)
		for _, word := range p {

			letters := make([]int, 26)
			for _, letter := range word {
				letters[letter-'a']++
			}

			encoded := strings.Trim(strings.Replace(fmt.Sprint(letters), " ", ",", -1), "[]")

			if ok := words[encoded]; ok {
				lock.Lock()
				wrong++
				lock.Unlock()
				break
			}
			words[encoded] = true
		}
	}

}

func main() {
	f, _ := os.Open("input")
	scan := bufio.NewScanner(f)
	pass := make(chan []string)
	count := 0
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(pass)
	}

	for scan.Scan() {
		line := scan.Text()
		count++
		phrase := strings.Fields(line)
		pass <- phrase
	}

	close(pass)
	wg.Wait()
	log.Println(count - wrong)
}
