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

	var N, K, I int
	fmt.Fscan(reader, &N, &K, &I)

	for i:=0 ; ; i++ {
		if K == I {
			fmt.Fprint(writer, i)
			break
		}
		if K%2 == 1 {
			K = K/2 + 1
		}else {
			K = K/2
		}
		if I%2 == 1 {
			I = I/2 + 1
		}else {
			I = I/2
		}
	}
}