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

	var Num1, Num2 int

	fmt.Fscan(reader, &Num1, &Num2)
	Num1 = (Num1/100) + (Num1/10 %10 *10)+ (Num1 %10 * 100)
	Num2 = (Num2/100) + (Num2/10 %10 *10)+ (Num2 %10 * 100)
	if Num1 < Num2 {
		fmt.Fprint(writer, Num2)
	}else{
		fmt.Fprint(writer, Num1)
	}
}