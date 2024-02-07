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

	var (
		W, H, X, Y, P int
		count int
	)
	fmt.Fscanf(reader, "%d %d %d %d %d\n", &W, &H, &X, &Y, &P)

	for i:=0 ; i<P ; i++ {
		var x, y int
		fmt.Fscanf(reader, "%d %d\n", &x, &y)
		if x < X {
			if (H/2) * (H/2) >= (X-x)*(X-x) + (Y+(H/2)-y)*(Y+(H/2)-y) {
				count++
			}
		} else if x > X+W{
			if (H/2) * (H/2) >= (X+W-x)*(X+W-x) + (Y+(H/2)-y)*(Y+(H/2)-y) {
				count++
			}
		}else if y >= Y && y <= Y+H {
			count++
		}
	}
	fmt.Fprintln(writer, count)
}