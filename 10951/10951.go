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
		_, err := fmt.Fscanln(reader, &A, &B)
		if err != nil {
			return
		}
		fmt.Fprintln(writer, A+B)
	}
}