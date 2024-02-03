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

	var input int
	fmt.Fscanf(reader, "%d\n", &input)

	PrintNumber(writer, input)
}

func PrintNumber(writer *bufio.Writer, number int) {
	for n:=number ; n>0 ; n-- {
		fmt.Fprintln(writer, n)
	}
}