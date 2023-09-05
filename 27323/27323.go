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

	var X, Y int
	fmt.Fscanln(reader, &X)
	fmt.Fscanln(reader, &Y)

	fmt.Fprint(writer, X*Y)
}