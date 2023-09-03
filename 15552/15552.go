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

	var Num int
	fmt.Fscan(reader, &Num)

	for i:=0 ; i<Num ; i++{
		var A, B int
		fmt.Fscan(reader, &A, &B)
		fmt.Fprintln(writer, A+B)
	}
}