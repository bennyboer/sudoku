package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Row-block-check pattern implementation.
type RowBlockCheck struct{}

// Apply pattern on Sudoku.
func (p *RowBlockCheck) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	changed = false

	// For each row
	for row := 0; row < model.SudokuSize; row++ {
		possibleRowValues := getRowPossibleValues(row, possibleValuesRef, true)
		rowValueOccurrences := countValueOccurrences(possibleRowValues)

		// For each block
		startBlock := row / model.BlockSize * model.BlockSize
		for block := 0; block < model.BlockSize; block++ {
			possibleBlockValues := getBlockPossibleValues(block+startBlock, possibleValuesRef, true)
			rowLookupsInBlock := possibleRowValues[block*model.BlockSize : block*model.BlockSize+model.BlockSize]

			if didChange := p.findPatternAndUpdate(rowValueOccurrences, possibleBlockValues, rowLookupsInBlock); didChange {
				changed = true
			}
		}
	}

	// For each column
	for column := 0; column < model.SudokuSize; column++ {
		possibleColumnValues := getColumnPossibleValues(column, possibleValuesRef, true)
		columnValueOccurrences := countValueOccurrences(possibleColumnValues)

		// For each block
		blockOffset := column / model.BlockSize
		for block := 0; block < model.BlockSize; block++ {
			possibleBlockValues := getBlockPossibleValues(block*model.BlockSize+blockOffset, possibleValuesRef, true)
			columnLookupsInBlock := possibleColumnValues[block*model.BlockSize : block*model.BlockSize+model.BlockSize]

			if didChange := p.findPatternAndUpdate(columnValueOccurrences, possibleBlockValues, columnLookupsInBlock); didChange {
				changed = true
			}
		}
	}

	return
}

// Find the row-block-check pattern and update the possible value lookups.
func (p *RowBlockCheck) findPatternAndUpdate(rowValueOccurrences *map[int]int,
	blockUnit []*map[int]bool,
	rowLookupsInBlock []*map[int]bool) bool {
	changed := false

	// Count values in overlapping cells.
	crossOcc := countValueOccurrences(rowLookupsInBlock)

	for value, count := range *crossOcc {
		totalOccurrences := (*rowValueOccurrences)[value]

		if totalOccurrences-count == 0 {
			// All occurrences of [value] are only in the block! -> Remove all other occurrences of [value] in block.
			for i := 0; i < len(blockUnit); i++ {
				blockLookupPtr := blockUnit[i]

				if blockLookupPtr == nil {
					continue
				}

				// Check if lookup ptr not overlapping with the row.
				overlapping := false
				for _, rowLookupPtr := range rowLookupsInBlock {
					if blockLookupPtr == rowLookupPtr {
						overlapping = true
						break
					}
				}

				if !overlapping {
					lookup := *blockLookupPtr

					possible, ok := lookup[value]
					if ok && possible {
						lookup[value] = false
						changed = true
					}
				}
			}
		}
	}

	return changed
}
