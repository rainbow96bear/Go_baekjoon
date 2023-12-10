package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	num   int
	value int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)
	list := make([][]int, M)
	for i := 0; i < M; i++ {
		list[i] = make([]int, 2)
		fmt.Fscanf(reader, "%d %d\n", &list[i][0], &list[i][1])
	}

	result := Process(N, M, list)

	for i := 0; i < N; i++ {
		fmt.Fprintf(writer, "%d ", result[i])
	}
}

func Process(N, M int, list [][]int) []int {
	childList := make(map[int][]int)
	valueList := make(map[int]int)
	for i := 0; i < M; i++ {
		childList[list[i][0]] = append(childList[list[i][0]], list[i][1])
		valueList[list[i][1]]++
	}
	stack := []int{}
	result := []int{}
	for i := 1; i <= N; i++ {
		if valueList[i] == 0 {
			stack = append(stack, i)
			result = append(result, i)
		}
	}

	for len(stack) > 0 {
		now := stack[0]
		stack = stack[1:]
		for i := 0; i < len(childList[now]); i++ {
			valueList[childList[now][i]]--
			if valueList[childList[now][i]] == 0 {
				stack = append(stack, childList[now][i])
				result = append(result, childList[now][i])
			}
		}
	}

	return result
}
