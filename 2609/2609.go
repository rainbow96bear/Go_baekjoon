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

	var A, B int
	fmt.Fscan(reader, &A, &B)
	fmt.Fprintln(writer, min(A, B))
	fmt.Fprintln(writer, max(A, B))
}

func max(a, b int) int {
	return a*b / min(a,b)
}

func min(a, b int) int {
	for b != 0{
		a, b = b, a%b
	}
	return a
}