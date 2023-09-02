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

	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)

	fmt.Fprintln(writer,a+b+c)
}