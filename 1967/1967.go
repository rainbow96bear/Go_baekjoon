package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
type node struct {
	child int
	length int
}
var diameterList []int = []int{}
func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	N, _ := strconv.Atoi(Input(scanner))

	tree := make(map[int][]node)
	for i:=1 ; i<N ; i++ {
		info := strings.Split(Input(scanner), " ")
		parent, _ := strconv.Atoi(info[0])
		child, _ := strconv.Atoi(info[1])
		length, _ := strconv.Atoi(info[2])
		tree[parent] = append(tree[parent], node{child, length})
	}
	if N == 1 {
		fmt.Fprintln(writer, 0)
		return
	}
	DFS(1, tree)

	sort.Slice(diameterList, func(i, j int) bool {return diameterList[i] > diameterList[j]})

	fmt.Fprintln(writer, diameterList[0])
}

func DFS(parent int, tree map[int][]node) (int) {
	if len(tree[parent]) == 0 {
		return 0
	}
	maxList := make([]int, 2)
	for _, v := range tree[parent] {
		length := DFS(v.child, tree)
		result := v.length + length
		if maxList[0] < result {
			maxList[0], maxList[1] = result, maxList[0]
		}else if maxList[1] < result {
			maxList[1] = result
		}
	}

	diameterList = append(diameterList, maxList[0] + maxList[1])
	return maxList[0]
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}