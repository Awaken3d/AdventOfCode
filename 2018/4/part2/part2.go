package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	layout = "2006-01-02 15:04"
	events = []event{}
	lock   sync.Mutex
)

type event struct {
	t time.Time
	e string
}

type guard struct {
	total int
	mins  [60]int
}

func breakup(line string) {

	line = strings.TrimPrefix(line, "[")
	elements := strings.Split(line, "]")

	t, err := time.Parse(layout, elements[0])
	if err != nil {
		log.Fatalf("error while parsing time %v", err)
	}

	events = append(events, event{t: t, e: strings.TrimSpace(elements[1])})

}

func generateGuards() map[string]*guard {
	guards := make(map[string]*guard)
	cur := ""
	var sleep time.Time
	for _, event := range events {
		if strings.Contains(event.e, "#") {
			cur = event.e
		} else if strings.Contains(event.e, "falls") {
			sleep = event.t
		} else {
			if _, ok := guards[cur]; !ok {
				guards[cur] = &guard{}
			}

			guards[cur].total += event.t.Minute() - sleep.Minute()
			for i := sleep.Minute(); i < event.t.Minute(); i++ {
				guards[cur].mins[i]++
			}
		}

	}
	return guards
}

func findGuardMinute(guards map[string]*guard) (string, int) {
	max := -1
	idGuard := ""
	minute := -1

	for key, val := range guards {
		for min, freq := range val.mins {
			if freq > max {
				max = freq
				idGuard = key
				minute = min
			}
		}
	}
	return idGuard, minute
}

func main() {

	f, _ := os.Open("input")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		breakup(scanner.Text())
	}

	sort.Slice(events, func(i, j int) bool { return events[i].t.Before(events[j].t) })

	guards := generateGuards()

	id, minute := findGuardMinute(guards)

	idNum := strings.Split(id, " ")[1]
	idNum = strings.TrimPrefix(idNum, "#")

	num, _ := strconv.Atoi(idNum)
	total := strconv.Itoa(num * minute)
	fmt.Println(total)
}
