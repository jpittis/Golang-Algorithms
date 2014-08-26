package main

import "fmt"
import "math/rand"

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
	var rNew, cNew = getClose(maze, r, c)
	for !(rNew == -1 && cNew == -1) {
		// break down walls
		breakWallsBetween(&maze[r][c], r, c, &maze[rNew][cNew], rNew, cNew)
		maze = backtrack(maze, rNew, cNew)
		rNew, cNew = getClose(maze, r, c)
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
	var close [][]int
	// hard coded north
	if inBounds(maze, r - 1, c) && notVisited(maze[r - 1][c]) {
		close = append(close, []int{r - 1, c})
	}

	// hard coded south
	if inBounds(maze, r + 1, c) && notVisited(maze[r + 1][c]) {
		close = append(close, []int{r + 1, c})
	}

	// hard coded east
	if inBounds(maze, r, c - 1) && notVisited(maze[r][c - 1]) {
		close = append(close, []int{r, c - 1})
	}

	// hard coded west
	if inBounds(maze, r, c + 1) && notVisited(maze[r][c + 1]) {
		close = append(close, []int{r, c + 1})
	}

	//randomize if not empty
	if (len(close) > 0) {
		var returnData = close[rand.Intn(len(close))]
		return returnData[0], returnData[1]
	} else {
		return -1 , -1
	}
}

func inBounds(maze [][]cell, r int, c int) bool {
	return r >= 0 && r < len(maze) && c >= 0 && c < len(maze[r])
}

func notVisited(c cell) bool {
	return !c.north && !c.south && !c.east && !c.west
}

// print a cell and its walls
func PrintMaze(maze [][]cell) {
	for r := 0; r < len(maze); r++ {
		for c := 0; c < len(maze[r]); c++ {
			if maze[r][c].north {
				fmt.Print("# ")
			} else {
				fmt.Print("##")
			}
		}
		fmt.Print("#")
		fmt.Println()
		for c := 0; c < len(maze[r]); c++ {
			if maze[r][c].west {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
			fmt.Print(" ")
			/*if maze[r][c].east  {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}*/
		}
		fmt.Print("#")
		fmt.Println()
		/*for c := 0; c < len(maze[r]); c++ {
			if maze[r][c].south {
				fmt.Print("# #")
			} else {
				fmt.Print("###")
			}
		}*/
	}
	for c := 0; c < len(maze[0]); c++ {
		fmt.Print("##")
	}
	fmt.Println("#")
}

func main() {
	PrintMaze(create(25, 25))
}