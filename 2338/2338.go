package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var A, B big.Int
	fmt.Fscanln(reader, &A)
	fmt.Fscanln(reader, &B)
	rst := new(big.Int)
	// fmt.Fprintln(writer,rst.Add(&A,&B))
	// fmt.Fprintln(writer,rst.Sub(&A,&B))
	// fmt.Fprintln(writer,rst.Mul(&A,&B))
	fmt.Fprintf(writer, "%v\n%v\n%v",rst.Add(&A,&B).String(),rst.Sub(&A,&B).String(),rst.Mul(&A,&B).String())
}