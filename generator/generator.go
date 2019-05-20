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
	Generate(difficulty float64) (*model.Sudoku, error)
}

type SudokuGeneratorBacktracking struct {
}

type SudokuGeneratorDifficulty struct {
}

func NewBacktrackingGenerator() *SudokuGeneratorBacktracking {
	rand.Seed(time.Now().UnixNano())

	return &SudokuGeneratorBacktracking{}
}

func NewDifficultyGenerator() *SudokuGeneratorDifficulty {
	rand.Seed(time.Now().UnixNano())

	return &SudokuGeneratorDifficulty{}
}

func AllGenerationAlgorithms() *map[string]SudokuGenerator {
	gMap := make(map[string]SudokuGenerator)

	gMap["difficulty"] = NewDifficultyGenerator()
	gMap["simple"] = NewBacktrackingGenerator()

	return &gMap
}
