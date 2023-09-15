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

	var N, newScore, P int
	fmt.Fscan(reader, &N, &newScore, &P)

	ranking := make([]int, P)
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader, &ranking[i])
	}
	cnt := 0
	check := 0
	for i:=0 ; i<P ; i++ {
		if ranking[i] <= newScore{
			if check == 0 {
				check = i+1
			}
			if ranking[i] < newScore || newScore == 0{
				break
			}
		}
		cnt++
	}
	if cnt == P {
		fmt.Fprint(writer, -1)
	}else {
		fmt.Fprint(writer, check)
	}
}