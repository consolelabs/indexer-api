package utils

import (
	"errors"

	"gorm.io/gorm"
)

func IsRecordNotFound(err error) bool {
	return err != nil && (errors.Is(err, gorm.ErrRecordNotFound) || err.Error() == "sql: no rows in result set")
}
