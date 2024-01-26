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

	for {
		var N int
		fmt.Fscanf(reader, "%d", &N)
		if N == 0 {
			return
		}
		brunch := 1
		for n:=0 ; n<N ; n++ {
			var sf, cut int
			fmt.Fscanf(reader, "%d %d", &sf, &cut)
			brunch = brunch*sf - cut
		}
		fmt.Fscanf(reader, "\n")
		fmt.Fprintln(writer, brunch)
	}
}