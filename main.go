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

func pressNumber(matrix [][]int, row, col int) {
	adjacent := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for _, pos := range adjacent {
		// apply the transform matrix
		newRow, newCol := row+pos[0], col+pos[1]

		if newRow >= 0 && newRow < 3 && newCol >= 0 && newCol < 3 {
			matrix[newRow][newCol] = (matrix[newRow][newCol] + 1) % 3
		}
	}

	// now update the button number itself
	matrix[row][col] = (matrix[row][col] + 1) % 3
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

func solveUsingGraph(matrix [][]int) int {

	queue := []State{{Matrix: matrix, Steps: 0}}
	visited := make(map[string]bool)
	var finalSteps int
	var finalMatrix [][]int

	for len(queue) > 0 {
		currentState := queue[0]
		queue = queue[1:]

		if isMatrixZero(currentState.Matrix) {
			fmt.Println("\nOptimal Solution (Steps:", currentState.Steps, "):")
			printMatrix(currentState.Matrix)
			return currentState.Steps + 1
		}

		// Generate next states by pressing numbers
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				nextState := currentState
				pressNumber(nextState.Matrix, i, j)

				// Check if the next state has been visited
				hash := fmt.Sprintf("%v", nextState.Matrix)
				if !visited[hash] {
					visited[hash] = true
					nextState.Steps++
					queue = append(queue, nextState)
				}
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
