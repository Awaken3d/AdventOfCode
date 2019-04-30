package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	checksum int
	wg       sync.WaitGroup
)

const (
	workers = 10
)

func worker(slices chan []string, diff chan int) {
	defer wg.Done()
	for slice := range slices {
		for i := 0; i < len(slice); i++ {
			num1, _ := strconv.Atoi(slice[i])
			for j := i + 1; j < len(slice); j++ {
				num2, _ := strconv.Atoi(slice[j])
				if num1%num2 == 0 {
					diff <- num1 / num2
				} else if num2%num1 == 0 {
					diff <- num2 / num1
				}
			}
		}
	}
}

func receiver(diff chan int, done chan struct{}) {
	for d := range diff {
		checksum += d
	}
	done <- struct{}{}
}
func main() {

	slices, diff, done := make(chan []string), make(chan int), make(chan struct{})

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(slices, diff)
	}

	go receiver(diff, done)
	go func() {
		wg.Wait()
		close(diff)
	}()
	f, _ := os.Open("input")
	scan := bufio.NewScanner(f)

	for scan.Scan() {
		line := scan.Text()
		nums := strings.Fields(line)
		slices <- nums
	}
	close(slices)
	<-done
	log.Println(checksum)
}
