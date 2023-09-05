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
	checkX := []int{}
	checkY := []int{}
	for i:=0 ; i<3; i++{
		var x, y int
		fmt.Fscan(reader, &x, &y)
		checkX = append(checkX,x)
		checkY = append(checkY,y)
	}
	for i:=0 ; i<3 ; i++{
		if checkX[i%3] == checkX[(i+1)%3]{
			fmt.Fprint(writer, checkX[(i+2)%3])
		}
	}
	fmt.Fprint(writer," ")
	for i:=0 ; i<3 ; i++{
		if checkY[i%3] == checkY[(i+1)%3]{
			fmt.Fprint(writer, checkY[(i+2)%3])
		}
	}
}