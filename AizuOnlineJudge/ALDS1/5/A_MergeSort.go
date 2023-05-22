package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func (s *Scanner) nextStr() string {
	s.sc.Scan()
	return s.sc.Text()
}

func (s *Scanner) nextInt() int {
	i, e := strconv.Atoi(s.nextStr())
	if e != nil {
		panic(e)
	}
	return i
}
func (s *Scanner) nextIntSlice(size int) []int {
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = s.nextInt()
	}
	return res
}

// index以降のsliceの要素を使って整数mを作れる場合はtrueを返す
func ExhaustiveSearch(slice []int, index int, m int) bool {
	if m == 0 {
		return true
	}
	if index >= len(slice) {
		return false
	}

	res := ExhaustiveSearch(slice, index+1, m) || ExhaustiveSearch(slice, index+1, m-slice[index])

	return res
}

func main() {
	sc := NewScanner()

	N := sc.nextInt()
	S := sc.nextIntSlice(N)

	Q := sc.nextInt()
	T := sc.nextIntSlice(Q)

	for _, m := range T {
		res := ExhaustiveSearch(S, 0, m)

		if res {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
