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
	pointNum := 2
	for i:=0 ; i<N ; i++{
		pointNum = pointNum*2 - 1
	}
	fmt.Fprint(writer, pointNum*pointNum)
}