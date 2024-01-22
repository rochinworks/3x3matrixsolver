package main

import (
	"fmt"
	"math/rand"
	"time"
)

type State struct {
	Matrix [][]int
	Steps  int
}

func pressNumber(matrix [][]int, row, col int) int {
	adjacent := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// if matrix[row][col] == 0 {
	// 	return -1
	// }

	// for _, pos := range adjacent {
	// 	newRow, newCol := row+pos[0], col+pos[1]

	// 	if newRow >= 0 && newRow < 3 && newCol >= 0 && newCol < 3 {
	// 		if matrix[newRow][newCol] == 0 {
	// 			return -1
	// 		}
	// 	}
	// }

	for _, pos := range adjacent {
		// apply the transform matrix
		newRow, newCol := row+pos[0], col+pos[1]

		if newRow >= 0 && newRow < 3 && newCol >= 0 && newCol < 3 {
			matrix[newRow][newCol] = (matrix[newRow][newCol] + 1) % 3
		}
	}

	// now update the button number itself
	matrix[row][col] = (matrix[row][col] + 1) % 3
	return 0
}

func isMatrixZero(matrix [][]int) bool {
	for _, row := range matrix {
		for _, val := range row {
			if val != 0 {
				return false
			}
		}
	}
	return true
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func randomStart(matrix [][]int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(3)
		}
	}
}

func newEmpty3x3() [][]int {

	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}

	return matrix
}

func solveUsingGraph(matrix [][]int) int {
	// newMatrix := newEmpty3x3()
	// copy(newMatrix, matrix)

	// visited := make(map[string]bool)
	// state := []State{{Matrix: newMatrix, Steps: 1}}
	// queue := make([][]int, 1)

	// for len(queue) > 0 {
	// 	currentIndex := queue[0]
	// 	// pop the bottom
	// 	queue = queue[1:]
	// 	currentState := state[len(state)-1]
	// 	nextIndex := currentIndex

	// 	for i := 0; i < 9; i++ {
	// 		row := i / 3
	// 		col := i % 3
	// 		if currentState.Matrix[row][col] == 0 {
	// 			// do nothing
	// 			continue
	// 		}

	// 		//

	// 		// nextIndex = []int{newRow, newColumn}
	// 	}
	// }
	return -1
}

func main() {

	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}

	randomStart(matrix)
	fmt.Println("Random Start:")
	printMatrix(matrix)

	var keepProgramRunning bool = true
	for keepProgramRunning {

		// stepstaken := solveUsingGraph(matrix)
		fmt.Println("enter new row coordinate:")
		var row, col int
		fmt.Scanln(&row)
		fmt.Println("enter new col coordinate:")
		fmt.Scanln(&col)
		if x := pressNumber(matrix, row, col); x == -1 {
			fmt.Println("nothing pressed")
		}
		fmt.Println("New Matrix:")
		printMatrix(matrix)
		// fmt.Println(stepstaken)
	}
}
