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

func BubbleSort(slice []int) int {
	n := len(slice)
	// 値を入れ替えたか管理する変数
	// ループを起動させるために初回だけtrueにしておく
	swapped := true
	cnt := 0

	// 最後まで走査して入れ替えが発生しなければ終了
	for swapped {
		swapped = false

		for i := 0; i < n-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				swapped = true
				cnt++
			}
		}

		// 最後まで走査する度に、最後の値がソート済みとなる
		// 不要な走査を減らすためにnをデクリメントして走査範囲を狭めていく
		n--
	}

	return cnt
}

func main() {
	sc.Split(bufio.ScanWords)

	N := NextInt()
	A := NextIntSlice(N)

	cnt := BubbleSort(A)

	PrintIntSlice(A)
	fmt.Println(cnt)
}
