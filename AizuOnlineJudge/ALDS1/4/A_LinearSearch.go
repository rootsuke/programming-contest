package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func NextString() string {
	sc.Scan()
	return sc.Text()
}

func NextInt() int {
	n, _ := strconv.Atoi(NextString())
	return n
}

func NextIntSlice(size int) []int {
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = NextInt()
	}
	return res
}

func PrintIntSlice(slice []int) {
	fmt.Println(strings.Trim(fmt.Sprint(slice), "[]"))
}

func LinearSearch(slice []int, target int) int {
	n := len(slice)

	// 番兵
	slice = append(slice, target)

	i := 0

	for slice[i] != target {
		i++
	}

	if i == n {
		return -1
	} else {
		return i
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N := NextInt()
	S := NextIntSlice(N)

	Q := NextInt()
	T := NextIntSlice(Q)

	cnt := 0
	for i := 0; i < Q; i++ {
		res := LinearSearch(S, T[i])
		if res >= 0 {
			cnt++
		}
	}

	fmt.Println(cnt)
}
