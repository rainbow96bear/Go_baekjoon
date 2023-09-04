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

	var N int
	var Nline string
	var zero string = "0"
	fmt.Fscan(reader, &N)
	fmt.Fscan(reader, &Nline)
	var answer int
	for i:=0 ; i<N ; i++{
		answer += int(Nline[i]-zero[0])
	}
	fmt.Fprint(writer, answer)
}