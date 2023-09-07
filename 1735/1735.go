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

	var a1, a2 int
	fmt.Fscan(reader, &a1, &a2)

	var b1, b2 int
	fmt.Fscan(reader, &b1, &b2)
	for i:=a2*b2 ; i>0 ; i-- {
		if (a1*b2+b1*a2)%i == 0 && (a2*b2)%i ==0 {
			fmt.Fprintf(writer,"%d %d",(a1*b2+b1*a2)/i,(a2*b2)/i)
			return
		}
	}
}