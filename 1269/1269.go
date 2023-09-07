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
	check := make(map[int]bool)
	for i:=0 ; i<N+M ; i++ {
		var input int
		fmt.Fscan(reader, &input)
		check[input] = !check[input]
	}
	answer:=0
	for _, v := range check {
		if v {
			answer ++
		}
	}
	fmt.Fprint(writer, answer)
}