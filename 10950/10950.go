package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var testNum int
	fmt.Fscan(reader, &testNum)
	
	for i:=0 ; i<testNum ; i++{
		var A, B int
		fmt.Fscan(reader, &A, &B)
		fmt.Fprintln(writer,A+B)
	}
}