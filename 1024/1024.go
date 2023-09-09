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

	var N, L int
	fmt.Fscan(reader, &N, &L)
	check := false 
	for i:=L ; i<=100 ; i++{
		if i%2 == 1{
			if N%i == 0 && N/i - i/2 >= 0{
				for j:=-(i/2) ; j<=i/2 ; j++ {
					fmt.Fprintf(writer, "%d ",N/i+j)
				}
				check=true
				break
			}
		}else if i%2 == 0 {
			if N%i == i/2 && (N/i - i/2) + 1 >= 0{
				for j:=-(i/2) ; j<i/2 ; j++ {
					fmt.Fprintf(writer, "%d ",N/i+j+1)
				}
				check=true
				break
			}
		}
	}
	if check == false {
		fmt.Fprint(writer, -1)
	}
}