package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"capital-gains/internal/domain"
)

func TestCalculateTaxes(t *testing.T) {
	t.Run("Case #1", func(t *testing.T) {
		t.Parallel()

		operations := domain.Operations{
			{
				Type:     domain.BUY,
				Quantity: 100,
				UnitCost: 10.00,
			},
			{
				Type:     domain.SELL,
				Quantity: 50,
				UnitCost: 15.00,
			},
			{
				Type:     domain.SELL,
				Quantity: 50,
				UnitCost: 15.00,
			},
		}

		expectedTaxes := domain.Taxes{
			{Tax: 0.00},
			{Tax: 0.00},
			{Tax: 0.00},
		}

		portfolio := domain.Portfolio{}
		taxes := portfolio.CalculateTaxes(operations)
		assert.Equal(t, expectedTaxes, taxes)
	})

	t.Run("Case #2", func(t *testing.T) {
		t.Parallel()

		operations := domain.Operations{
			{
				Type:     domain.BUY,
				Quantity: 10000,
				UnitCost: 10.00,
			},
			{
				Type:     domain.SELL,
				Quantity: 5000,
				UnitCost: 20.00,
			},
			{
				Type:     domain.SELL,
				Quantity: 5000,
				UnitCost: 5.00,
			},
		}

		expectedTaxes := domain.Taxes{
			{Tax: 0.00},
			{Tax: 10000.00},
			{Tax: 0.00},
		}

		portfolio := domain.Portfolio{}
		taxes := portfolio.CalculateTaxes(operations)
		assert.Equal(t, expectedTaxes, taxes)
	})

	t.Run("Case #3", func(t *testing.T) {
		t.Parallel()

		operations := domain.Operations{
			{
				Type:     domain.BUY,
				Quantity: 10000,
				UnitCost: 10.00,
			},
			{
				Type:     domain.SELL,
				Quantity: 5000,
				UnitCost: 5.00,
			},
			{
				Type:     domain.SELL,
				Quantity: 3000,
				UnitCost: 20.00,
			},
		}

		expectedTaxes := domain.Taxes{
			{Tax: 0.00},
			{Tax: 0.00},
			{Tax: 1000.00},
		}

		portfolio := domain.Portfolio{}
		taxes := portfolio.CalculateTaxes(operations)
		assert.Equal(t, expectedTaxes, taxes)
	})

	t.Run("Case #4", func(t *testing.T) {
		t.Parallel()

		operations := domain.Operations{
			{
				Type:     domain.BUY,
				Quantity: 10000,
				UnitCost: 10.00,
			},
			{
				Type:     domain.BUY,
				Quantity: 5000,
				UnitCost: 25.00,
			},
			{
				Type:     domain.SELL,
				Quantity: 10000,
				UnitCost: 15.00,
			},
		}

		expectedTaxes := domain.Taxes{
			{Tax: 0.00},
			{Tax: 0.00},
			{Tax: 0.00},
		}

		portfolio := domain.Portfolio{}
		taxes := portfolio.CalculateTaxes(operations)
		assert.Equal(t, expectedTaxes, taxes)
	})

	t.Run("Case #5", func(t *testing.T) {
		t.Parallel()

		operations := domain.Operations{
			{
				Type:     domain.BUY,
				Quantity: 10000,
				UnitCost: 10.00,
			},
			{
				Type:     domain.BUY,
				Quantity: 5000,
				UnitCost: 25.00,
			},
			{
				Type:     domain.SELL,
				Quantity: 10000,
				UnitCost: 15.00,
			},
			{
				Type:     domain.SELL,
				Quantity: 5000,
				UnitCost: 25.00,
			},
		}

		expectedTaxes := domain.Taxes{
			{Tax: 0.00},
			{Tax: 0.00},
			{Tax: 0.00},
			{Tax: 10000.00},
		}

		portfolio := domain.Portfolio{}
		taxes := portfolio.CalculateTaxes(operations)
		assert.Equal(t, expectedTaxes, taxes)
	})

	t.Run("Case #6", func(t *testing.T) {
		t.Parallel()

		operations := domain.Operations{
			{
				Type:     domain.BUY,
				Quantity: 10000,
				UnitCost: 10.00,
			},
			{
				Type:     domain.SELL,
				Quantity: 5000,
				UnitCost: 2.00,
			},
			{
				Type:     domain.SELL,
				Quantity: 2000,
				UnitCost: 20.00,
			},
			{
				Type:     domain.SELL,
				Quantity: 2000,
				UnitCost: 20.00,
			},
			{
				Type:     domain.SELL,
				Quantity: 1000,
				UnitCost: 25.00,
			},
		}

		expectedTaxes := domain.Taxes{
			{Tax: 0.00},
			{Tax: 0.00},
			{Tax: 0.00},
			{Tax: 0.00},
			{Tax: 3000.00},
		}

		portfolio := domain.Portfolio{}
		taxes := portfolio.CalculateTaxes(operations)
		assert.Equal(t, expectedTaxes, taxes)
	})

	t.Run("Case #7", func(t *testing.T) {
		t.Parallel()

		operations := domain.Operations{
			{Type: domain.BUY, Quantity: 10000, UnitCost: 10.00},
			{Type: domain.SELL, Quantity: 5000, UnitCost: 2.00},
			{Type: domain.SELL, Quantity: 2000, UnitCost: 20.00},
			{Type: domain.SELL, Quantity: 2000, UnitCost: 20.00},
			{Type: domain.SELL, Quantity: 1000, UnitCost: 25.00},
			{Type: domain.BUY, Quantity: 10000, UnitCost: 20.00},
			{Type: domain.SELL, Quantity: 5000, UnitCost: 15.00},
			{Type: domain.SELL, Quantity: 4350, UnitCost: 30.00},
			{Type: domain.SELL, Quantity: 650, UnitCost: 30.00},
		}

		expectedTaxes := domain.Taxes{
			{Tax: 0.00},
			{Tax: 0.00},
			{Tax: 0.00},
			{Tax: 0.00},
			{Tax: 3000.00},
			{Tax: 0.00},
			{Tax: 0.00},
			{Tax: 3700.00},
			{Tax: 0.00},
		}

		portfolio := domain.Portfolio{}
		taxes := portfolio.CalculateTaxes(operations)
		assert.Equal(t, expectedTaxes, taxes)
	})

	t.Run("Case #8", func(t *testing.T) {
		t.Parallel()

		operations := domain.Operations{
			{Type: domain.BUY, Quantity: 10000, UnitCost: 10.00},
			{Type: domain.SELL, Quantity: 10000, UnitCost: 50.00},
			{Type: domain.BUY, Quantity: 10000, UnitCost: 20.00},
			{Type: domain.SELL, Quantity: 10000, UnitCost: 50.00},
		}

		expectedTaxes := domain.Taxes{
			{Tax: 0.00},
			{Tax: 80000.00},
			{Tax: 0.00},
			{Tax: 60000.00},
		}

		portfolio := domain.Portfolio{}
		taxes := portfolio.CalculateTaxes(operations)
		assert.Equal(t, expectedTaxes, taxes)
	})
}
