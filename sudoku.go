package main

import "fmt"

func numExistBox(num int, row int, column int, grid [9][9]int) bool {
	var baseRow = row - (row % 3)
	var baseColumn = column - (column % 3)
	for r := baseRow; r < baseRow + 3; r++ {
		for c := baseColumn; c < baseColumn + 3; c++ {
			if grid[r][c] == num {
				return true
			}
		}
	}
	return false
}

func numExistRow(num int, row int, grid [9][9]int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return true
		}
	}
	return false
}

func numExistColumn(num int, column int, grid [9][9]int) bool {
	for i := 0; i < 9; i++ {
		if grid[i][column] == num {
			return true
		}
	}
	return false
}

func isValid(num int, row int, column int, grid [9][9]int) bool {
	var inRow = numExistRow(num, row, grid)
	var inColumn = numExistColumn(num, column, grid)
	var inBox = numExistBox(num, row, column, grid)
	return !inRow && !inColumn && !inBox
}

func nextEmpty(grid [9][9]int) (int, int) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if grid[r][c] == 0 {
				return r, c
			}
		}
	}
	return -1, -1
}

func isSolved(grid [9][9]int) bool {
	for r := 0; r < 9; r++ {
		for i := 0; i < 9; i++ {
			if !numExistRow(i + 1, r, grid) {
				return false
			}
		}
	}
	return true
}

func solve(grid [9][9]int, num int) [9][9]int {
	// base case full grid
	var r, c int = nextEmpty(grid)
	if r == -1 && c == -1 {
		return grid
	}

	for num <= 9 {
		if isValid(num, r, c, grid) {
			grid[r][c] = num
			grid = solve(grid, 1)
		}
		num++
	}
	if !isSolved(grid) {
		grid[r][c] = 0
	}
	return grid
}

func printGrid(grid [9][9]int) {
	for i := 0; i < 9; i++ {
		fmt.Println(grid[i])
	}
}

func main() {
	// empty sudoku grid [row][column]
	var grid = [9][9]int{}

	// hard coded puzzle for testing
	// to be replaced with generator, file or user entry
	grid[0][1] = 7
	grid[0][7] = 2

	grid[1][0] = 5
	grid[1][4] = 8
	grid[1][8] = 4

	grid[2][2] = 2
	grid[2][5] = 6
	grid[2][6] = 9

	grid[3][0] = 3
	grid[3][3] = 6
	grid[3][6] = 5

	grid[4][1] = 5
	grid[4][4] = 4
	grid[4][7] = 8

	grid[5][2] = 6
	grid[5][5] = 1
	grid[5][8] = 9

	grid[6][2] = 5
	grid[6][3] = 4
	grid[6][6] = 2

	grid[7][0] = 4
	grid[7][4] = 6
	grid[7][8] = 8

	grid[8][1] = 9
	grid[8][7] = 3

	printGrid(solve(grid, 1))
}