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

	var A, B int
	fmt.Fscanf(reader, "%d %d\n", &A, &B)

	min, max := A+B, A+B

	ten := 1
	for A > 0 {
		if A%10 == 5 {
			max += ten
		}
		if A%10 == 6 {
			min -= ten
		}
		A /= 10
		ten *= 10
	}

	ten = 1
	for B > 0 {
		if B%10 == 5 {
			max += ten
		}
		if B%10 == 6 {
			min -= ten
		}
		B /= 10
		ten *= 10
	}

	fmt.Fprintln(writer, min, max)
}