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

	minute = minute + hour*60

	alamHour := (minute - 45)/60
	alamMinute := (minute - 45)%60
	if alamMinute < 0 {
		alamHour--
		alamMinute += 60
	}
	if alamHour < 0 {
		alamHour = 23
	}

	fmt.Fprintf(writer,"%d %d",alamHour, alamMinute)
}