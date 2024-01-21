package main

import (
	"fmt"
	"math"
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

func newEmpty3x3() [][]int {

	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}

	return matrix
}

func solveUsingGraph(matrix [][]int) int {
	newMatrix := newEmpty3x3()
	copy(newMatrix, matrix)

	// visited := make(map[string]bool)
	var finalSteps int
	var finalMatrix [][]int
	queue := make([]int, 0)
	start := 0

	for i := 0; i < 9; i++ {
		newRow := i / 3
		newCol := i % 3
		start += newMatrix[newRow][newCol]
	}
	// capacity equal to 1 shifted m * n (3 x 3)
	dist := make([]int, 1<<9)
	dist[0] = math.MaxInt
	for i := 1; i < len(dist); i *= 2 {
		copy(dist[i:], dist[:i])
	}

	dist[start] = 0

	queue = append(queue, start)
	for len(queue) > 0 {
		finalSteps++
		// initialize steps to 1 and our current state to the original matrix
		// currentState := State{Matrix: newMatrix, Steps: 1}
		currentNode := queue[0]
		queue = queue[1:]

		d := dist[currentNode]
		if currentNode == 0 {
			return d
		}

		for i := 0; i < 9; i++ {
			row := i / 3
			col := i % 3

			next := currentNode
			next ^= 1 << i
			if col > 0 {
				next ^= 1 << (i - 1)
			}

			if col < 2 { // n - 1 :(n = 3)
				next ^= 1 << (i + 1)
			}

			if row > 0 {
				next ^= 1 << (i - 3)
			}

			if row < 2 {
				next ^= 1 << (i + 3)
			}

			if d+1 < dist[next] {
				dist[next] = d + 1
				queue = append(queue, next)
			}
		}
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
