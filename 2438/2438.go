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

	var Num int
	fmt.Fscan(reader, &Num)

	for i:=1 ; i<=Num ; i++{
		fmt.Fprintln(writer, strings.Repeat("*",i))
	}
}