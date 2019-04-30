package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	elements []int
}

func (s *stack) push(el int) {
	s.elements = append(s.elements, el)
}

func (s *stack) pop() int {
	num := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return num
}

func (s *stack) isEmpty() bool {
	return len(s.elements) == 0
}

func convToIntSlice(s []string) []int {
	result := []int{}
	for _, str := range s {
		num, _ := strconv.Atoi(str)
		result = append(result, num)
	}

	return result
}

// func parse(c, m, index int) int {

// 	for i := 0; i < c; i++ {
// 		child := index + 1
// 		meta := index + 2
// 		index += 2
// 		//parse(nums[child], nums[meta], index)
// 	}
// 	return 0
// }
func main() {

	f, _ := os.Open("input")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
	}

	strNums := strings.Split(line, " ")
	nums := convToIntSlice(strNums)
	//log.Println(nums)

	children, meta := &stack{}, &stack{}
	children.push(nums[0])
	meta.push(nums[1])
	index := 2
	result := 0
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(strconv.Itoa(result))
		}
	}()
	for (!children.isEmpty() || !meta.isEmpty()) && index < len(nums) {
		//for index < len(nums) {
		childrenLen := 0
		if !children.isEmpty() {
			childrenLen = children.pop()
		}
		metaLen := 0
		if !meta.isEmpty() {
			metaLen = meta.pop()
		}
		//fmt.Println(children.elements)
		// if children.isEmpty() && !meta.isEmpty() {
		// 	fmt.Println(metaLen)
		// }
		met := true
		i := 1
		//if !children.isEmpty() {
		for i = 0; i < childrenLen; i++ {
			fmt.Println(strconv.Itoa(index))
			fmt.Println(strconv.Itoa(len(nums)))
			if nums[index] > 0 {
				if i < childrenLen-1 {
					children.push(childrenLen - i - 1)
				}
				children.push(nums[index])
				index++
				meta.push(metaLen)
				meta.push(nums[index])
				index++
				met = false
			} else {
				index++
				if i != childrenLen-1 {
					children.push(childrenLen - i - 1)
				}
				meta.push(metaLen)
				metaLen = nums[index]
				index++
				break
			}
		}

		//}
		for j := 0; met && j < metaLen; j++ {
			fmt.Println("adding " + strconv.Itoa(nums[index]))
			result += nums[index]
			index++
		}

	}

	fmt.Println(strconv.Itoa(result))

}
