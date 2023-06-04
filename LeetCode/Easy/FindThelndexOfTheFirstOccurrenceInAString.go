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

func (s *Scanner) nextStrSlice(size int) []string {
	res := make([]string, size)
	for i := 0; i < size; i++ {
		res[i] = s.nextStr()
	}
	return res
}

func strStr(haystack string, needle string) int {
	firstOccurrenceIdx := 0
	j := 0

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[j] {
			if j == len(needle)-1 {
				return firstOccurrenceIdx
			}

			j++
		} else {
			i = firstOccurrenceIdx
			j = 0
			firstOccurrenceIdx++
		}
	}

	return -1
}

func main() {
	sc := NewScanner()
	haystack := sc.nextStr()
	needle := sc.nextStr()

	fmt.Println(strStr(haystack, needle))
}
