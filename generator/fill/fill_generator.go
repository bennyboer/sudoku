package fill

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"math/rand"
	"time"
)

// Generator randomly filling a newly created Sudoku.
type FillGenerator struct{}

// Will fill a Sudoku randomly from scratch.
func (g *FillGenerator) Generate(difficulty float32) *model.Sudoku {
	//sudoku := model.EmptySudoku()
	// TODO Implement
	return nil
}

// Retrieve a random integer between zero and the passed max integer.
func randomInt(max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max)
}
