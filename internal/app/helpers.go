package app

import "time"

func formatDate(date time.Time) string {
	return date.Format("02.01.2006")
}

func formatCurrency(cur string) string {
	switch cur {
	case "CURRENCY_RUB":
		return "₽"
	case "CURRENCY_EURO":
		return "€"
	case "CURRENCY_DOLLAR_US":
		return "$"
	default:
		return cur
	}
}
