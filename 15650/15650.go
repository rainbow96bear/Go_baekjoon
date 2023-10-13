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

	arr := make([]int, M+1)
	for i:=0 ; i<=M ; i++ {
		arr[i] = i
	}
	if M == 1 {
		for i:=1 ; i<=N ; i++ {
			fmt.Fprintln(writer, i)
		}
		return
	}
	for arr[M] <= N && arr[0] == 0{
		for i:=1 ; i<=M ; i++ {
			fmt.Fprintf(writer, "%d ",arr[i])
		}
		fmt.Fprintln(writer)
		arr[M]++
		
		for i:=M ; i>0 ; i-- {
			if arr[i] > N-M+i {
				arr[i-1] += 1
				for j:=0 ; i+j<=M ; j++ {
					arr[i+j] = arr[i+j-1]+1
				}
			}
		}
	}
}