package module

import (
	"strconv"
	"fmt"
	"github.com/addonrizky/valasConvertValidator/library"
	"github.com/addonrizky/valasConvertValidator/model"
	"github.com/addonrizky/valasConvertValidator/constant"
)

func ConvertValidate(
	debitAccount string, 
	creditAccount string, 
	debitAmount string, 
	creditAmount string, 
	debitCurrency string, 
	creditCurrency string, 
	buyRate string, 
	sellRate string,
) model.Validation {

	var creditAmountCalculated float64
	var debitAmountCalculated float64
	
	calculationCreditAmountValid := true
	calculationDebitAmountValid := true

	/** converting the debitAmount variable into a float using ParseFloat method */
	debitAmountFloat, err := strconv.ParseFloat(debitAmount, 64)
	if err != nil {
		return library.GetValidationResult(constant.RC_X0, constant.DESC_X0)
	}

	/** converting the debitAmount variable into a float using ParseFloat method */
	creditAmountFloat, err := strconv.ParseFloat(creditAmount, 64)
	if err != nil {
		return library.GetValidationResult(constant.RC_X1, constant.DESC_X1)
	}

	/** converting the debitAmount variable into a float using ParseFloat method */
	sellRateFloat, err := strconv.ParseFloat(sellRate, 64)
	if err != nil {
		return library.GetValidationResult(constant.RC_X2, constant.DESC_X2)
	}

	/** converting the debitAmount variable into a float using ParseFloat method */
	buyRateFloat, err := strconv.ParseFloat(buyRate, 64)
	if err != nil {
		return library.GetValidationResult(constant.RC_X3, constant.DESC_X3)
	}

	if debitAmountFloat <= 0{
		return library.GetValidationResult(constant.RC_V0, constant.DESC_V0)
	}

	if creditAmountFloat <= 0{
		return library.GetValidationResult(constant.RC_V1, constant.DESC_V1)
	}
	
	if creditCurrency == debitCurrency {
		return library.GetValidationResult(constant.RC_V2, constant.DESC_V2)
	}

	if debitAccount[4:6] == creditAccount[4:6] {
		return library.GetValidationResult(constant.RC_V3, constant.DESC_V3)
	}

	//IDR TO VALAS handler
	if debitCurrency == "IDR" {
		creditAmountCalculated = debitAmountFloat / sellRateFloat
		debitAmountCalculated = creditAmountFloat * sellRateFloat

		if fmt.Sprintf("%.2f", creditAmountCalculated) != fmt.Sprintf("%.2f", creditAmountFloat) {
			calculationCreditAmountValid = false
		}

		if fmt.Sprintf("%.2f", debitAmountCalculated) != fmt.Sprintf("%.2f", debitAmountFloat) {
			calculationDebitAmountValid = false
		}

		if !(calculationDebitAmountValid || calculationCreditAmountValid) {
			return library.GetValidationResult(constant.RC_V4, constant.DESC_V4)
		}
	}

	//VALAS TO IDR Handler
	if debitCurrency != "IDR" && creditCurrency == "IDR" {
		creditAmountCalculated = debitAmountFloat * buyRateFloat
		debitAmountCalculated = creditAmountFloat / buyRateFloat

		if fmt.Sprintf("%.2f", creditAmountCalculated) != fmt.Sprintf("%.2f", creditAmountFloat) {
			calculationCreditAmountValid = false
		}

		if fmt.Sprintf("%.2f", debitAmountCalculated) != fmt.Sprintf("%.2f", debitAmountFloat) {
			calculationDebitAmountValid = false
		}

		if !(calculationDebitAmountValid || calculationCreditAmountValid) {
			return library.GetValidationResult(constant.RC_V4, constant.DESC_V4)
		}
	}

	//VALAS TO VALAS Handler
	if debitCurrency != "IDR" && creditCurrency != "IDR" {
		creditAmountCalculated = debitAmountFloat * buyRateFloat / sellRateFloat
		debitAmountCalculated = creditAmountFloat / buyRateFloat * sellRateFloat

		if fmt.Sprintf("%.2f", creditAmountCalculated) != fmt.Sprintf("%.2f", creditAmountFloat) {
			calculationCreditAmountValid = false
		}

		if fmt.Sprintf("%.2f", debitAmountCalculated) != fmt.Sprintf("%.2f", debitAmountFloat) {
			calculationDebitAmountValid = false
		}

		if !(calculationDebitAmountValid || calculationCreditAmountValid) {
			return library.GetValidationResult(constant.RC_V4, constant.DESC_V4)
		}
	}

	return library.GetValidationResult("00", "konversi valas valid")
}
