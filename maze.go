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