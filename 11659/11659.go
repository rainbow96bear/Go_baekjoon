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

	numSumList := []int{0}
	for i:=0 ; i<N ; i++ {
		var input int
		fmt.Fscan(reader, &input)

		numSumList = append(numSumList, numSumList[i]+input)
	}

	for i:=0 ; i<M ; i++ {
		var start, end int
		fmt.Fscan(reader, &start, &end)
		fmt.Fprintln(writer, numSumList[end]-numSumList[start-1])
	}
}