package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/kahakai/simple-bank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	currency, ok := fieldLevel.Field().Interface().(string)
	if !ok {
		return false
	}

	return util.IsSupportedCurrency(currency)
}
