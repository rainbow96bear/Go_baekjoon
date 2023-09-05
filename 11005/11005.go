package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, B int
	fmt.Fscan(reader, &N, &B)
	if N==0{
		fmt.Fprint(writer, 0)
		return
	}
	var rst []string
	for N>0{
		if N%B <10 {
			rst = append(rst,strconv.Itoa(N%B))
		}else {
			rst = append(rst,string('A'+N%B-10))
		}
		N = N/B
	}
	for i:=len(rst)-1; i>=0 ; i--{
		fmt.Fprintf(writer,"%s",rst[i])
	}
}