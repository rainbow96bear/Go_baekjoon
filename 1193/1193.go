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

	var X int
	var toggle bool = false
	fmt.Fscan(reader, &X)
	check := 0
	point := 0
	for i:=1 ; check<X ; i++{
		check += i
		point = i
		toggle = !toggle
	}
	if toggle {
		fmt.Fprintf(writer,"%d/%d",check-X+1,point-check+X)
		}else {
		fmt.Fprintf(writer,"%d/%d",point-check+X,check-X+1)
	}
}
