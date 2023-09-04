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
		var inputText string
		fmt.Fscan(reader, &inputText)
		fmt.Fprintf(writer,"%s%s\n", string(inputText[0]),string(inputText[len(inputText)-1]))
	}
}