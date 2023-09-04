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

	var testNum int
	fmt.Fscan(reader, &testNum)
	for i:=0 ; i<testNum ; i++{
		var answer int = 1
		var a, b int
		fmt.Fscan(reader, &a, &b)
		for j:=0 ; j<b ; j++{
			answer = answer*a%10
		}
		if answer == 0 {
			fmt.Fprintln(writer,10)
		}else {
			fmt.Fprintln(writer,answer)
		}
	}
}