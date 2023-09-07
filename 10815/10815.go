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

	cards := make(map[int]bool,N)
	for i:=0 ; i<N ; i++{
		var input int
		fmt.Fscan(reader, &input)
		cards[input]=true
	}

	var M int
	fmt.Fscan(reader, &M)

	for i:=0 ; i<M ; i++{
		var input int
		fmt.Fscan(reader, &input)
		if cards[input] {
			fmt.Fprint(writer, "1 ")
		}else {
			fmt.Fprint(writer, "0 ")
		}
	}
}