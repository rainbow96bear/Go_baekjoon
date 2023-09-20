package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
func main(){

	defer writer.Flush()

	var M int
	fmt.Fscan(reader, &M)

	list := make([]bool, 21)
	for i:=0 ; i<M ; i++ {
		var command string
		var num int
		fmt.Fscan(reader, &command)

		if command != "all" && command != "empty" {
			fmt.Fscan(reader, &num)
		}
		process(list, command, num)
	}
}

func process(list []bool, command string, num int){
	switch command {
	case "add" :
		list[num] = true
	case "remove" :
		list[num] = false
	case "check" :
		if list[num] == true {
			fmt.Fprintln(writer, 1)
			return
		}
		fmt.Fprintln(writer, 0)
	case "toggle" :
		list[num] = !list[num]
	case "all" :
		for i:=1 ; i<=20 ; i++ {
			list[i] = true
		}
	case "empty" :
		for i:=1 ; i<=20 ; i++ {
			list[i] = false
		}
	}
}
