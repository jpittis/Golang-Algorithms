package main

import "fmt"
//import "math/rand"

type cell struct {
	north bool
	south bool
	east bool
	west bool
}

func create(width int, height int) [][]cell {
	var maze = make([][]cell, height)
	for r := 0; r < len(maze); r++ {
		for c := 0; c < width; c++ {
			maze[r] = append(maze[r], cell{})
		}
	}

	return backtrack(maze, 1, 1)
}

func backtrack(maze [][]cell, r int, c int) [][]cell {
	// base case
	if noMoreToVisit(maze) {
		return maze
	}

	var rNew, cNew = getClose(maze, r, c)
	if rNew == -1 && cNew == -1 {
		return maze
	} else {
		// break down walls
		breakWallsBetween(&maze[r][c], r, c, &maze[rNew][cNew], rNew, cNew)
		maze = backtrack(maze, rNew, cNew)
	}

	return maze
}

func breakWallsBetween(cell1 *cell, r1 int, c1 int, cell2 *cell, r2 int, c2 int) {
	if r1 == r2 {
		if c2 > c1 {
			cell1.east = true
			cell2.west = true
		} else {
			cell1.west = true
			cell2.east = true
		}
	} else {
		if r2 > r1 {
			cell1.south = true
			cell2.north = true
		} else {
			cell1.north = true
			cell2.south = true
		}
	}
}

func getClose(maze [][]cell, r int, c int) (int, int) {
	var close = [4][2]int{
		{-1, -1,},
		{-1, -1,},
		{-1, -1,},
		{-1, -1,},
	}
	// hard coded north
	if inBounds(maze, r - 1, c) && notVisited(maze[r][c]) {
		close[0][0] = r - 1
		close[0][1] = c
	}

	// hard coded south
	if inBounds(maze, r + 1, c) && notVisited(maze[r][c]) {
		close[1][0] = r + 1
		close[1][1] = c
	}

	// hard coded east
	if inBounds(maze, r, c - 1) && notVisited(maze[r][c]) {
		close[2][0] = r
		close[2][1] = c - 1
	}

	// hard coded west
	if inBounds(maze, r, c + 1) && notVisited(maze[r][c]) {
		close[3][0] = r
		close[3][1] = c + 1
	}
	//this is not randomized for testing purposes
	for i := 0; i < 4; i++ {
		if !(close[i][0] == -1 && close[i][1] == -1) {
			return close[i][0], close[i][1]
		}
	}
	return -1, -1
}

func inBounds(maze [][]cell, r int, c int) bool {
	return r >= 0 && r < len(maze) && c >= 0 && c < len(maze[r])
}

func noMoreToVisit(maze [][]cell) bool {
	for r := 0; r < len(maze); r++ {
		for c := 0; c < len(maze[r]); c++ {
			if notVisited(maze[r][c]) {
				return false
			}
		}
	}
	return true
}

func notVisited(c cell) bool {
	return !c.north && !c.south && !c.east && !c.west
}

// print a cell and its walls
func PrintMaze(maze [][]cell) {
	for r := 0; r < len(maze); r++ {
		for c := 0; c < len(maze[r]); c++ {
			if maze[r][c].north {
				fmt.Print("   ")
			} else {
				fmt.Print(" # ")
			}
		}
		fmt.Println()
		for c := 0; c < len(maze[r]); c++ {
			if maze[r][c].west {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
			fmt.Print("0")
			if maze[r][c].east  {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
		for c := 0; c < len(maze[r]); c++ {
			if maze[r][c].south {
				fmt.Print("   ")
			} else {
				fmt.Print(" # ")
			}
		}
		fmt.Println()
	}
}

func main() {
	PrintMaze(create(3, 3))
}