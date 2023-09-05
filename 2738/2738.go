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

	rst := make([]int,N*M)
	for i:=0 ; i<2*N*M ; i++ {
			var input int
			fmt.Fscan(reader, &input)
			rst[i%(N*M)] += input
	}
	for i:=0 ; i<N*M ; i++ {
		fmt.Fprintf(writer,"%d ", rst[i%(N*M)])
		if i%M == M-1{
			fmt.Fprint(writer,"\n")
		}
	}
}