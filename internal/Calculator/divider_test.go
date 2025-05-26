package calculator_test

import (
	"testing"

	calculator "github.com/lautaromdelgado/calculator-test/internal/Calculator"
	cuestom_errors "github.com/lautaromdelgado/calculator-test/internal/errors"
	"github.com/stretchr/testify/suite"
)

type DividerTestSuite struct {
	suite.Suite
	underTest calculator.IDivider
}

func TestDividerTestSuite(t *testing.T) {
	suite.Run(t, new(DividerTestSuite))
}

func (d *DividerTestSuite) SetupTest() {
	d.underTest = calculator.NewDivider()
}

func (d *DividerTestSuite) TestDivide() {
	result, err := d.underTest.Divide(10, 2)
	d.Nil(err)
	d.Equal(5.0, result)
}

func (d *DividerTestSuite) TestDivide_ByZero() {
	result, err := d.underTest.Divide(10, 0)
	d.Equal(cuestom_errors.ErrDivisionByZero, err)
	d.Equal(0.0, result)
}
