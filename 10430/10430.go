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

	var A, B, C int
	fmt.Fscan(reader,&A, &B, &C)

	fmt.Fprintln(writer, (A+B)%C)
	fmt.Fprintln(writer, ((A%C)+(B%C))%C)
	fmt.Fprintln(writer, (A*B)%C)
	fmt.Fprintln(writer, ((A%C)*(B%C))%C)
}