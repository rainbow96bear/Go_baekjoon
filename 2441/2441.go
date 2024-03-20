package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	for i:=0 ; i<N ; i++{
		fmt.Fprintf(writer, "%s%s\n", strings.Repeat(" ", i), strings.Repeat("*", N-i))
	}
}