package calculator

import errors "github.com/lautaromdelgado/calculator-test/internal/errors"

type IDivider interface {
	Divide(a, b float64) (float64, error)
}

type divider struct{}

func NewDivider() IDivider {
	return &divider{}
}

func (d *divider) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.ErrDivisionByZero
	}
	return a / b, nil
}
