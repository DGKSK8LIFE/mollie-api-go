package mollie

import (
	"errors"
	"strings"
)

const (
	//CentesimalBase is a helper to transform a cents integer into
	//a float
	CentesimalBase = 100

	//DefaultCurrency is used when an empty currency code is provided
	DefaultCurrency = "EUR"
)

var (
	errInvalidISO4217Code = errors.New("currency code with invalid lenght")
)

//Currency is a monetary value representation
type Currency struct {
	Value        string `json:"value"`
	CurrencyCode string `json:"currency"`
}

//NewCurrency builds a currency struct.
func NewCurrency(code string, val string) (Currency, error) {
	if strings.TrimSpace(code) == "" {
		code = DefaultCurrency
	}

	if len(code) != 3 {
		return Currency{}, errInvalidISO4217Code
	}

	return Currency{
		CurrencyCode: code,
		Value:        val,
	}, nil
}
