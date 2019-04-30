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
		min, max := int(^uint(0)>>1), -int(^uint(0)>>1)-1

		for _, numStr := range slice {
			num, _ := strconv.Atoi(numStr)
			if num < min {
				min = num
			}

			if num > max {
				max = num
			}
		}
		diff <- max - min
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
