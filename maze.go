package main

import "fmt"

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


	return maze
}

func backtrack(maze [][]cell) {
	// base case
	if nextNotVisited(maze) == -1, -1 {
		return maze
	}
	return maze

}

func nextNotVisited(maze [][]cell) (int, int) {
	for r := 0; r < len(maze); r++ {
		for c := 0; c < len(maze[r]); c++ {
			if notVisited(maze[r][c]) {
				return r, c
			}
		}
	}
	return -1, -1
}

func notVisited(c cell) {
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