package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var testCase int
	fmt.Fscan(reader,&testCase)
	for i:=0 ; i<testCase ; i++ {
		var R int
		var P string
		fmt.Fscan(reader, &R, &P)
		for _, v := range P {
			fmt.Fprint(writer, strings.Repeat(string(v),R))
		}
		fmt.Fprintln(writer)
	}
}