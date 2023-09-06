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
	var maxX, maxY int = -10000, -10000
	var minX, minY int = 10000, 10000
	for i:=0 ; i<N ; i++{
		var x, y int
		fmt.Fscan(reader, &x, &y)
		if maxX < x {
			maxX = x
		}
		if maxY < y {
			maxY = y
		}
		if minX > x {
			minX = x
		}
		if minY > y {
			minY = y
		}
	}
	fmt.Fprint(writer, (maxX-minX) * (maxY-minY))
}