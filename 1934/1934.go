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
		var A, B int
		fmt.Fscan(reader, &A, &B)
		bigNum := A
		smallNum := B
		if A < B {
			bigNum, smallNum = B, A
		}
		for i:=1 ; true ; i++{
			if (smallNum*i)%bigNum == 0 {
				fmt.Fprintln(writer, smallNum*i)
				break
			}
		}
	}
}