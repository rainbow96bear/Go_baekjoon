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
	for i:=1 ; i<=N ; i++{
		sum := i
		check := i
		for check!=0{
			sum += check%10
			check/=10
		}
		if sum == N {
			fmt.Fprint(writer,i)
			return
		} 
	}
	fmt.Fprint(writer,0)
}