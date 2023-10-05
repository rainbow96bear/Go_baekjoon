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

	var T int
	fmt.Fscan(reader, &T)

	for i:=0 ; i<T ; i++{
		var N int
		fmt.Fscan(reader, &N)
		MBTIList := make([]string, N) 
		for j:=0 ; j<N ; j++ {
			fmt.Fscan(reader, &MBTIList[j])
		}
		if N > 64 {
			fmt.Fprintln(writer, 0)
			continue
		}
		min := 16
		for j:=0 ; j<N-2 ; j++ {
			for k:=j+1 ; k<N-1 ; k++ {
				for l:=k+1 ; l<N ; l++ {
					result := 0
					for m:=0 ; m<4 ; m++ {
						if MBTIList[j][m] == MBTIList[k][m] && MBTIList[k][m] == MBTIList[l][m] && MBTIList[l][m] == MBTIList[j][m] {
							continue
						}
						result+=2
					}
					if min > result {
						min = result
					}
				}
			}
		}
		fmt.Fprintln(writer, min)
	}
}