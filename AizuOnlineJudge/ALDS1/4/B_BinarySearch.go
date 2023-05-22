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
func BinarySearch(slice []int, target int) int {
	n := len(slice)
	left := 0
	right := n

	for left < right {
		mid := (right + left) / 2

		if slice[mid] == target {
			return mid
		} else if target < slice[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	sc.Split(bufio.ScanWords)

	N := NextInt()
	S := NextIntSlice(N)

	Q := NextInt()
	T := NextIntSlice(Q)

	cnt := 0
	for i := 0; i < Q; i++ {
		res := BinarySearch(S, T[i])
		if res >= 0 {
			cnt++
		}
	}

	fmt.Println(cnt)
}
