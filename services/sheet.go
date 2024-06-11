package services

import (
	"mandalart.com/schemas"
	"mandalart.com/utils"
)

func CreateSheetByOwnerId(sheetID string) []schemas.Cell {
	var cells []schemas.Cell

	utils.DB.Where("sheet_id = ?", sheetID).Find(&cells)

	return cells
} 