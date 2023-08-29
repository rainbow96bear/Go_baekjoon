package main

import (
	"bufio"
	"fmt"
	"os"
)

var rx = bufio.NewReader(os.Stdin)
var wx = bufio.NewWriter(os.Stdout)
var check0, check1 int

func main() {

	defer wx.Flush()

	var testNum int

	fmt.Fscanln(rx, &testNum)

	for i:=0 ; i<testNum ; i++{
		var num int
		fmt.Fscanln(rx,&num)
		check0, check1 = fibonachi(num)
		fmt.Fprintln(wx,check0, check1 )
	}
}

func fibonachi(n int) (check0, check1 int) {
	if n == 0 {
		return 1, 0
	}else if n == 1 {
		return 0, 1
	}else {
		var previous, now int = 1, 1
		for i:=0;i<n-2;i++{
			previous, now = now, previous + now
		}
		return previous, now
	}
}