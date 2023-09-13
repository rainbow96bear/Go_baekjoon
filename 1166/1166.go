package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	len := make([]float64,3)
	fmt.Fscan(reader, &N, &len[0], &len[1], &len[2])
	
	left, right := 0.0, math.MaxFloat64
	rst := ""
	for i:=0; i<10000 ; i++{
		mid := (left+right)/2

		if int(len[0]/mid) * int(len[1]/mid) * int(len[2]/mid) >= N {
			left = mid
			rst = fmt.Sprintf("%.10f",mid)
		} else {
			right = mid
		}
	}
	fmt.Fprint(writer,rst)
}