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

	var N int
	fmt.Fscan(reader, &N)

	tp := make([]int,N)
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader, &tp[i])
	}

	gap := []int{}
	for i:=1 ; i<N ; i++{
		gap = append(gap, tp[i]-tp[i-1])
	}
	gdcV := gap[0]
	for i:=1 ; i<len(gap) ; i++{
		gdcV = gcd(gdcV,gap[i])
	}
	fmt.Fprint(writer, (tp[len(tp)-1]-tp[0])/gdcV - len(tp) + 1)
}

func gcd(a, b int)int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}