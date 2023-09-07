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

	var N, M int
	fmt.Fscan(reader, &N, &M)

	groupS := make(map[string]bool,N)
	for i:=0 ; i<N ; i++{
		var input string
		fmt.Fscan(reader, &input)
		groupS[input]=true
	}
	cnt := 0
	for i:=0 ; i<M ; i++{
		var input string
		fmt.Fscan(reader, &input)
		if groupS[input] {
			cnt++
		}
	}
	fmt.Fprint(writer, cnt)
}