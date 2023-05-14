package api

import (
	"github.com/Chien179/MusicPlayerBE/util"
	"github.com/go-playground/validator/v10"
)

var validDirection validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsValidDirection(currency)
	}

	return false
}
