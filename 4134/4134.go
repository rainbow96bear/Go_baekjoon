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
	for i:=0 ; i<N ; i++{
		var input uint
		fmt.Fscan(reader, &input)
		if input <= 2{
			fmt.Fprintln(writer, 2)
			continue
		}
		for j:=input ; ; j++ {
			var k uint
			check := true
            for k = 2 ; k*k<=j ; k++{
				if j%k == 0 {
					check = false
					break
				}
			}
			if check {
				fmt.Fprintln(writer, j)
				break
			}
		}
	}
}