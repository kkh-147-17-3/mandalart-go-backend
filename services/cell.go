package services

import (
	"mandalart.com/utils"
	"mandalart.com/schemas"
)

func GetMainBySheetID(sheetID string) []schemas.Cell {
	var cells []schemas.Cell

	utils.DB.Where("sheetID = ?", sheetID).Find(&cells)

	return cells
} 