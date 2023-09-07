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

	var T int
	fmt.Fscan(reader, &T)

	for i:=0 ; i<T ;i++{
		var input string
		fmt.Fscan(reader, &input)

		strings.TrimSpace(input)
		strSli := strings.Split(input,"")
		check := []string{}
		do := true
		for _, v := range strSli {
			if v == "("{
				check = append(check, "(")
			}else if len(check) == 0 && v ==")"{
				fmt.Fprintln(writer, "NO")
				do = false
				break
			}else {
				check = check[:(len(check)-1)]
			}
		}
		if do {
			if len(check) > 0 {
				fmt.Fprintln(writer, "NO")
			}else {
				fmt.Fprintln(writer, "YES")
			}
		}
	}
}