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

	var max, num int
	for i:=1 ; i<=9 ; i++{
		var input int
		fmt.Fscan(reader, &input)
		if max < input {
			max = input
			num = i
		}
	}
	fmt.Fprintf(writer,"%d\n%d", max,num)
}