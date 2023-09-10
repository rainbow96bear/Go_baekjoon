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
	cnt := 0
	for i:=1 ; i<=N ; i++{
		num := i
		
		if num < 10 {
			cnt++
			continue
		}

		check := true
		gap := num%10 - (num/10)%10
		
		for num/10 != 0{
			if num%10 - (num/10)%10 != gap {
				check = false 
				break
			}
			num /= 10
		}
		if check {
			cnt++
		}
	}
	fmt.Fprint(writer, cnt)
}