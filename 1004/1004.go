package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	rx := bufio.NewReader(os.Stdin)
	wx := bufio.NewWriter(os.Stdout)
	
	defer wx.Flush()
	var testNum int

	fmt.Fscan(rx, &testNum)

	for i:=0 ; i<testNum ; i++ {
		var answer int = 0
		var x1, y1, x2, y2 int
		fmt.Fscan(rx, &x1, &y1, &x2, &y2)
		var planetsNum int
		fmt.Fscan(rx,&planetsNum)
		for j := 0 ; j<planetsNum ; j++ {
			var x, y, r int
			fmt.Fscan(rx, &x, &y, &r)
			if checkIn(x1, y1, x, y, r) != checkIn(x2, y2, x, y, r) {
				answer++
			}
		}
		fmt.Fprintln(wx, answer)
	}
}

func checkIn(cx1,cy1,x1,y1,r int) bool{
	return mathPow(x1-cx1) + mathPow(y1-cy1) < mathPow(r)
}

func mathPow(v int) int{
	return v*v
}