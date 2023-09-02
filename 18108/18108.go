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

	var bulgiYear int
	fmt.Fscan(reader, &bulgiYear)
	seogiYear := bulgiYear - (2541-1998)
	fmt.Fprint(writer, seogiYear)
}