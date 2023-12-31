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
	fmt.Fscanf(reader, "%d %d\n", &A, &B)

	fmt.Fprintln(writer, A*A-B*B)
}