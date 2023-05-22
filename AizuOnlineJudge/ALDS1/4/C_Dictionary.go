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

// 1次ハッシュ
func primaryHash(key int, n int) int {
	return key % n
}

// 2次ハッシュ
func secondaryHash(key int, n int) int {
	// 再計算したハッシュ値が同じにならないように
	// 1以上かつハッシュテーブルの長さ未満の値を返す
	return 1 + (key % (n - 1))
}

func calcHash(key int, i int, n int) int {
	return (primaryHash(key, n) + i*secondaryHash(key, n)) % n
}

func Insert(hashTable []int, key int) int {
	n := len(hashTable)
	i := 0

	for i < n {
		hash := calcHash(key, i, n)
		if hashTable[hash] == 0 {
			// 格納しそのハッシュ値を返す
			hashTable[hash] = key
			return hash
		} else if hashTable[hash] == key {
			// すでに同じ値が格納されていればそのハッシュ値を返す
			return hash
		} else {
			// 衝突した場合はハッシュ値を再計算する
			i++
		}
	}

	return -1
}

func Search(hashTable []int, key int) int {
	n := len(hashTable)
	i := 0
	for i < n {
		hash := calcHash(key, i, n)
		if hashTable[hash] == key {
			// 見つかったらハッシュ値を返す
			return hash
		} else if hashTable[hash] == 0 {
			// 空だった場合は格納されていないと判定し-1を返す
			return -1
		} else {
			// 違う値が格納されていた場合はハッシュ値を再計算する
			i++
		}
	}

	return -1
}

func convertToKey(str string) int {
	sum := 0
	digit := 1
	for _, v := range str {
		sum += digit * int(v-'A'+1)
		digit *= 5
	}
	return sum
}

const HASH_TABLE_LENGTH = 1000003

func main() {
	scanner := NewScanner()
	writer := bufio.NewWriter(os.Stdout)
	N := scanner.nextInt()

	hashTable := [HASH_TABLE_LENGTH]int{}

	for i := 0; i < N; i++ {
		cmd := scanner.nextStr()
		str := scanner.nextStr()
		key := convertToKey(str)
		if cmd == "insert" {
			Insert(hashTable[:], key)
		} else {
			res := Search(hashTable[:], key)
			if res >= 0 {
				fmt.Fprintln(writer, "yes")
			} else {
				fmt.Fprintln(writer, "no")
			}
		}
	}
	writer.Flush()
}
