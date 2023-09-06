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

	var a, b, c, d, e, f int
	fmt.Fscan(reader, &a, &b, &c, &d, &e, &f)
	fmt.Fprint(writer,(c*e-f*b)/(a*e - d*b),(c*d-f*a)/(b*d-e*a))
}