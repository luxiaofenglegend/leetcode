package utils

import "fmt"

type PrintUtil struct {
}

func (p PrintUtil) PrintArrInt(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func (p PrintUtil) Print2DInt(mat [][]int) {
	for _, arr := range mat {
		p.PrintArrInt(arr)
	}
}

func (p PrintUtil) Print3DimInt(cube [][][]int) {
	for i, _ := range cube {
		for j, _ := range cube[i] {
			if len(cube[i][j]) == 0 {
				fmt.Print("*")
			} else {
				for _, v := range cube[i][j] {
					fmt.Print(v, "-")
				}
			}

			fmt.Print(" ")
		}
		fmt.Println()
	}
}
