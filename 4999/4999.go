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

	var first, second string

	fmt.Fscanf(reader, "%s\n", &first)
	fmt.Fscanf(reader, "%s\n", &second)

	if len(first) >= len(second) {
		fmt.Fprintln(writer, "go")
	}else {
		fmt.Fprintln(writer, "no")
	}

}