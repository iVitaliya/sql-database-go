package database

import (
	"github.com/iVitaliya/colors-go"
	"github.com/iVitaliya/logger-go"
)

func logSetID(table string, id int64) {
	logger.Info("Table `"+table+"` has inserted row with ID:", id)
}

func logError(errored_on string, message string) {
	logger.Error(colors.BrightRed("Database Error"), colors.BrightBlack(":"), colors.BrightRed(errored_on), colors.BrightBlack(">>"), message)
}

func logAffectedRows(table string, affected_rows int64) {
	logger.Info("Number of rowsaffected in table`"+table+"`:", affected_rows)
}
