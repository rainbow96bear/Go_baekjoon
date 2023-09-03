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

	var N, X int
	fmt.Fscan(reader, &N, &X)

	for i:=0 ; i<N ; i++{
		var Num int
		fmt.Fscan(reader, &Num)
		if Num < X{
			fmt.Fprintf(writer,"%d ", Num)
		}
	}
}