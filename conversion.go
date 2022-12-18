package database

import (
	"fmt"

	"github.com/iVitaliya/logger-go"
)

func valueToInt(key string, value any) int {
	i, err := value.(int)
	if !err {
		logger.Error(`The value for the key "` + key + `" couldn't be transformed into a number. Original Value: ` + fmt.Sprint(value))
		return -1
	}

	return i
}
