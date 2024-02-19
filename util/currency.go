package util

// Constants for all supported currencies
const (
	USD = "USD"
	ZAR = "ZAR"
	BWP = "BWP"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, ZAR, BWP:
		return true
	}
	return false
}
