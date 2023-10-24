package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	tree := make(map[string][]string)
	isNotRoot := make(map[string]bool)
	for i:=0 ; i<N ; i++ {
		var parent, left, right string
		fmt.Fscan(reader, &parent, &left, &right)
		tree[parent] = []string{left, right}
		isNotRoot[left] = true
		isNotRoot[right] = true
	}
	root := ""
	for i:='A' ; i<rune(int('A')+N) ; i++ {
		if isNotRoot[string(i)] == false {
			root = string(i)
			break
		}
	}
	front := frontProcess(tree, root)
	middle := middleProcess(tree, root)
	back := backProcess(tree, root)

	fmt.Fprintln(writer, front)
	fmt.Fprintln(writer, middle)
	fmt.Fprintln(writer, back)
}

func frontProcess(tree map[string][]string, root string) string {
	leftResult := ""
	rightResult := ""
	if tree[root][0] != "." {
		leftResult = frontProcess(tree, tree[root][0])
	}
	if tree[root][1] != "." {
		rightResult = frontProcess(tree, tree[root][1])
	}
	result := root + leftResult + rightResult
	return result
}

func middleProcess(tree map[string][]string, root string) string {
	leftResult := ""
	rightResult := ""
	if tree[root][0] != "." {
		leftResult = middleProcess(tree, tree[root][0])
	}
	if tree[root][1] != "." {
		rightResult = middleProcess(tree, tree[root][1])
	}
	result := leftResult + root + rightResult
	return result
}

func backProcess(tree map[string][]string, root string) string {
	leftResult := ""
	rightResult := ""
	if tree[root][0] != "." {
		leftResult = backProcess(tree, tree[root][0])
	}
	if tree[root][1] != "." {
		rightResult = backProcess(tree, tree[root][1])
	}
	result := leftResult + rightResult + root
	return result
}
