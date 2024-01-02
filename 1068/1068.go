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
	var N, removeNum int
	fmt.Fscanln(reader, &N)

	list := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &list[i])
	}
	fmt.Fscanln(reader)
	
	fmt.Fscanln(reader, &removeNum)
    result := FindLeafNodeNum(N, list, removeNum)
    fmt.Println(result)
}

func FindLeafNodeNum(N int, list []int, removeNum int) int {
	childList := make([][]int, N)

	rootNode := 0
	for i:=0 ; i<N ; i++ {
		if list[i] == -1 {
			rootNode = i
		}else if i != removeNum{
			childList[list[i]] = append(childList[list[i]], i)
		}
	}

	if rootNode == removeNum {
		return 0
	}
	count := 0
	queue := []int{rootNode}
	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]
		if len(childList[now]) == 0 {
			count ++
		}
		for i:=0 ; i<len(childList[now]) ; i++ {
			if childList[now][i] != removeNum {
				queue = append(queue, childList[now][i])
			}
		}
	}
	return count
}