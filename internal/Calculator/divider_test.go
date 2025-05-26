package calculator_test

import (
	"testing"

	calculator "github.com/lautaromdelgado/calculator-test/internal/Calculator"
	cuestom_errors "github.com/lautaromdelgado/calculator-test/internal/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DividerTestSuite struct {
	suite.Suite
	underTest calculator.IDivider
}

func TestDivide(t *testing.T) {
	divider := calculator.NewDivider()

	result, err := divider.Divide(10, 2)
	assert.Nil(t, err)
	assert.Equal(t, 5.0, result)
}

func TestDivide_ByZero(t *testing.T) {
	divider := calculator.NewDivider()

	result, err := divider.Divide(10, 0)
	assert.Equal(t, cuestom_errors.ErrDivisionByZero, err)
	assert.Equal(t, 0.0, result)
}
