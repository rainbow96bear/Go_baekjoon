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

	for {
		var A, B int
		fmt.Fscan(reader, &A, &B)
		if A == 0 && B == 0 {
			return
		}
		fmt.Fprintln(writer,A+B)
	}
}