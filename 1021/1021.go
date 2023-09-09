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
	
	var N, M int
	fmt.Fscan(reader, &N, &M)

	circle := make([]int,N)
	for i:=0 ; i<N ; i++ {
		circle[i]=i+1
	}
	answer:=0
	for i:=0 ; i<M ; i++{
		var input int
		fmt.Fscan(reader, &input)
		cnt:=0
		for j:=0 ; j<len(circle) ; j++{
			if input == circle[j]{
				break
			}
			cnt++
		}

		if cnt < len(circle)-cnt{
			answer += cnt
		}else {
			answer += len(circle)-cnt
		}	
		circle = append(circle[cnt+1:], circle[:cnt]...)
	}
	fmt.Fprint(writer, answer)
}