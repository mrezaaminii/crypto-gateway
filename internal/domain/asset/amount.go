package asset

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

type Amount struct {
	value *big.Int
}

var ErrInvalidFormat = errors.New("Invalid Amount provided!")

func NewAmount(v *big.Int) Amount {
	return Amount{value: v}
}

func ParseAmount(s string, decimals int) (Amount, error) {
	s = strings.TrimSpace(s)

	if s == "" {
		return Amount{}, ErrInvalidFormat
	}

	whole, frac, _ := strings.Cut(s, ".")
	if whole == "" {
		whole = "0"
	}

	if !isDigits(whole) || (frac != "" && !isDigits(frac)) {
		return Amount{}, fmt.Errorf("%w: %q", ErrInvalidFormat, s)
	}

	if len(frac) > decimals {
		return Amount{}, fmt.Errorf("%w: Too many decimal (max is %d)", ErrInvalidFormat, decimals)
	}

	frac += strings.Repeat("0", decimals - len(frac))

	value, ok := new(big.Int).SetString(whole + frac, 10)

	if !ok {
		return Amount{}, fmt.Errorf("%w: %q", ErrInvalidFormat, s) 
	}

	return Amount{value: value}, nil
}

func (a *Amount) String(decimals int) string {
	if a.value == nil {
		return "0"
	}

	digits := a.value.String()

	if len(digits) <= decimals {
		digits = strings.Repeat("0", decimals - len(digits)+1) + digits
	}

	split := len(digits) - decimals
	whole := digits[:split]
	frac := strings.TrimRight(digits[split:], "0")

	if frac == "" {
		return whole
	}

	return whole + "." + frac

}

func isDigits(s string) bool {
	if s == "" {
		return false
	}

	for _, digit := range s {
		if digit < '0' || digit > '9' {
			return false
		}
	}

	return true
}