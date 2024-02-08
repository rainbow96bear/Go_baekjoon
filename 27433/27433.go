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

	var N int64
	fmt.Fscanf(reader, "%d\n", &N)
	var answer int64
	answer = 1
	for i:=int64(1) ; i<=N ; i++ {
		answer *= i
	}

	fmt.Fprintln(writer, answer)
}