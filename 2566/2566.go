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

	var input int
	max := 0
	var N, M int
	for i:=0 ; i<9*9 ; i++{
		fmt.Fscan(reader,&input)
		if max <= input{
			max = input
			N = (i/9)+1
			M = (i%9) +1
		}
	}
	fmt.Fprintf(writer,"%d\n%d %d",max,N,M)
}