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

func findGuard(guards map[string]*guard) (g *guard, id string) {
	max := -1
	idGuard := ""
	var actual *guard

	for key, val := range guards {
		//fmt.Println("for guard " + key + " total is " + strconv.Itoa(guards[key].total))
		if val.total > max {
			max = val.total
			idGuard = key
			actual = val
		}
	}

	return actual, idGuard
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

	g, id := findGuard(guards)

	m := 0
	ind := 0
	for index, val := range g.mins {
		if val > m {
			m = val
			ind = index
		}
	}

	idNum := strings.Split(id, " ")[1]
	idNum = strings.TrimPrefix(idNum, "#")

	num, _ := strconv.Atoi(idNum)
	total := strconv.Itoa(num * ind)
	fmt.Println(total)
}
