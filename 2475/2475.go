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

	var answer int
	for i:=0 ; i<5 ; i++{
		var input int
		fmt.Fscan(reader, &input)
		answer += input*input
	}
	fmt.Fprint(writer,answer%10)
}