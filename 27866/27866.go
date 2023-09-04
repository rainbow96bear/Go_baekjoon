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

	var text string
	var num int
	fmt.Fscan(reader, &text)
	fmt.Fscan(reader, &num)

	fmt.Fprint(writer, string(text[num-1]))
}