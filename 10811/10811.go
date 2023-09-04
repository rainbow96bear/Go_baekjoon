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

	var N, M int
	fmt.Fscan(reader, &N, &M)
	bucket := make([]int,N+1)

	for i:=1 ; i<=N ; i++{
		bucket[i]=i
	}
	for i:=0 ; i<M ; i++{
		var start, end int
		fmt.Fscan(reader, &start, &end)
		for j:=0 ; j<=(end-start)/2 ; j++{
			bucket[start+j], bucket[end-j] =bucket[end-j], bucket[start+j]
		}
	}
	bucket=bucket[1:]
	for _, v := range bucket{
		fmt.Fprintf(writer, "%d ",v)
	}
}