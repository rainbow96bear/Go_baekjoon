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

	var TotalPrice int
	fmt.Fscan(reader, &TotalPrice)

	var Num int
	fmt.Fscan(reader, &Num)

	var PriceSum int
	for i:=0 ; i<Num ; i++{
		var a, b int
		fmt.Fscan(reader, &a, &b)
		PriceSum += a*b
	}
	if TotalPrice == PriceSum {
		fmt.Fprint(writer, "Yes")
	}else {
		fmt.Fprint(writer, "No")
	}
}