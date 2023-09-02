package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct{
	previous int
	next int
	value int
}
type CircleList struct{
	list []node
}

func (c *CircleList)push(n node){
	n.previous=len(c.list)-1
	c.list[n.previous].next = len(c.list)
	c.list = append(c.list, n)
}
func (c *CircleList)remove(n node, index int){	
	c.list = append(c.list[:index], c.list[index+1:]...)
	c.list[len(c.list)-1].next = 0
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	
	defer writer.Flush()

	var testNum int
	fmt.Fscan(reader,&testNum)

	for i:=0 ; i<testNum ; i++{
		var areaNum int
		var members int
		fmt.Fscan(reader, &areaNum, &members)
		answer := members
		var enemyNums []CircleList
		var check [][]int
		for j:=0 ; j<2 ; j++{
			for k:=0 ; k<areaNum ; k++{
				fmt.Fscan(reader, &enemyNums[j].list[k])
				if enemyNums[j].list[k].value <= members/2 {
					check[j][k] = 1
				}
			}
		}
		for j:=0 ; j<2 ; j++{
			for k:=0 ; k<areaNum ; k++{
				if check[j][k] == 1 {
					
					// 주변과 더해서 members보다 작은지 확인
					// 작다면 일단 대기
					// 작으면 작은 개수 체크
				}
			}
		}
		// 최소를 구하는 방법을 생각해야한다.
		// 경우의 수가 적은 것 부터 짝을 잡는다.
		// 합해서 members보다 작거나 같으면 answer--
		// 깊이 우선? 넓이 우선?
		fmt.Fprintln(writer, &answer)
	}
}