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

	var testCase int
	fmt.Fscan(reader, &testCase)

	for i:=0 ; i<testCase ; i++ {
		var C int
		fmt.Fscan(reader, &C)

		fmt.Fprintf(writer,"%d ",C/25)
		C%=25
		fmt.Fprintf(writer,"%d ",C/10)
		C%=10
		fmt.Fprintf(writer,"%d ",C/5)
		C%=5
		fmt.Fprintf(writer,"%d\n",C)
	}
}