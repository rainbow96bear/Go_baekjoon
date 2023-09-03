package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var Num int
	fmt.Fscan(reader, &Num)

	for i:=1 ; i<10 ; i++ {
		fmt.Fprintf(writer, "%d * %d = %d\n", Num, i, Num*i )
	}
}