package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Star struct {
	x float64
	y float64
}

type Bridge struct {
	IDX1 int
	IDX2 int
	length float64
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int 
	fmt.Fscanf(reader, "%d\n", &n)

	starList := []Star{}
	for i:=0 ; i<n ; i++ {
		var x, y float64
		fmt.Fscanf(reader, "%v %v\n", &x, &y)
		starList = append(starList, Star{x, y})
	}
	result := FindMinValue(n, starList)
	fmt.Fprintln(writer, float64(int(result*100))/100)
}

func FindMinValue(n int, starList []Star) float64 {
	bridgeList := []Bridge{}
	for i:=0 ; i<len(starList) ; i++ {
		for j:=i+1 ; j<len(starList) ; j++ {
			if i != j {
				bridgeList = append(bridgeList, Bridge{i, j, CalcLength(starList[i], starList[j])})
			}
		}
	}
	sort.Slice(bridgeList, func(i, j int)bool{return bridgeList[i].length < bridgeList[j].length})
	return FindRoute(n, bridgeList)
}

func CalcLength(star1, star2 Star) float64 {
	Xgap := star1.x - star2.x
	Ygap := star1.y - star2.y
	length := math.Sqrt(Xgap*Xgap + Ygap*Ygap)
	return length
}

func FindRoute(n int, BridgeList []Bridge) float64 {
	Parent := make(map[int]int)
	for i:=0 ; i<n ; i++ {
		Parent[i] = i
	}
	var total float64
	for i:=0; i<len(BridgeList) ; i++ {
		Parent1 := FindParent(BridgeList[i].IDX1, Parent)
		Parent2 := FindParent(BridgeList[i].IDX2, Parent)
		if Parent[Parent1] != Parent[Parent2] {
			total += BridgeList[i].length
			if BridgeList[i].IDX1 < BridgeList[i].IDX2 {
				Parent[Parent1] = Parent2
			}else {
				Parent[Parent2] = Parent1
			}
		}
	}
	return total
}

func FindParent(num int, Parent map[int]int) int {
	if num != Parent[num] {
		Parent[num] = FindParent(Parent[num], Parent)
		return Parent[num] 
	}
	return Parent[num]
}