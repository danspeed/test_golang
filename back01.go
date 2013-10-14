package main

import (
	"fmt"
)

const (
	pCounts = 6
	pCap    = 27
)

var V [pCounts + 1][pCap]int
var weight, price, x [pCounts]int

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func KnapSack() {
	for i := 0; i < pCounts; i++ {
		V[i][0] = 0
	}
	for i := 0; i < pCap; i++ {
		V[0][i] = 0
	}
	for i := 0; i < pCounts; i++ {
		//fmt.Printf("i=%d : %d\n", i, price[i])
		for j := 0; j < pCap; j++ {
			if j < weight[i] {
				V[i+1][j] = V[i][j]
			} else {
				V[i+1][j] = max(V[i][j], V[i][j-weight[i]]+price[i])
			}
		}
		//fmt.Println(i, " || ", V[i])
	}

	fmt.Println("=======")

	j := pCap - 1
	for i := pCounts - 1; i >= 0; i-- {
		//fmt.Printf("==>V[%d][%d]=%d,V[%d][%d]=%d\n", i+1, j, V[i+1][j], i, j, V[i][j])
		if V[i+1][j] > V[i][j] {
			x[i] = 1
			j = j - weight[i]
		} else {
			x[i] = 0
		}
	}

	//fmt.Println("x:", x)
}

func printResult() {
	fmt.Printf("Max Value:\t")
	for i := 0; i < pCounts-1; i++ {
		fmt.Printf("%d * %d + ", x[i], price[i])
	}
	fmt.Printf("%d * %d = %d\n", x[pCounts-1], price[pCounts-1], V[pCounts][pCap-1])
	fmt.Printf("Max Cap:\t")
	maxcap := 0
	for i := 0; i < pCounts-1; i++ {
		fmt.Printf("%d * %d + ", x[i], weight[i])
		if x[i] > 0 {
			maxcap += weight[i]
		}
	}
	fmt.Printf("%d * %d = %d\n", x[pCounts-1], weight[pCounts-1], maxcap+weight[pCounts-1])
}

func main() {
	weight = [pCounts]int{2, 9, 7, 10}
	price = [pCounts]int{2, 15, 9, 10}
	KnapSack()
	printResult()
}
