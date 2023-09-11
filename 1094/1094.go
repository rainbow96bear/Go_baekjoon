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

	var X int
	fmt.Fscan(reader, &X)
	cnt := 0
	for X != 0 {
		if X%2 == 1{
			cnt++
		}
		X/=2
	}
	fmt.Fprintln(writer, cnt)
}