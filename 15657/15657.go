package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, M int
	fmt.Fscan(reader, &N, &M)

	arr := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	sort.Slice(arr, func(i, j int) bool {return arr[i] < arr[j]})
	Process(arr, M, "", writer)
}

func Process(arr []int, cnt int, s string, writer *bufio.Writer){
	if cnt > 1 {
		for i:=0 ; i<len(arr) ; i++ {
			Process(arr[i:], cnt-1, fmt.Sprintf("%s%d ", s, arr[i]), writer)
		}
	}else {
		for i:=0 ; i<len(arr) ; i++ {
			fmt.Fprintf(writer, "%s%d\n", s, arr[i])
		}
	}
}