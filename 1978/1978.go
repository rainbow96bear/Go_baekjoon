package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)
	answer := 0
	for i:=0 ; i<N ; i++{
		var input int
		fmt.Fscan(reader, &input)
		cnt := 0
		if input == 1{
			continue
		}
		for j:=2 ; j <= int(math.Pow(float64(input),0.5)) ; j++{
			if j != input && input % j == 0 {
				cnt++
				break
			}
		}
		if cnt == 0{
			answer++
		}
	}
	fmt.Fprint(writer, answer)
}