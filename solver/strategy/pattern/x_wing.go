package pattern

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Implementation of the XY-Wing pattern.
type XWing struct{}

// Apply pattern on Sudoku.
func (p *XWing) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	pv := *possibleValuesRef

	// Count possible value occurrences for each row and column.
	rowValueOccurrences := make([]*map[int]int, model.SudokuSize)
	columnValueOccurrences := make([]*map[int]int, model.SudokuSize)
	for i := 0; i < model.SudokuSize; i++ {
		rowValueOccurrences[i] = countValueOccurrences(getRowPossibleValues(i, possibleValuesRef, false))
		columnValueOccurrences[i] = countValueOccurrences(getColumnPossibleValues(i, possibleValuesRef, false))
	}

	// Calculate all combinations of every two rows and every two columns.
	combinations := getCombinations(2, model.SudokuSize)
	combinationCount := len(combinations)

	// For all row and column combinations
	for rowComb := 0; rowComb < combinationCount; rowComb++ {
		rowCombination := combinations[rowComb]

		for columnComb := 0; columnComb < combinationCount; columnComb++ {
			columnCombination := combinations[columnComb]

			if cuttingCellLookups := p.findCuttingCells(possibleValuesRef, rowCombination, columnCombination); cuttingCellLookups != nil {
				occurrences := *countValueOccurrences(cuttingCellLookups)

				// Check if all 4 cell lookups have a possible value in common
				for value, count := range occurrences {
					if count == 4 {
						// Check if both rows or both columns do not contain more occurrences of the value
						rowHaveMore := p.unitsContainMoreOccurrencesOfValue(rowValueOccurrences, rowCombination, value, 2)
						columnHaveMore := p.unitsContainMoreOccurrencesOfValue(columnValueOccurrences, columnCombination, value, 2)

						if rowHaveMore != columnHaveMore {
							fmt.Printf("Can reduce value %d in rows %v and columns %v!\n", value, rowCombination, columnCombination)

							changed = false

							// Can reduce value now in either rows or columns based on which of rowHaveMore or columnHaveMore is true
							if rowHaveMore {
								for _, row := range rowCombination {
									for column := 0; column < model.SudokuSize; column++ {
										process := true
										for _, crossColumn := range columnCombination {
											if column == crossColumn {
												process = false
												break
											}
										}

										if process {
											lookupPtr := pv[row][column]

											if lookupPtr != nil {
												lookup := *lookupPtr

												possible, ok := lookup[value]
												if ok && possible {
													lookup[value] = false
													changed = true
												}
											}
										}
									}
								}
							} else {
								for _, column := range columnCombination {
									for row := 0; row < model.SudokuSize; row++ {
										process := true
										for _, crossRow := range rowCombination {
											if row == crossRow {
												process = false
												break
											}
										}

										if process {
											lookupPtr := pv[row][column]

											if lookupPtr != nil {
												lookup := *lookupPtr

												possible, ok := lookup[value]
												if ok && possible {
													lookup[value] = false
													changed = true
												}
											}
										}
									}
								}
							}

							return changed
						}
					}
				}
			}
		}
	}

	return false
}

// Find all cutting cells of the passed row and column index combinations.
// Will return nil if not all cells could be found.
func (p *XWing) findCuttingCells(possibleValuesRef *[][]*map[int]bool, rowCombination []int, columnCombination []int) []*map[int]bool {
	pv := *possibleValuesRef

	cellLookups := make([]*map[int]bool, 0, len(rowCombination)*len(columnCombination))
	for _, row := range rowCombination {
		for _, column := range columnCombination {
			cellLookup := pv[row][column]

			if cellLookup == nil {
				return nil
			}

			cellLookups = append(cellLookups, cellLookup)
		}
	}

	return cellLookups
}

// Check if the passed rows or columns (units) contain more occurrences (count) for the passed [value] in the passed combination.
func (p *XWing) unitsContainMoreOccurrencesOfValue(unitsValueOccurrences []*map[int]int, combination []int, value int, count int) bool {
	for i := 0; i < len(combination); i++ {
		unitOcc := *unitsValueOccurrences[combination[i]]

		occurrences, _ := unitOcc[value]
		if occurrences > count {
			return true
		}
	}

	return false
}
