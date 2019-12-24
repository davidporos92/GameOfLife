package grid

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type NextGenerationForLivingCellTestSuite struct {
	suite.Suite
	grid Grid
	investigatedX, investigatedY int
}

func (s *NextGenerationForLivingCellTestSuite) getInvestigatedCell() bool {
	return s.grid.getCell(s.investigatedX, s.investigatedY)
}

func (s *NextGenerationForLivingCellTestSuite) nextGenerationForInvestigatedCell() {
	currentGen := CopyGrid(s.grid)
	s.grid.nextGenerationForCell(s.investigatedX, s.investigatedY, currentGen)
}

func (s *NextGenerationForLivingCellTestSuite) SetupSuite() {
	s.investigatedX, s.investigatedY = 1, 1
}

func (s *NextGenerationForLivingCellTestSuite) SetupTest() {
	s.grid = NewGrid(3, 3)
	s.grid.setCell(s.investigatedX, s.investigatedY, true)
}

func (s *NextGenerationForLivingCellTestSuite) TestWithoutLivingNeighbours() {
	s.nextGenerationForInvestigatedCell()
	s.Equal(false, s.getInvestigatedCell())
}

func (s *NextGenerationForLivingCellTestSuite) TestWithOneLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(false, s.getInvestigatedCell())
}

func (s *NextGenerationForLivingCellTestSuite) TestWithTwoLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(true, s.getInvestigatedCell())
}

func (s *NextGenerationForLivingCellTestSuite) TestWithThreeLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(0, 1, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(true, s.getInvestigatedCell())
}

func (s *NextGenerationForLivingCellTestSuite) TestWithFourLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(0, 1, true)
	s.grid.setCell(0, 2, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(false, s.getInvestigatedCell())
}

func (s *NextGenerationForLivingCellTestSuite) TestWithFiveLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(0, 1, true)
	s.grid.setCell(0, 2, true)
	s.grid.setCell(1, 0, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(false, s.getInvestigatedCell())
}

func (s *NextGenerationForLivingCellTestSuite) TestWithSixLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(0, 1, true)
	s.grid.setCell(0, 2, true)
	s.grid.setCell(1, 0, true)
	s.grid.setCell(1, 2, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(false, s.getInvestigatedCell())
}

func (s *NextGenerationForLivingCellTestSuite) TestWithSevenLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(0, 1, true)
	s.grid.setCell(0, 2, true)
	s.grid.setCell(1, 0, true)
	s.grid.setCell(1, 2, true)
	s.grid.setCell(2, 0, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(false, s.getInvestigatedCell())
}

func (s *NextGenerationForLivingCellTestSuite) TestWithEightLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(0, 1, true)
	s.grid.setCell(0, 2, true)
	s.grid.setCell(1, 0, true)
	s.grid.setCell(1, 2, true)
	s.grid.setCell(2, 0, true)
	s.grid.setCell(2, 1, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(false, s.getInvestigatedCell())
}

type NextGenerationForNonLivingCellTestSuite struct {
	suite.Suite
	grid Grid
	investigatedX, investigatedY int
}

func (s *NextGenerationForNonLivingCellTestSuite) getInvestigatedCell() bool {
	return s.grid.getCell(s.investigatedX, s.investigatedY)
}

func (s *NextGenerationForNonLivingCellTestSuite) nextGenerationForInvestigatedCell() {
	currentGen := CopyGrid(s.grid)
	s.grid.nextGenerationForCell(s.investigatedX, s.investigatedY, currentGen)
}

func (s *NextGenerationForNonLivingCellTestSuite) SetupSuite() {
	s.investigatedX, s.investigatedY = 1, 1
}

func (s *NextGenerationForNonLivingCellTestSuite) SetupTest() {
	s.grid = NewGrid(3, 3)
}

func (s *NextGenerationForNonLivingCellTestSuite) TestWithoutLivingNeighbours() {
	s.nextGenerationForInvestigatedCell()
	s.Equal(false, s.getInvestigatedCell())
}

func (s *NextGenerationForNonLivingCellTestSuite) TestWithOneLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(false, s.getInvestigatedCell())
}

func (s *NextGenerationForNonLivingCellTestSuite) TestWithTwoLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(false, s.getInvestigatedCell())
}

func (s *NextGenerationForNonLivingCellTestSuite) TestWithThreeLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(0, 1, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(true, s.getInvestigatedCell())
}

func (s *NextGenerationForNonLivingCellTestSuite) TestWithFourLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(0, 1, true)
	s.grid.setCell(0, 2, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(false, s.getInvestigatedCell())
}

func (s *NextGenerationForNonLivingCellTestSuite) TestWithFiveLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(0, 1, true)
	s.grid.setCell(0, 2, true)
	s.grid.setCell(1, 0, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(false, s.getInvestigatedCell())
}

func (s *NextGenerationForNonLivingCellTestSuite) TestWithSixLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(0, 1, true)
	s.grid.setCell(0, 2, true)
	s.grid.setCell(1, 0, true)
	s.grid.setCell(1, 2, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(false, s.getInvestigatedCell())
}

func (s *NextGenerationForNonLivingCellTestSuite) TestWithSevenLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(0, 1, true)
	s.grid.setCell(0, 2, true)
	s.grid.setCell(1, 0, true)
	s.grid.setCell(1, 2, true)
	s.grid.setCell(2, 0, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(false, s.getInvestigatedCell())
}

func (s *NextGenerationForNonLivingCellTestSuite) TestWithEightLivingNeighbours() {
	s.grid.setCell(0, 0, true)
	s.grid.setCell(0, 1, true)
	s.grid.setCell(0, 2, true)
	s.grid.setCell(1, 0, true)
	s.grid.setCell(1, 2, true)
	s.grid.setCell(2, 0, true)
	s.grid.setCell(2, 1, true)
	s.grid.setCell(2, 2, true)
	s.nextGenerationForInvestigatedCell()

	s.Equal(false, s.getInvestigatedCell())
}

func TestCopyGrid(t *testing.T) {
	grid := NewGrid(100, 100)
	grid.Populate(100)
	copiedGrid := CopyGrid(grid)

	gridBuffer := new(bytes.Buffer)
	copiedGridBuffer := new(bytes.Buffer)

	err := grid.Show(gridBuffer)
	assert.NoError(t, err)
	err = copiedGrid.Show(copiedGridBuffer)
	assert.NoError(t, err)

	assert.Equal(t, gridBuffer.String(), copiedGridBuffer.String())
}

func TestGrid_SetGetCell(t *testing.T) {
	grid := NewGrid(1, 1)

	assert.Equal(t, false, grid.getCell(0, 0))
	assert.Equal(t, 0, grid.LivingCells())

	grid.setCell(0, 0, true)
	assert.Equal(t, true, grid.getCell(0, 0))
	assert.Equal(t, 1, grid.LivingCells())

	grid.setCell(0, 0, false)
	assert.Equal(t, false, grid.getCell(0, 0))
}

func TestGrid_GetSize(t *testing.T) {
	expectedSizeX, expectedSizeY := 5, 9
	grid := NewGrid(expectedSizeX, expectedSizeY)
	sizeX, sizeY := grid.getSize()

	assert.Equal(t, expectedSizeX, sizeX)
	assert.Equal(t, expectedSizeY, sizeY)
}

func TestGrid_GetCellVisualOutput(t *testing.T) {
	grid := NewGrid(1, 2)
	grid.setCell(0, 0, true)

	assert.Equal(t, "1", grid.getCellVisualOutput(0, 0))
	assert.Equal(t, " ", grid.getCellVisualOutput(0, 1))
}

func TestGrid_CountLiveNeighbours_AndCanCellLiveOn(t *testing.T) {
	investigatedX, investigatedY := 1, 1
	grid := NewGrid(3, 3)
	grid.setCell(investigatedX, investigatedY, true)

	assert.Equal(t, 0, grid.countLiveNeighbours(investigatedX, investigatedY))
	assert.Equal(t, false, grid.canCellLiveOn(investigatedX, investigatedY))

	grid.setCell(0, 0, true)
	assert.Equal(t, 1, grid.countLiveNeighbours(investigatedX, investigatedY))
	assert.Equal(t, false, grid.canCellLiveOn(investigatedX, investigatedY))

	grid.setCell(0, 1, true)
	assert.Equal(t, 2, grid.countLiveNeighbours(investigatedX, investigatedY))
	assert.Equal(t, true, grid.canCellLiveOn(investigatedX, investigatedY))

	grid.setCell(0, 2, true)
	assert.Equal(t, 3, grid.countLiveNeighbours(investigatedX, investigatedY))
	assert.Equal(t, true, grid.canCellLiveOn(investigatedX, investigatedY))

	// Reproduction test
	grid.setCell(investigatedX, investigatedY, false)
	assert.Equal(t, true, grid.canCellLiveOn(investigatedX, investigatedY))
	grid.setCell(investigatedX, investigatedY, true)

	grid.setCell(1, 0, true)
	assert.Equal(t, 4, grid.countLiveNeighbours(investigatedX, investigatedY))
	assert.Equal(t, false, grid.canCellLiveOn(investigatedX, investigatedY))

	grid.setCell(1, 2, true)
	assert.Equal(t, 5, grid.countLiveNeighbours(investigatedX, investigatedY))
	assert.Equal(t, false, grid.canCellLiveOn(investigatedX, investigatedY))

	grid.setCell(2, 0, true)
	assert.Equal(t, 6, grid.countLiveNeighbours(investigatedX, investigatedY))
	assert.Equal(t, false, grid.canCellLiveOn(investigatedX, investigatedY))

	grid.setCell(2, 1, true)
	assert.Equal(t, 7, grid.countLiveNeighbours(investigatedX, investigatedY))
	assert.Equal(t, false, grid.canCellLiveOn(investigatedX, investigatedY))

	grid.setCell(2, 2, true)
	assert.Equal(t, 8, grid.countLiveNeighbours(investigatedX, investigatedY))
	assert.Equal(t, false, grid.canCellLiveOn(investigatedX, investigatedY))
}

func TestGrid_Populate(t *testing.T) {
	expectedPopulation := 100
	grid := NewGrid(100, 100)

	assert.Equal(t, 0, grid.LivingCells())

	grid.Populate(expectedPopulation)
	assert.Equal(t, expectedPopulation, grid.LivingCells())
}

func TestGrid_NextGenerationForCell(t *testing.T) {
	suite.Run(t, new(NextGenerationForLivingCellTestSuite))
	suite.Run(t, new(NextGenerationForNonLivingCellTestSuite))
}

func TestGrid_NextGeneration(t *testing.T) {
	grid := NewGrid(3, 3)

	grid.setCell(0, 0, true)
	grid.setCell(1, 0, true)
	grid.setCell(0, 1, true)
	grid.setCell(2, 2, true)
	grid.NextGeneration()

	// Generation 1
	assert.Equal(t, true, grid.getCell(0, 0))
	assert.Equal(t, true, grid.getCell(0, 1))
	assert.Equal(t, false, grid.getCell(0, 2))
	assert.Equal(t, true, grid.getCell(1, 0))
	assert.Equal(t, false, grid.getCell(1, 1))
	assert.Equal(t, false, grid.getCell(1, 2))
	assert.Equal(t, false, grid.getCell(2, 0))
	assert.Equal(t, false, grid.getCell(2, 1))
	assert.Equal(t, false, grid.getCell(2, 2))

	grid.NextGeneration()

	// Generation 2
	assert.Equal(t, true, grid.getCell(0, 0))
	assert.Equal(t, true, grid.getCell(0, 1))
	assert.Equal(t, false, grid.getCell(0, 2))
	assert.Equal(t, true, grid.getCell(1, 0))
	assert.Equal(t, true, grid.getCell(1, 1))
	assert.Equal(t, false, grid.getCell(1, 2))
	assert.Equal(t, false, grid.getCell(2, 0))
	assert.Equal(t, false, grid.getCell(2, 1))
	assert.Equal(t, false, grid.getCell(2, 2))

	grid.NextGeneration()

	// Generation 3
	assert.Equal(t, true, grid.getCell(0, 0))
	assert.Equal(t, true, grid.getCell(0, 1))
	assert.Equal(t, false, grid.getCell(0, 2))
	assert.Equal(t, true, grid.getCell(1, 0))
	assert.Equal(t, true, grid.getCell(1, 1))
	assert.Equal(t, false, grid.getCell(1, 2))
	assert.Equal(t, false, grid.getCell(2, 0))
	assert.Equal(t, false, grid.getCell(2, 1))
	assert.Equal(t, false, grid.getCell(2, 2))
}

func TestGrid_Show(t *testing.T) {
	expectedOutput := fmt.Sprintf(headerFormat, 0, 2)
	expectedOutput += "1 \n 1\n"

	grid := NewGrid(2, 2)
	grid.setCell(0, 0, true)
	grid.setCell(1, 1, true)

	gridBuffer := new(bytes.Buffer)
	err := grid.Show(gridBuffer)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, gridBuffer.String())
}
