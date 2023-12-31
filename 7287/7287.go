package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	fmt.Fprintln(writer, "362")
	fmt.Fprintln(writer, "rainbow96bear")
}