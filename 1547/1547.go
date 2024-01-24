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

	var M int
	fmt.Fscanf(reader, "%d\n", &M)
	answer := 1
	for m:=0 ; m<M ; m++ {
		var a, b int
		fmt.Fscanf(reader,"%d %d\n", &a, &b)

		if a == answer {
			answer = b
		}else if b == answer {
			answer = a
		}
	}
	fmt.Fprintln(writer, answer)
}