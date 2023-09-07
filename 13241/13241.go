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

	var a, b int
	fmt.Fscan(reader, &a, &b)
	big, small := a, b
	if a < b {
		big, small = b, a
	}
	for i:=1 ; true ; i++{
		if (small*i) % big == 0{
			fmt.Fprint(writer, small*i)
			return
		}
	}
}