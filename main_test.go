package main

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPressNumber(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
		row      int
		col      int
	}{{
		name: "press middle number",
		matrix: [][]int{
			{0, 1, 2},
			{1, 2, 0},
			{1, 1, 1},
		},
		expected: [][]int{
			{0, 2, 2},
			{2, 0, 1},
			{1, 2, 1},
		},
		row: 1,
		col: 1,
	}, {
		name: "press top middle number",
		matrix: [][]int{
			{0, 1, 2},
			{1, 2, 0},
			{1, 1, 1},
		},
		expected: [][]int{
			{1, 2, 0},
			{1, 0, 0},
			{1, 1, 1},
		},
		row: 0,
		col: 1,
	}, {
		name: "press bottom middle number",
		matrix: [][]int{
			{0, 1, 2},
			{1, 2, 0},
			{1, 1, 1},
		},
		expected: [][]int{
			{0, 1, 2},
			{1, 0, 0},
			{2, 2, 2},
		},
		row: 2,
		col: 1,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pressNumber(test.matrix, test.row, test.col)
			assert.DeepEqual(t, test.matrix, test.expected)

		})
	}
}

func TestIsMatrixZero(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected bool
	}{{
		name: "matrix is not all zeros",
		matrix: [][]int{
			{0, 1, 2},
			{1, 2, 0},
			{1, 1, 1},
		},
		expected: false,
	}, {
		name: "matrix is all zeros",
		matrix: [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
		expected: true,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			isZero := isMatrixZero(test.matrix)
			assert.Equal(t, isZero, test.expected)

		})
	}
}

func TestPrintMatrix(t *testing.T) {
	matrix := [][]int{
		{0, 1, 0},
		{2, 0, 1},
		{1, 1, 0},
	}

	printMatrix(matrix)
}

func TestRandomStart(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	randomStart(matrix)
	printMatrix(matrix)
}

func TestMain(t *testing.T) {

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
