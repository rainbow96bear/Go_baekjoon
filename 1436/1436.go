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
	cnt:=0
	for i:=1;true;i++{
		check := i
		for check>0 {
			if check%10 == 6 && check%1000 == 666{
				cnt++
				break
			}
			check/=10
		}
		if cnt == N{
			fmt.Fprint(writer,i)
			break
		}
	}
}