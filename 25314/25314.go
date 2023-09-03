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

	var byte int
	fmt.Fscan(reader, &byte)

	for i:=0 ; i<(byte/4) ; i++{
		fmt.Fprint(writer,"long ")
	}
	fmt.Fprint(writer, "int")
}