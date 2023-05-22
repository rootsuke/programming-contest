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

func PrintIntSlice(slice []int) {
	fmt.Println(strings.Trim(fmt.Sprint(slice), "[]"))
}

var MAX = 10001

func CountingSort(slice []int) []int {
	C := make([]int, MAX)

	for i := 0; i < len(slice); i++ {
		C[slice[i]]++
	}

	for i := 1; i < len(C); i++ {
		C[i] = C[i] + C[i-1]
	}

	B := make([]int, len(slice))

	for i := len(slice) - 1; i >= 0; i-- {
		B[C[slice[i]]-1] = slice[i]
		C[slice[i]]--
	}

	return B
}

func main() {
	sc := NewScanner()

	N := sc.nextInt()
	S := sc.nextIntSlice(N)

	res := CountingSort(S)

	PrintIntSlice(res)
}
