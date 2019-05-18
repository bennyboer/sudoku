package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Block-row-check pattern implementation.
type BlockRowCheck struct{}

// Apply pattern on Sudoku.
func (p *BlockRowCheck) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	changed = false

	// For each block
	for block := 0; block < model.SudokuSize; block++ {
		possibleBlockValues := getBlockPossibleValues(block, possibleValuesRef, true)
		blockValueOccurrences := countValueOccurrences(possibleBlockValues)

		// For each row
		startRow := block / model.BlockSize * model.BlockSize
		normalizedBlock := block % model.BlockSize
		for row := 0; row < model.BlockSize; row++ {
			possibleRowValues := getRowPossibleValues(row+startRow, possibleValuesRef, true)
			rowLookupsInBlock := possibleRowValues[normalizedBlock*model.BlockSize : normalizedBlock*model.BlockSize+model.BlockSize]

			if didChange := p.findPatternAndUpdate(blockValueOccurrences, possibleBlockValues, possibleRowValues, rowLookupsInBlock); didChange {
				changed = true
			}
		}

		// Recalculate occurrences, since they could have changed
		blockValueOccurrences = countValueOccurrences(possibleBlockValues)

		// For each column
		startColumn := (block % model.BlockSize) * model.BlockSize
		blockOffset := block / model.BlockSize
		for column := 0; column < model.BlockSize; column++ {
			possibleColumnValues := getColumnPossibleValues(column+startColumn, possibleValuesRef, true)
			columnLookupsInBlock := possibleColumnValues[blockOffset*model.BlockSize : blockOffset*model.BlockSize+model.BlockSize]

			if didChange := p.findPatternAndUpdate(blockValueOccurrences, possibleBlockValues, possibleColumnValues, columnLookupsInBlock); didChange {
				changed = true
			}
		}
	}

	return changed
}

// Find the block-row-check pattern and update the possible value lookups.
func (p *BlockRowCheck) findPatternAndUpdate(blockValueOccurrences *map[int]int,
	blockUnit []*map[int]bool,
	rowUnit []*map[int]bool,
	rowLookupsInBlock []*map[int]bool) bool {
	changed := false

	// Count values in overlapping cells.
	crossOcc := countValueOccurrences(rowLookupsInBlock)

	for value, count := range *crossOcc {
		totalOccurrences := (*blockValueOccurrences)[value]

		if totalOccurrences-count == 0 {
			// All occurrences of [value] are only in the row (within the block)!
			// -> Remove all other occurrences of [value] in the row.
			for _, rowLookupPtr := range rowUnit {
				if rowLookupPtr == nil {
					continue
				}

				// Check if lookup ptr not overlapping with the block.
				overlapping := false
				for _, overlappingPtr := range rowLookupsInBlock {
					if rowLookupPtr == overlappingPtr {
						overlapping = true
						break
					}
				}

				if !overlapping {
					lookup := *rowLookupPtr

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
