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
	fmt.Fscanf(reader, "%d\n", &N)

	answer := 0
	for i:=1 ; N != 0 ; i++ {
		if i <= N {
			N -= i
			answer++ 
		}else {
			i = 0
		}
	}
	fmt.Fprintln(writer, answer)
}