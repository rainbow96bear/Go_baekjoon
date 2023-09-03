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

	var x, y int

	fmt.Fscan(reader, &x)
	fmt.Fscan(reader, &y)

	if x >0 {
		if y > 0 {
			fmt.Fprintln(writer, 1)
		}else{
			fmt.Fprintln(writer, 4)
		}
	} else {
		if y > 0 {
			fmt.Fprintln(writer, 2)
		}else{
			fmt.Fprintln(writer, 3)
		}
	}

}