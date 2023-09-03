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

	var Num int
	fmt.Fscan(reader, &Num)
	var answer int
	for i:=1 ; i<=Num ; i++{
		answer += i
	}
	fmt.Fprintln(writer, answer)
}