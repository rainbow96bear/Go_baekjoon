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
	var score int
	fmt.Fscan(reader, &score)

	switch {
	case 100 >= score && score >= 90 :
		fmt.Fprintln(writer,"A")
	case 89 >= score && score >= 80 :
		fmt.Fprintln(writer, "B")
	case 79 >= score && score >= 70:
		fmt.Fprintln(writer, "C")
	case 69 >= score && score >= 60 :
		fmt.Fprintln(writer, "D")
	default :
		fmt.Fprintln(writer, "F")
	}
}