package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Row-block-check pattern implementation.
type RowBlockCheck struct{}

// Apply pattern on Sudoku.
func (p *RowBlockCheck) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	// For each row
	for row := 0; row < model.SudokuSize; row++ {
		possibleRowValues := getRowPossibleValues(row, possibleValuesRef, true)
		rowValueOccurrences := countValueOccurrences(possibleRowValues)

		// For each block
		startBlock := row / model.BlockSize * model.BlockSize
		for block := 0; block < model.BlockSize; block++ {
			possibleBlockValues := getBlockPossibleValues(block+startBlock, possibleValuesRef, true)
			rowLookupsInBlock := possibleRowValues[block*model.BlockSize : block*model.BlockSize+model.BlockSize]

			if changed := p.findPatternAndUpdate(rowValueOccurrences, possibleRowValues, possibleBlockValues, rowLookupsInBlock); changed {
				return true
			}
		}
	}

	return false
}

// Find the row-block-check pattern and update the possible value lookups.
func (p *RowBlockCheck) findPatternAndUpdate(rowValueOccurrences *map[int]int,
	rowUnit []*map[int]bool,
	blockUnit []*map[int]bool,
	rowLookupsInBlock []*map[int]bool) bool {
	changed := false

	// Count values in overlapping cells.
	crossOcc := countValueOccurrences(rowLookupsInBlock)

	for value, count := range *crossOcc {
		totalOccurrences, _ := (*rowValueOccurrences)[value]

		if totalOccurrences-count == 0 {
			// All occurrences of [value] are only in the block! -> Remove all other occurrences of [value] in block.
			for _, blockLookupPtr := range blockUnit {
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
