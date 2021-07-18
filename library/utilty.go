package library 

import (
	"github.com/addonrizky/valasConvertValidator/model"
)

//GetValidationResult, retrieve object validation
func GetValidationResult(code string, desc string) model.Validation{
	validationObject := model.Validation{
		Code: code,
		Desc: desc,
	}
	return validationObject
}