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

func (s *Scanner) nextIntSlice(size int) []int {
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = s.nextInt()
	}
	return res
}

func PrintIntSlice(slice []int) {
	fmt.Println(strings.Trim(fmt.Sprint(slice), "[]"))
}

type Node struct {
	parent       int
	leftChild    int
	rightSibling int
	depth        int
}

func CalcDepth(tree []Node, nodeValue int, depth int) {
	tree[nodeValue].depth = depth

	if tree[nodeValue].leftChild != -1 {
		CalcDepth(tree, tree[nodeValue].leftChild, depth+1)
	}

	if tree[nodeValue].rightSibling != -1 {
		CalcDepth(tree, tree[nodeValue].rightSibling, depth)
	}
}

func printTree(tree []Node) {
	for i := 0; i < len(tree); i++ {
		nodeType := ""
		if tree[i].parent == -1 {
			nodeType = "root"
		} else if tree[i].leftChild == -1 {
			nodeType = "leaf"
		} else {
			nodeType = "internal node"
		}

		children := []int{}
		if tree[i].leftChild != -1 {
			childNode := tree[i].leftChild
			for {
				if childNode == -1 {
					break
				}

				children = append(children, childNode)
				childNode = tree[childNode].rightSibling
			}
		}

		fmt.Printf("node %v: parent = %v, depth = %v, %v, ", i, tree[i].parent, tree[i].depth, nodeType)
		childrenFormat(children)
		fmt.Print("\n")
	}
}

func childrenFormat(children []int) {
	fmt.Print("[")
	for i := 0; i < len(children); i++ {
		if i == len(children)-1 {
			fmt.Print(children[i])
		} else {
			fmt.Printf("%v, ", children[i])
		}
	}
	fmt.Print("]")
}

func AppendNode(tree []Node, nodeValue int, children []int) []Node {
	for j := 0; j < len(children); j++ {
		if j == 0 {
			tree[nodeValue].leftChild = children[0]
		} else {
			tree[children[j-1]].rightSibling = children[j]
		}
		tree[children[j]].parent = nodeValue
	}

	return tree
}

func FindRootNodeIndex(tree []Node) int {
	rootNodeIndex := 0
	for i := 0; i < len(tree); i++ {
		if tree[i].parent == -1 {
			rootNodeIndex = i
			break
		}
	}

	return rootNodeIndex
}

func main() {
	sc := NewScanner()

	N := sc.nextInt()
	tree := make([]Node, N)

	for i := 0; i < N; i++ {
		tree[i] = Node{parent: -1, leftChild: -1, rightSibling: -1}
	}

	for i := 0; i < N; i++ {
		nodeValue := sc.nextInt()
		childrenCount := sc.nextInt()
		children := sc.nextIntSlice(childrenCount)

		tree = AppendNode(tree, nodeValue, children)
	}

	rootNodeIndex := FindRootNodeIndex(tree)

	CalcDepth(tree, rootNodeIndex, 0)

	printTree(tree)
}
