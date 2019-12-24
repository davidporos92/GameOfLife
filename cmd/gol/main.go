package main

import (
	"github.com/davidporos92/GameOfLife/internal/grid"
	"github.com/gosuri/uilive"
	"log"
	"time"
)

const GridSizeY = 100
const GridSizeX = 70
const StartingPopulation = 1000

var generations = 1000

func main() {
	o := uilive.New()
	o.Start()

	g := grid.NewGrid(GridSizeX, GridSizeY)
	g.Populate(StartingPopulation)

	if err := g.Show(o); err != nil {
		log.Fatal(err)
	}

	for generations > 0 {
		time.Sleep(100 * time.Millisecond)

		g.NextGeneration()
		generations--

		if err := g.Show(o); err != nil {
			log.Fatal(err)
		}

		if g.LivingCells() <= 0 {
			return
		}
	}

	o.Stop()
}
