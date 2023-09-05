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

	for {
		var input int
		fmt.Fscan(reader, &input)
		if input == -1 {
			break
		}
		var check int = 0
		var text string
		for i:=1 ; i<input ; i++{
			if input%i==0{
				check+=i
				text = fmt.Sprintf("%s %d +", text,i)
			}
		}
		if check == input {
			text = text[:len(text)-2]
			var rst string = fmt.Sprintf("%d =%s",input,text)
			fmt.Fprintln(writer,rst)
		}else {
			var rst string = fmt.Sprint(input," is NOT perfect.")
			fmt.Fprintln(writer,rst)
		}
	}
}