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

	var name string
	fmt.Fscan(reader, &name)
	name = name+"??!"
	fmt.Fprintln(writer,name)
}