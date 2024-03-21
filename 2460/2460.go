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

	max := 0
	people := 0
	for i := 0 ; i< 10 ; i++ {
		var in, out int
		fmt.Fscanf(reader, "%d %d\n", &out, &in)
		people += in
		people -= out
		if max < people {
			max = people
		}
	}
	fmt.Fprintln(writer, max)
}