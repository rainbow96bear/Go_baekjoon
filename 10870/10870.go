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
	fmt.Fscanf(reader, "%d\n", &N)

	if N == 0 {
		fmt.Fprintln(writer, "0")
	}else if N == 1 || N == 2{
		fmt.Fprintln(writer, "1")
	}else {
		beforeNum := int64(1)
		nowNum := int64(1)
		for i:=1 ; i<=N-2 ; i++ {
			nowNum, beforeNum = nowNum+beforeNum, nowNum
		}
		fmt.Fprintln(writer, nowNum)
	}
}