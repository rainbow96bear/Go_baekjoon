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

	for n:=N ; n>0 ; n-- {
		str := strings.Repeat("*", n)
		fmt.Fprintln(writer, str)
	}
}