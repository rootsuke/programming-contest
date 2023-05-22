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

func ShellSort(slice []int) []int {
	n := len(slice)
	intervals := calcInterval(n)

	for i := 0; i < len(intervals); i++ {
		insertionSort(slice, intervals[i])
	}

	return slice
}

func calcInterval(n int) []int {
	var h = []int{1}

	for i := 4; i < n; i *= 3 + 1 {
		h = append([]int{i}, h...)
	}

	fmt.Println(len(h))
	PrintIntSlice(h)
	return h
}

var cnt = 0

func insertionSort(slice []int, interval int) {
	n := len(slice)

	for i := interval; i < n; i++ {
		crr := slice[i]
		j := i - interval

		for j >= 0 && crr < slice[j] {
			slice[j+interval] = slice[j]
			j -= interval
			cnt++
		}

		slice[j+interval] = crr
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N := NextInt()
	slice := []int{}
	for i := 1; i <= N; i++ {
		num := NextInt()
		slice = append(slice, num)
	}

	slice = ShellSort(slice)

	fmt.Println(cnt)

	for _, v := range slice {
		fmt.Println(v)
	}
}
