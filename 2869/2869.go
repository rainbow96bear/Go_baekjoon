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

	var A, B, V int
	fmt.Fscan(reader, &A, &B ,&V)
	if (V-A)%(A-B) > 0{
		fmt.Fprint(writer,(V-A)/(A-B) + 2)
	} else{
		fmt.Fprint(writer,(V-A)/(A-B) + 1)
	}
}