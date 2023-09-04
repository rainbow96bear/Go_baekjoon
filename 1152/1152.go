package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	s,_ := reader.ReadString('\n')
	// s = strings.TrimSpace(s)
	stringSlice := strings.Split(s," ")
	
	answer := len(stringSlice)
	
	fmt.Fprintln(writer, answer)
}