package app

import (
	"fmt"
	"strconv"
	"time"
)

const (
	invalidBalance = "Invalid balance"
)

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

func convertCopTuRub(cop string) string {
	copInt, err := strconv.ParseInt(cop, 10, 64)
	if err != nil {
		return invalidBalance
	}

	switch {
	case 0 <= copInt && copInt < 10:

		return fmt.Sprintf("0.0%d", copInt)

	case 10 <= copInt && copInt < 100:

		return fmt.Sprintf("0.%d", copInt)

	case 100 <= copInt:

		return fmt.Sprintf("%d.%s", copInt/100, cop[len(cop)-2:])

	default:

		return invalidBalance

	}
}
