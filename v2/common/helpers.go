package common

import (
	"bytes"
	"math"
)

// RoundPriceToTickSize rounds the price to the nearest tick size
func RoundPriceToTickSize(price, tickSize float64) float64 {
	if tickSize == 0 {
		return price // To avoid division by zero
	}
	// Calculate the factor by which to multiply and divide the price to conform to the tick size.
	factor := 1 / tickSize

	// Round the price to the nearest tick size.
	roundedPrice := math.Round(price*factor) / factor

	return roundedPrice
}

// ToJSONList convert v to json list if v is a map
func ToJSONList(v []byte) []byte {
	if len(v) > 0 && v[0] == '{' {
		var b bytes.Buffer
		b.Write([]byte("["))
		b.Write(v)
		b.Write([]byte("]"))
		return b.Bytes()
	}
	return v
}
