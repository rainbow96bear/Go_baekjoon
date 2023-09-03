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

	var hour, minute int
	fmt.Fscan(reader, &hour, &minute)

	var time int
	fmt.Fscan(reader, &time)

	minute += time

	hour += minute/60
	hour = hour%24
	minute = minute%60

	fmt.Fprintf(writer, "%d %d", hour, minute)
}