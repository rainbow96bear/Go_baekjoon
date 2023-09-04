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

	var inputText string
	fmt.Fscan(reader, &inputText)

	fmt.Fprint(writer, inputText[0])
}