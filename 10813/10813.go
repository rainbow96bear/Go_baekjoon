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
	arr := make([]int,N)
	for i:=0 ; i<N ; i++{
		arr[i]=i+1
	}
	for i:=0 ; i<M ; i++{
		var a, b int
		fmt.Fscan(reader, &a, &b)
		arr[a-1], arr[b-1] = arr[b-1], arr[a-1]
	}
	for _, v := range arr{
		fmt.Fprintf(writer,"%d ",v)
	}
}

