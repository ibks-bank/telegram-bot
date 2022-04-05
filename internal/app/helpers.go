package app

import "time"

func formatDate(date time.Time) string {
	return date.Format("02.01.2006")
}

func formatCurrency(cur string) string {
	switch cur {
	case "rub":
		return "₽"
	case "eur":
		return "€"
	case "usd":
		return "$"
	default:
		return cur
	}
}
