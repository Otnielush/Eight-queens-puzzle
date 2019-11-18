package main

import (
	"fmt"
)

const (
	dim = 14
)

type board struct {
	cellOccu    [dim][dim]bool
	Queens      [dim]possition
	Prohibitios [dim][dim][dim]int // row,column,num of queen
	ProhibCheck [dim][dim]bool
}

type possition struct {
	row    int
	column int // column num=>a-h
}

var (
	p     = fmt.Println
	pf    = fmt.Printf
	step  int
	bo    board
	clmns = map[int]string{
		0:  "a",
		1:  "b",
		2:  "c",
		3:  "d",
		4:  "e",
		5:  "f",
		6:  "g",
		7:  "h",
		8:  "i",
		9:  "j",
		10: "k",
		11: "l",
		12: "m",
		13: "n",
		14: "o",
		15: "p",
	}
	pathCurr    [dim]possition
	paths       [4000][dim]possition
	pathNum     int
	cellToMove  []int
	solutions   [2700][dim]possition
	solutionNum int
)

func main() {

	bo.walker(bo.FreePlaces(0))

	p("Paths walked:", pathNum)
	p("Solutions found:", solutionNum)

}

func (bo *board) addQueen(row int) {
	bo.cellOccu[row][step] = true
	bo.Queens[step].row = row
	bo.Queens[step].column = step
	pathCurr[step].row = row
	pathCurr[step].column = step
}

func (bo *board) calcProhib() {
	// nullizer
	for r := 0; r < dim; r++ {
		for c := 0; c < dim; c++ {
			bo.ProhibCheck[r][c] = false
			for i := 0; i < dim; i++ {
				bo.Prohibitios[r][c][i] = 0
			}
		}
	}

	// num queen
	for i := 0; i <= step; i++ {
		rQ := bo.Queens[i].row
		cQ := bo.Queens[i].column
		// p("rQ", rQ, " cQ", cQ)
		// rows
		for r := 0; r < dim; r++ {
			// give num of queen
			bo.Prohibitios[r][cQ][i] = i + 1
		}
		// columns
		for c := 0; c < dim; c++ {
			bo.Prohibitios[rQ][c][i] = i + 1
		}
		// diagonal
		for r, c := 1, 1; rQ+r < dim && cQ+c < dim; r, c = r+1, c+1 {
			bo.Prohibitios[rQ+r][cQ+c][i] = i + 1
		}
		for r, c := -1, 1; rQ+r >= 0 && cQ+c < dim; r, c = r-1, c+1 {
			bo.Prohibitios[rQ+r][cQ+c][i] = i + 1
		}
		for r, c := 1, -1; rQ+r < dim && cQ+c >= 0; r, c = r+1, c-1 {
			bo.Prohibitios[rQ+r][cQ+c][i] = i + 1
		}
		for r, c := -1, -1; rQ+r >= 0 && cQ+c >= 0; r, c = r-1, c-1 {
			bo.Prohibitios[rQ+r][cQ+c][i] = i + 1
		}
	}

	// cells not under attack
	for r := 0; r < dim; r++ {
		for c := 0; c < dim; c++ {
			for i := 0; i <= step; i++ {
				if bo.Prohibitios[r][c][i] != 0 {
					bo.ProhibCheck[r][c] = true
					break
				}
			}
		}
	}
}

func (bo *board) printProh() {
	var cellS string
	for r := dim - 1; r >= 0; r-- {
		pf("\n")
		pf("%d ", r+1)
		for rows := 1; rows <= 2; rows++ {
			if rows == 2 {
				pf("  ")
			}
			for c := 0; c < dim; c++ {
				for i := 4*rows - 4; i <= 4*rows-1; i++ {
					if bo.Prohibitios[r][c][i] != 0 {
						cellS += fmt.Sprintf("%d,", bo.Prohibitios[r][c][i])
					} else {
						cellS += "  "
					}
				}
				pf("[%s]", cellS)
				cellS = ""
			}
			pf("\n")
		}
		// pf("\n")
	}
	p("      a)        b)        c)        d)        e)        f)        g)        h)\n")
}

func (bo *board) printPrChech() {
	for r := dim - 1; r >= 0; r-- {
		pf("%d ", r+1)
		if r < 9 {
			pf(" ")
		}
		for c := 0; c < dim; c++ {
			if bo.cellOccu[r][c] {
				pf("[Q]")
			} else if bo.ProhibCheck[r][c] {
				pf("[-]")
			} else {
				pf("[ ]")
			}
		}
		pf("\n")
	}
	pf("    a  b  c  d  e  f  g  h")
	if dim > 8 {
		for i := 8; i < dim; i++ {
			pf("  %s", clmns[i])
		}
	}
	pf("\n")
}

func (bo *board) FreePlaces(col int) []int {
	free := make([]int, 0, dim)
	for i := 0; i < dim; i++ {
		if !bo.ProhibCheck[i][col] {
			free = append(free, i)
		}
	}
	return free
}

func (bo *board) queenNullizer() {
	for r := 0; r < dim; r++ {
		for c := 0; c < dim; c++ {
			bo.cellOccu[r][c] = false
		}
	}
}

func (bo *board) remQueen() {
	bo.cellOccu[bo.Queens[step].row][bo.Queens[step].column] = false
	pathCurr[step].row = 0
	pathCurr[step].column = 0
}

func (bo *board) walker(freeCells []int) {
	for _, r := range freeCells {
		bo.addQueen(r)
		bo.calcProhib()

		if step == dim-1 {
			p("Path found! step:", pathNum+1, "solution:", solutionNum+1)
			for i := 0; i < dim; i++ {
				pf("{%s%d}", clmns[pathCurr[i].column], pathCurr[i].row+1)
			}
			pf("\n")
			bo.printPrChech()
			if dim < 10 {
				paths[pathNum] = pathCurr
			}
			pathNum++
			if dim < 12 {
				solutions[solutionNum] = pathCurr
			}
			solutionNum++
			bo.remQueen()
			continue
		}

		nextFCells := bo.FreePlaces((step) + 1)
		if len(nextFCells) == 0 {
			if dim < 10 {
				pf("%d Stuck on step: %d path: ", pathNum+1, step+1)
				for i := 0; i <= step; i++ {
					pf("{%s%d}", clmns[pathCurr[i].column], pathCurr[i].row+1)
				}
				pf("\n")

				paths[pathNum] = pathCurr
			}
			pathNum++
			bo.remQueen()
			continue
		}
		step++
		bo.walker(nextFCells)

		bo.remQueen()
	}
	// step back
	step--
}
