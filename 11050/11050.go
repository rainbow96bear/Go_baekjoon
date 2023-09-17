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

	var N, K int
	fmt.Fscan(reader, &N, &K)

	answer := 1
	for i:=0 ; i<K ; i++ {
		answer *= N-i
	}

	for i:=1 ; i<=K ; i++ {
		answer /= i
	}
	fmt.Fprintln(writer, answer)
}