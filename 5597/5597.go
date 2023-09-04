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

	arr := make(map[int]int,30)
	for i:=0 ; i<28 ; i++{
		var input int
		fmt.Fscan(reader, &input)
		arr[input]=1
	}
	for i:=1 ; i<=30 ; i++{
		if arr[i] == 0 {
			fmt.Fprintln(writer,i)
		}
	}
}