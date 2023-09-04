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

	var input int
	fmt.Fscan(reader,&input)

	for i:=1 ; i<=input ; i++{
		fmt.Fprint(writer,strings.Repeat(" ",input-i))
		fmt.Fprintln(writer,strings.Repeat("*",2*i-1))
	}
	for i:=input-1 ; i>=1 ; i--{
		fmt.Fprint(writer,strings.Repeat(" ",input-i))
		fmt.Fprintln(writer,strings.Repeat("*",2*i-1))
	}
}