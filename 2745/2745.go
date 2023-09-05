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

	var N string
	var B int
	checkPoint := "0A"
	var answer float64
	fmt.Fscan(reader, &N, &B)
	for i:=0 ; i<len(N) ; i++ {
		if int(N[i])-int(checkPoint[1]) >= 0 {
			answer+=float64(N[i]-checkPoint[1]+10) * (math.Pow(float64(B), float64(len(N)-1-i)))
		}else if int(N[i])-int(checkPoint[0]) >= 0 {
			answer+=float64(N[i]-checkPoint[0]) * (math.Pow(float64(B), float64(len(N)-1-i)))
		}
	}
	fmt.Fprint(writer,int64(answer))
}
