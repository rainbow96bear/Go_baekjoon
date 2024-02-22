package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
var (
	list = []int{}
	numList = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	isAdd = make(map[int]bool)
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	
	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)
	dfs(0, 0)
	sort.Slice(list, func(i, j int)bool{return list[i]<list[j]})
	if N > len(list) {
		fmt.Fprintln(writer, "-1")
	}else {
		fmt.Fprintln(writer, list[N-1])
	}
}

func dfs(num, index int) {
	if isAdd[num] == false {
		list = append(list, num)
		isAdd[num] = true
	}
	if index >= 10 {
		return
	}


	dfs(num*10+numList[index], index+1)
	dfs(num, index+1)
}