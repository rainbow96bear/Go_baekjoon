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

	sort.Slice(arr, func(i, j int) bool {return arr[i] < arr[j]} )

	PrintProcess(arr, "", M, writer)

}

func PrintProcess(arr []int, s string, M int, writer *bufio.Writer){
	if M > 1 {
		for i:=0 ; i<len(arr) ; i++ {
			nextS := s + fmt.Sprintf("%d ",arr[i])
			nextArr := []int{}
			nextArr = append(nextArr, arr[:i]...)
			nextArr = append(nextArr, arr[i+1:]...)
			PrintProcess(nextArr, nextS, M-1, writer)
		}
	}else {
		for i:=0 ; i<len(arr) ; i++ {
			fmt.Fprintf(writer, "%s%d\n", s, arr[i])
		}
		return
	}
}