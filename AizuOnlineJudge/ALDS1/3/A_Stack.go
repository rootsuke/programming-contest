package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Scanner struct {
	sc *bufio.Scanner
}

func NewScanner() *Scanner {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &Scanner{sc}
}

type Stack struct {
	stack []int
}

func (s *Stack) Push(value int) {
	s.stack = append(s.stack, value)
}

func (s *Stack) Pop() int {
	lastValue := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return lastValue
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	str := sc.Text()
	strs := strings.Split(str, " ")
	stack := &Stack{}

	for _, v := range strs {
		switch v {
		case "+":
			b := stack.Pop()
			a := stack.Pop()
			stack.Push(a + b)
		case "-":
			b := stack.Pop()
			a := stack.Pop()
			stack.Push(a - b)
		case "*":
			b := stack.Pop()
			a := stack.Pop()
			stack.Push(a * b)
		default:
			num, _ := strconv.Atoi(v)
			stack.Push(num)
		}
	}

	fmt.Println(stack.Pop())
}
