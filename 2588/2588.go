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
	fmt.Fscan(reader, &Num1)
	fmt.Fscan(reader, &Num2)
	answer := Num1 * Num2
	fmt.Fprintln(writer,Num1*(Num2%10))
	Num2=Num2/10
	fmt.Fprintln(writer,Num1*(Num2%10))
	Num2=Num2/10
	fmt.Fprintln(writer,Num1*(Num2%10))

	fmt.Fprintln(writer,answer)
}