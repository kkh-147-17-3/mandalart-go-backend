//go:build exclude
package services

import (
	"mandalart.com/utils"
	"mandalart.com/schemas"
)

func GetMainBySheetID(sheetID string) []schemas.Cell {
	var cells []schemas.Cell

	utils.DB.Where("sheet_id = ? AND depth = 1", sheetID).Find(&cells)

	return cells
} 