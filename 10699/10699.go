package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main(){
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	currentTime := time.Now().Add(9*time.Hour)

	formattedDate := currentTime.Format("2006-01-02")
	fmt.Fprintln(writer,formattedDate)
}