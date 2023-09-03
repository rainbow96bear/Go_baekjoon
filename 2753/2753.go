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

	var year int
	fmt.Fscan(reader, &year)

	if year % 4 == 0 {
		if year % 400 == 0 {
			fmt.Fprintln(writer, 1)
		}else if year % 100 == 0 {
			fmt.Fprintln(writer, 0)
		}else {
			fmt.Fprintln(writer, 1)
		}
	}else {
		fmt.Fprintln(writer, 0)
	}
}