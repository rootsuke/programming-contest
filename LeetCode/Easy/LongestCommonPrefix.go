package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func longestCommonPrefix(strs []string) string {
	ans := ""

	for i := 0; i < len(strs[0]); i++ {
		current := string(strs[0][i])

		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || current != string(strs[j][i]) {
				return ans
			}
		}

		ans += current
	}

	return ans
}

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
func (s *Scanner) nextStrSlice(size int) []string {
	res := make([]string, size)
	for i := 0; i < size; i++ {
		res[i] = s.nextStr()
	}
	return res
}

func main() {
	sc := NewScanner()
	N := sc.nextInt()
	strs := sc.nextStrSlice(N)

	fmt.Println(longestCommonPrefix(strs))
}
