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
	for i:=0 ; i<M ; i++{
		var start, end, ballNum int
		fmt.Fscan(reader, &start, &end, &ballNum)
		for j:=start-1 ; j<=end-1 ; j++{
			arr[j]=ballNum
		}
	}
	for _, v := range arr{
		fmt.Fprintf(writer,"%d ",v)
	}
}