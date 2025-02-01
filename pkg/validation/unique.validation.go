package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func (v *Validate) isUnique(fl validator.FieldLevel) bool {
	// Extract parameters (e.g., "tablename,columnName")
	fmt.Println("fl", fl.Param())
	params := strings.Split(fl.Param(), ";")
	fmt.Println("params", params)
	if len(params) != 2 {
		return false // Invalid parameters
	}

	tableName := params[0]
	columnName := params[1]
	value := fl.Field().String() // Value to validate

	// Perform the database query
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = $1", tableName, columnName)
	v.log.Info("query", zap.String("query", query), zap.String("value", value), zap.String("tableName", tableName), zap.String("columnName", columnName))
	if err := v.db.Get(&count, query, value); err != nil {
		v.log.Error("Database error during uniqueness validation", zap.Error(err))
		return false // Return false on database errors
	}

	return count == 0 // Return true if the value is unique
}
