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

	for {
		var N int
		fmt.Fscan(reader, &N)
		if N == 0 {
			break
		}else{
			cnt := 0
			for i:=N+1 ; i<=2*N ; i++ {
				if i == 2 {
					cnt=1
					break
				}
				check := false
				for j:=2 ; j*j<=i ; j++{
					if i%j == 0 {
						check = true
						break
					}
				}
				if check == false {
					cnt++
				}
			}
			fmt.Fprintln(writer, cnt)
		}
	}
}