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
	fmt.Fscan(reader, &text)
	fmt.Fprint(writer,len(text))
}