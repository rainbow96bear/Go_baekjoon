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
	total := 1
	for n:=0 ; n<N ; n++ {
		var input int
		fmt.Fscanf(reader, "%d\n", &input)
		total += (input-1)
	}
	fmt.Fprintln(writer, total)
}

