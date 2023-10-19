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
	sort.Slice(arr, func(i, j int) bool {return arr[i]<arr[j]})
	Process(arr, "", M, writer)
}

func Process(arr []int, s string, cnt int, writer *bufio.Writer){
	if cnt == 0 {
		fmt.Fprintln(writer, s)
		return
	}
	isPrint := make(map[int]bool)
	for i:=0 ; i<len(arr) ; i++ {
		if isPrint[arr[i]] == false {
			isPrint[arr[i]] = true
			nextString := fmt.Sprintf("%s%d ", s, arr[i])
			nextArr := []int{}
			nextArr = append(nextArr, arr[:i]...)
			nextArr = append(nextArr, arr[i+1:]...)
			Process(nextArr, nextString, cnt-1, writer)
		}
	}
}