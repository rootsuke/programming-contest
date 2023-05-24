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

type Process struct {
	name string
	time int
}

type Queue struct {
	queue []Process
}

func (q *Queue) Enqueue(process *Process) {
	q.queue = append(q.queue, *process)
}

func (q *Queue) Dequeue() *Process {
	process := q.queue[0]
	q.queue = q.queue[1:]

	return &process
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}

func main() {
	sc := NewScanner()

	N := sc.nextInt()
	quantum := sc.nextInt()

	queue := &Queue{}
	for i := 0; i < N; i++ {
		name := sc.nextStr()
		time := sc.nextInt()
		process := &Process{name: name, time: time}
		queue.Enqueue(process)
	}

	totalSpentTime := 0
	for !queue.IsEmpty() {
		currentProcess := queue.Dequeue()
		if currentProcess.time > quantum {
			totalSpentTime += quantum
			currentProcess.time -= quantum
			queue.Enqueue(currentProcess)
		} else {
			totalSpentTime += currentProcess.time
			fmt.Printf("%v %v\n", currentProcess.name, totalSpentTime)
			currentProcess.time = 0
		}
	}
}
