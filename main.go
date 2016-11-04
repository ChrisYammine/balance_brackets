package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []string

func (s stack) Push(str string) stack {
	return append(s, str)
}

func (s stack) Pop() (stack, string, error) {
	l := len(s)
	if l == 0 {
		return s, "", errors.New("Empty Stack")
	}

	return s[:l-1], s[l-1], nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var arrays = make([][]string, 0)

	for i := 0; i < n; i++ {
		scanner.Scan()
		chars := strings.Split(scanner.Text(), "")
		arrays = append(arrays, chars)
	}

	for _, array := range arrays {
		if isBalanced(array) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func isBalanced(chars []string) bool {
	closers := make(stack, 0)

	for _, char := range chars {
		switch char {
		case "{":
			closers = closers.Push("}")
		case "[":
			closers = closers.Push("]")
		case "(":
			closers = closers.Push(")")
		case "}":
			m_closers, popped, err := closers.Pop()
			if err != nil {
				return false
			}
			if popped != char {
				return false
			}
			closers = m_closers
		case "]":
			m_closers, popped, err := closers.Pop()
			if err != nil {
				return false
			}
			if popped != char {
				return false
			}
			closers = m_closers
		case ")":
			m_closers, popped, err := closers.Pop()
			if err != nil {
				return false
			}
			if popped != char {
				return false
			}
			closers = m_closers
		}
	}

	if len(closers) == 0 {
		return true
	}
	return false
}
