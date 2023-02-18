package main

import (
	"fmt"

	lib "github.com/nurfan27/number-of-island/lib/numberofislands"
)

func main() {

	var grid [][]int = [][]int{
		{1, 1, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{1, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1},
	}

	land := lib.NewNumberOfIsland()
	output := land.CountIslands(grid)
	fmt.Println(output)
}
