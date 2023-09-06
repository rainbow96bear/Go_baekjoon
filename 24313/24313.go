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

	var a1, a0 int
	fmt.Fscan(reader, &a1, &a0)

	var c int
	fmt.Fscan(reader, &c)

	var n0 int
	fmt.Fscan(reader, &n0)
	if c-a1 >=0{
		if a0 <= (c-a1)*n0 {
			fmt.Fprint(writer, 1)
		}else {
			fmt.Fprint(writer, 0)
		}
	}else {
		fmt.Fprint(writer, 0)
	} 
}