package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func PrintIntSlice(slice []int) {
	fmt.Println(strings.Trim(fmt.Sprint(slice), "[]"))
}

type Node struct {
	key  int
	next *Node
	prev *Node
}

func newNode() *Node {
	node := &Node{
		next: nil,
		prev: nil,
	}

	node.next = node
	node.prev = node

	return node
}

type linkedList struct {
	head *Node
}

func newLinkedList() *linkedList {
	node := newNode()
	return &linkedList{head: node}
}

func (dll *linkedList) insert(key int) {
	next := dll.head.next
	node := &Node{key: key, next: next, prev: dll.head}
	dll.head.next = node
	next.prev = node
}

func (dll *linkedList) delete(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (dll *linkedList) deleteKey(key int) {
	for n := dll.head.next; n != dll.head; n = n.next {
		if n.key == key {
			dll.delete(n)
			return
		}
	}
}

func (dll *linkedList) deleteFirst() {
	dll.delete(dll.head.next)
}

func (dll *linkedList) deleteLast() {
	dll.delete(dll.head.prev)
}

func printList(dll *linkedList) {
	slice := []int{}
	for n := dll.head.next; n != dll.head; n = n.next {
		slice = append(slice, n.key)
	}
	PrintIntSlice(slice)
}

func main() {
	sc := NewScanner()
	N := sc.nextInt()

	dll := newLinkedList()

	for i := 0; i < N; i++ {
		command := sc.nextStr()
		switch command {
		case "insert":
			key := sc.nextInt()
			dll.insert(key)
		case "delete":
			key := sc.nextInt()
			dll.deleteKey(key)
		case "deleteFirst":
			dll.deleteFirst()
		case "deleteLast":
			dll.deleteLast()
		}
	}

	printList(dll)
}
