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

	cards := make(map[int]int,N)
	for i:=0 ; i<N ; i++{
		var input int
		fmt.Fscan(reader, &input)
		cards[input]++
	}

	var M int
	fmt.Fscan(reader, &M)

	for i:=0 ; i<M ; i++{
		var input int
		fmt.Fscan(reader, &input)
		fmt.Fprintf(writer, "%d ",cards[input])
	}
}