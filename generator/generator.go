package generator

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"math/rand"
	"time"
)

// Interface for Sudoku generators.
type SudokuGenerator interface {
	// Generates a Sudoku with the passed difficulty of range [0.0; 1.0],
	// where 0.0 is as easy as possible and 1.0 as difficult as possible.
	// Specify a timeout after which the generator should complete.
	Generate(difficulty float64, timeout time.Duration) (*model.Sudoku, error)
}

type SudokuGeneratorSimple struct {
	sudoku      *model.Sudoku
	isCancelled bool
}

type SudokuGeneratorDifficulty struct {
	difficulty  float64
	sudokuSrc   *[][]int
	isCancelled bool
}

func NewBacktrackingGenerator() *SudokuGeneratorSimple {
	rand.Seed(time.Now().UnixNano())

	return &SudokuGeneratorSimple{}
}

func NewDifficultyGenerator() *SudokuGeneratorDifficulty {
	rand.Seed(time.Now().UnixNano())

	return &SudokuGeneratorDifficulty{0, nil, false}
}

func AllGenerationAlgorithms() *map[string]SudokuGenerator {
	gMap := make(map[string]SudokuGenerator)

	gMap["difficulty"] = NewDifficultyGenerator()
	gMap["simple"] = NewBacktrackingGenerator()

	return &gMap
}
