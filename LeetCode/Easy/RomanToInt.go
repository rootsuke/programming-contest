package main

import (
	"bufio"
	"os"
)

func romanToInt(s string) int {
	sum := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'M':
			sum += 1000
		case 'D':
			sum += 500
		case 'L':
			sum += 50
		case 'V':
			sum += 5
		case 'C':
			if i == len(s)-1 {
				sum += 100
				break
			}

			switch s[i+1] {
			case 'M':
				sum += 900
				i++
			case 'D':
				sum += 400
				i++
			default:
				sum += 100
			}
		case 'X':
			if i == len(s)-1 {
				sum += 10
				break
			}

			switch s[i+1] {
			case 'C':
				sum += 90
				i++
			case 'L':
				sum += 40
				i++
			default:
				sum += 10
			}
		case 'I':
			if i == len(s)-1 {
				sum += 1
				break
			}

			switch s[i+1] {
			case 'X':
				sum += 9
				i++
			case 'V':
				sum += 4
				i++
			default:
				sum += 1
			}
		}
	}

	return sum
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

func main() {
	sc := NewScanner()
	str := sc.nextStr()

	romanToInt(str)
}
