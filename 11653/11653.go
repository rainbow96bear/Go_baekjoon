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
	fmt.Fscan(reader, &N)

	Num := N
	for i:=2 ; i<=N ; i++{
		for Num%i == 0 {
			fmt.Fprintln(writer, i)
			Num/=i
		}
	}
}