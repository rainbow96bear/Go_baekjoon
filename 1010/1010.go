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

	var testCase int
	fmt.Fscan(reader, &testCase)

	for i:=0 ; i<testCase ; i++{
		var N, M int
		fmt.Fscan(reader, &N ,&M)
		answer := 1
		for j:=1 ; j<=N ; j++{
			answer *= (M-j+1)
			answer /= j
		}
		fmt.Fprintln(writer, answer)
	}
}