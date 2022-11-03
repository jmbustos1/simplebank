package util

// Suported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

//IsSupportedCurrency returns a true if the currency is supported

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
