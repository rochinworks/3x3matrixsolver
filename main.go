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

	if matrix[row][col] == 0 {
		return -1
	}

	for _, pos := range adjacent {
		newRow, newCol := row+pos[0], col+pos[1]

		if newRow >= 0 && newRow < 3 && newCol >= 0 && newCol < 3 {
			if matrix[newRow][newCol] == 0 {
				return -1
			}
		}
	}

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

func new3x3() [][]int {

	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}

	return matrix
}

func solveUsingGraph(matrix [][]int) int {
	newMatrix := new3x3()
	copy(newMatrix, matrix)

	visited := make(map[string]bool)
	var finalSteps int
	var finalMatrix [][]int
	queue := make([][]int, 1) // next position to visit

	for len(queue) > 0 {
		// initialize steps to 1 and our current state to the original matrix
		currentState := State{Matrix: newMatrix, Steps: 1}

		if isMatrixZero(currentState.Matrix) {
			fmt.Println("\nOptimal Solution (Steps:", currentState.Steps, "):")
			printMatrix(currentState.Matrix)
			return currentState.Steps + 1
		}

		for i := 0; i < 9; i++ {
			hash := fmt.Sprintf("%v", []int{i, i % 3})
			if x := pressNumber(currentState.Matrix, i, i%3); x == -1 {
				if !visited[hash] {
					visited[hash] = true
					continue
				}
			} else {
				currentState.Steps++
				queue = append(queue, []int{i, i % 3})
			}

		}

		finalSteps = currentState.Steps + 1
		finalMatrix = currentState.Matrix
	}
	printMatrix(finalMatrix)
	return finalSteps
}

func main() {

	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}

	randomStart(matrix)
	fmt.Println("Random Start:")
	printMatrix(matrix)

	stepstaken := solveUsingGraph(matrix)
	fmt.Println(stepstaken)
}
