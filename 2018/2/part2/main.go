package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/agnivade/levenshtein"
)

func main() {

	// Open the file.
	f, _ := os.Open("input")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	ids := []string{}
	for scanner.Scan() {
		id := scanner.Text()
		ids = append(ids, id)
	}

	var s1, s2 string
	for i := range ids {
		s1 = ids[i]
		for j := range ids {

			s2 = ids[j]
			distance := levenshtein.ComputeDistance(s1, s2)

			if distance == 1 {
				fmt.Printf("The distance between %s and %s is %d.\n", s1, s2, distance)
				os.Exit(0)
			}
		}
	}

}
