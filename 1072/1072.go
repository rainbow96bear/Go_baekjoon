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

	var X, Y int
	fmt.Fscan(reader, &X, &Y)

	Z := Y*100/X

	if 99 <= Z {
		fmt.Fprint(writer, -1)
		return
	}
	if (Z*X+X-100*Y)%(99 - Z) == 0 {
		fmt.Fprint(writer, (Z*X+X-100*Y)/(99 - Z))
	}else {
		fmt.Fprint(writer, (Z*X+X-100*Y)/(99 - Z)+1)
	}

}