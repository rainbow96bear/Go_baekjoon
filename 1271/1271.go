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

	var N, M big.Int

	fmt.Fscan(reader, &N,&M)

	quotient := new(big.Int)
    remainder := new(big.Int)

    quotient.DivMod(&N, &M, remainder)

	fmt.Fprintf(writer, "%v\n%v",quotient,remainder)
}