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

	total, max := 0, 0
	for i:=0 ; i<4; i++ {
		var out, in int
		fmt.Fscanf(reader, "%d %d\n", &out, &in)

		total -= out
		if total > max {
			max = total
		}

		total += in
		if total > max {
			max = total
		}
	}
	fmt.Fprintln(writer, max)
}