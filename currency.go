package main

import "strconv"

var currencySymbols = map[string]string{
	"USD": "$",
	"EUR": "€",
	"GBP": "£",
	"JPY": "¥",
	"CNY": "¥",
	"INR": "₹",
	"RUB": "₽",
	"KRW": "₩",
	"BRL": "R$",
	"SGD": "SGD$",
}

var currencyFractionDigits = map[string]int{
	"JPY": 0,
}

func formatCurrency(currency string, amount float64) string {
	digits := 2
	if value, ok := currencyFractionDigits[currency]; ok {
		digits = value
	}

	formatted := strconv.FormatFloat(amount, 'f', digits, 64)
	if currency == "JPY" {
		return formatted + " 円"
	}

	return currencySymbols[currency] + formatted
}
