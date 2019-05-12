package main

import (
	"bufio"
	"log"
	"os"
	"runtime/trace"
	"strconv"
	"strings"
	"sync"
)

var (
	checksum int
	wg       sync.WaitGroup
	lock     sync.Mutex
)

const (
	workers = 3
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

func workerNoRec(slices chan []string, diff chan int) {
	defer wg.Done()
	for slice := range slices {
		for i := 0; i < len(slice); i++ {
			num1, _ := strconv.Atoi(slice[i])
			for j := i + 1; j < len(slice); j++ {
				num2, _ := strconv.Atoi(slice[j])
				if num1%num2 == 0 {
					lock.Lock()
					checksum += num1 / num2
					lock.Unlock()
					//diff <- num1 / num2
				} else if num2%num1 == 0 {
					lock.Lock()
					checksum += num2 / num1
					lock.Unlock()
					//diff <- num2 / num1
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

	trace.Start(os.Stdout)
	defer trace.Stop()

	slices, diff := make(chan []string, 17), make(chan int, 17)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go workerNoRec(slices, diff)
	}

	// go receiver(diff, done)
	// go func() {
	// 	wg.Wait()
	// 	close(diff)
	// }()
	f, _ := os.Open("input")
	scan := bufio.NewScanner(f)

	for scan.Scan() {
		line := scan.Text()
		nums := strings.Fields(line)
		slices <- nums
	}
	close(slices)
	close(diff)
	wg.Wait()
	//<-done
	log.Println(checksum)
}
