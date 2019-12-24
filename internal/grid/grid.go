package grid

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

const headerFormat = "Generation: %d\tLiving cells: %d\n"

type Grid interface {
	Init()
	Populate(population int)
	Show(out io.Writer) error
	NextGeneration()
	LivingCells() int

	canCellLiveOn(x, y int) bool
	setCell(x, y int, value bool)
	getCell(x, y int) bool
	countLiveNeighbours(x, y int) (n int)
	getSize() (x, y int)
	getCellVisualOutput(x int, y int) string
	nextGenerationForCell(x, y int, currentGen Grid)
}

type grid struct {
	cells        map[int]map[int]bool
	sizeX, sizeY int
	gen          int
	livingCells  int
}

func CopyGrid(g Grid) Grid {
	sizeX, sizeY := g.getSize()
	newGrid := NewGrid(sizeX, sizeY)

	for x := 0; x < sizeX; x++ {
		for y := 0; y < sizeY; y++ {
			newGrid.setCell(x, y, g.getCell(x, y))
		}
	}

	return newGrid
}

func NewGrid(sizeX, sizeY int) Grid {
	g := &grid{
		cells: make(map[int]map[int]bool),
		sizeX: sizeX,
		sizeY: sizeY,
	}
	g.Init()

	return g
}

func (g *grid) Init() {
	for x := 0; x < g.sizeX; x++ {
		g.cells[x] = make(map[int]bool)

		for y := 0; y < g.sizeY; y++ {
			g.setCell(x, y, false)
		}
	}
}

func (g *grid) Populate(population int) {
	rand.Seed(time.Now().UnixNano())

	for population > 0 {
		x := rand.Intn(g.sizeX - 1)
		y := rand.Intn(g.sizeY - 1)
		if !g.getCell(x, y) {
			g.setCell(x, y, true)
			population--
		}
	}
}

func (g *grid) LivingCells() int {
	return g.livingCells
}

func (g *grid) NextGeneration() {
	currentGen := CopyGrid(g)

	g.livingCells = 0
	for x := 0; x < g.sizeX; x++ {
		for y := 0; y < g.sizeY; y++ {
			g.nextGenerationForCell(x, y, currentGen)
		}
	}

	g.gen++
}

func (g *grid) Show(out io.Writer) error {
	var output string

	output += fmt.Sprintf(headerFormat, g.gen, g.livingCells)

	for x := 0; x < g.sizeX; x++ {
		for y := 0; y < g.sizeY; y++ {
			output += g.getCellVisualOutput(x, y)
		}

		output += "\n"
	}

	if _, err := fmt.Fprintf(out, output); err != nil {
		return err
	}

	return nil
}

func (g *grid) getCell(x, y int) bool {
	return g.cells[x][y]
}

func (g *grid) setCell(x, y int, value bool) {
	g.cells[x][y] = value
	if value {
		g.livingCells++
	}
}

func (g *grid) getSize() (x, y int) {
	return g.sizeX, g.sizeY
}

func (g *grid) getCellVisualOutput(x int, y int) string {
	if g.getCell(x, y) {
		return "1"
	}

	return " "
}

func (g *grid) canCellLiveOn(x, y int) bool {
	aliveNeighbours := g.countLiveNeighbours(x, y)
	if g.getCell(x, y) &&
		(aliveNeighbours == 2 || aliveNeighbours == 3) {
		return true
	}

	if !g.getCell(x, y) && aliveNeighbours == 3 {
		return true
	}

	return false
}

func (g *grid) countLiveNeighbours(x, y int) (n int) {
	if g.cells[x-1][y-1] {
		n++
	}

	if g.cells[x-1][y] {
		n++
	}

	if g.cells[x-1][y+1] {
		n++
	}

	if g.cells[x][y-1] {
		n++
	}

	if g.cells[x][y+1] {
		n++
	}

	if g.cells[x+1][y-1] {
		n++
	}

	if g.cells[x+1][y] {
		n++
	}

	if g.cells[x+1][y+1] {
		n++
	}

	return
}

func (g *grid) nextGenerationForCell(x, y int, currentGen Grid) {
	canLive := currentGen.canCellLiveOn(x, y)
	g.setCell(x, y, canLive)
}
