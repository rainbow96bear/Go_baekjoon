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

	var (
		A, B, C int
	)
	fmt.Fscanf(reader, "%d %d %d\n", &A, &B, &C)

	if C > B {
		fmt.Fprintln(writer, A/(C-B)+1)
	}else {
		fmt.Fprintln(writer, "-1")
	}
}