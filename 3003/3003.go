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

	answer := []int{1, 1, 2, 2, 2, 8}

	var input int
	for _ , v := range answer {
		fmt.Fscan(reader, &input)
		fmt.Fprintf(writer,"%d ",v - input)
	}
}