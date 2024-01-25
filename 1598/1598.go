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

	var a, b int
	fmt.Fscanf(reader, "%d %d\n", &a, &b)
	width := (a-1)/4 - (b-1)/4
	height := (a-1)%4 - (b-1)%4
	fmt.Fprintln(writer, abs(width)+abs(height))
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}