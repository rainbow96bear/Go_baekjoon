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

	var ang [3]int
	var sum int = 0
	for i:=0 ; i<3 ; i++{
		fmt.Fscan(reader, &ang[i])
		sum += ang[i]
	}
	if sum != 180{
		fmt.Fprint(writer, "Error")
		return
	}
	if ang[0] == 60 && ang[1] == 60 && ang[2] == 60{
		fmt.Fprint(writer, "Equilateral")
		return
	}
	if ang[0] != ang[1] && ang[1] != ang[2] && ang[2] != ang[0]{
		fmt.Fprint(writer, "Scalene")
	}else {
		fmt.Fprint(writer, "Isosceles")
	}
}