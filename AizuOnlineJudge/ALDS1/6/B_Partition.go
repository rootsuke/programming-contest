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

func Partition(slice []int, p int, r int) int {
	x := slice[r]
	i := p - 1

	for j := p; j < r; j++ {
		if slice[j] <= x {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[r] = slice[r], slice[i+1]

	return i + 1
}

func main() {
	sc := NewScanner()

	N := sc.nextInt()
	S := sc.nextIntSlice(N)

	res := Partition(S, 0, len(S)-1)

	for i, v := range S {
		if i == res {
			fmt.Printf("[%v] ", v)
		} else if i == len(S)-1 {
			fmt.Printf("%v\n", v)
		} else {
			fmt.Printf("%v ", v)
		}
	}
}
