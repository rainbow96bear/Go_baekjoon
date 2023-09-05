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

	for {
		var A, B int
		fmt.Fscan(reader, &A, &B)
		if A == B {
			break
		} else if A < B && B%A==0{
			fmt.Fprint(writer,"factor\n")
		} else if A > B && A%B==0 {
			fmt.Fprint(writer, "multiple\n")
		} else {
			fmt.Fprint(writer,"neither\n")
		}
	}
}