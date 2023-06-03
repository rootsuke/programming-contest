package main

import (
	"bufio"
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

func removeDuplicates(nums []int) int {
	targetIdx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[targetIdx] = nums[i]
			targetIdx++
		}
	}

	return targetIdx
}

func main() {
	sc := NewScanner()
	N := sc.nextInt()
	nums := sc.nextIntSlice(N)

	removeDuplicates(nums)
}
