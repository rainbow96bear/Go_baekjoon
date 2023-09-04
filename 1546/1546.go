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

	var ScoreNum int
	fmt.Fscan(reader, &ScoreNum)

	var sum, max float64
	for i:=0 ; i<ScoreNum ; i++{
		var score float64
		fmt.Fscan(reader, &score)
		if max < score {
			max = score
		}
		sum += score
	}
	fmt.Fprintln(writer, sum/float64(ScoreNum)*100/max)
}