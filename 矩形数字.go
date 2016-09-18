package main

import (
	"flag"
	"fmt"
)

// 输出类似以下图形，给定中心点为4
/*
 1  1  1  1  1  1  1
 1  2  2  2  2  2  1
 1  2  3  3  3  2  1
 1  2  3  4  3  2  1
 1  2  3  3  3  2  1
 1  2  2  2  2  2  1
 1  1  1  1  1  1  1
*/

var (
	num int = 4
)

func init() {
	flag.IntVar(&num, "num", 4, "num:0-100")
	flag.Parse()
}

func main() {
	for row := 1; row < 2*num; row++ {
		for column := 1; column < 2*num; column++ {
			rowDu := abs(row - num)
			columnDu := abs(column - num)
			if rowDu > columnDu {
				fmt.Printf("%2d ", num-rowDu)
			} else {
				fmt.Printf("%2d ", num-columnDu)
			}
		}
		fmt.Println("")
	}
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
