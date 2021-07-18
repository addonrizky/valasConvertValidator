package main

import (
	"fmt"
	"github.com/addonrizky/valasConvertValidator/module"
)

func main() {
	//debitAccount ,creditAccount debitAmount,creditAmount,debitCurrency,creditCurrency,buyRate,sellRate
	cek := module.ConvertValidate("020601087063504","020602087064504","406.04","101","SAR","USD","3891.60","15645.00")
	fmt.Println(cek)
}
