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
	
	var N int
	fmt.Fscan(reader, &N)
	A := make([]int, N)
	sortA := make([]int, N)
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader, &A[i])
		sortA[i]=A[i]
	}
	sort.Ints(sortA)
	check := make([]int,N)
	for i:=0 ; i<N ; i++{
		index := sort.SearchInts(sortA, A[i])
		fmt.Fprintf(writer, "%d ",index+check[index])
		check[index]++
	}
}