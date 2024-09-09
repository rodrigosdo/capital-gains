package domain

// MaxNoTaxOperationAmount is the maximum amount for a sale operation to be tax-free.
const MaxNoTaxOperationAmount = 20000

// Portfolio represents a portfolio of operations.
type Portfolio struct {
	AccumulatedLoss float64
	TotalQuantity   int64
	WeightedAvgCost float64
}

// CalculateTaxes calculates the taxes for a list of operations.
func (p *Portfolio) CalculateTaxes(operations Operations) Taxes {
	taxes := make(Taxes, len(operations))
	for i, operation := range operations {
		tax := p.calculateTax(operation)
		if tax != nil {
			taxes[i] = *tax
		}
	}
	return taxes
}

func (p *Portfolio) calculateTax(operation Operation) *Tax {
	switch operation.Type {
	case BUY:
		return p.buy(operation.Quantity, operation.UnitCost)
	case SELL:
		return p.sell(operation.Quantity, operation.UnitCost)
	default:
		return nil
	}
}

func (p *Portfolio) buy(quantity int64, unitCost float64) *Tax {
	totalCost := p.WeightedAvgCost * float64(p.TotalQuantity)
	newCost := unitCost * float64(quantity)
	p.TotalQuantity += quantity
	p.WeightedAvgCost = (totalCost + newCost) / float64(p.TotalQuantity)
	return &Tax{Tax: 0}
}

func (p *Portfolio) sell(quantity int64, unitCost float64) *Tax {
	profit := p.calculateProfit(quantity, unitCost)

	if unitCost*float64(quantity) <= MaxNoTaxOperationAmount {
		p.handleNoTaxSale(quantity, profit)
		return &Tax{Tax: 0}
	}

	if profit <= 0 {
		p.handleLossSale(quantity, profit)
		return &Tax{Tax: 0}
	}

	return p.handleProfitSale(quantity, profit)
}

func (p *Portfolio) calculateProfit(quantity int64, unitCost float64) float64 {
	return (unitCost - p.WeightedAvgCost) * float64(quantity)
}

func (p *Portfolio) handleNoTaxSale(quantity int64, profit float64) {
	if profit <= 0 {
		p.AccumulatedLoss += -profit
	}
	p.TotalQuantity -= quantity
}

func (p *Portfolio) handleLossSale(quantity int64, profit float64) {
	p.AccumulatedLoss += -profit
	p.TotalQuantity -= quantity
}

func (p *Portfolio) handleProfitSale(quantity int64, profit float64) *Tax {
	if p.AccumulatedLoss > 0 {
		if p.AccumulatedLoss >= profit {
			p.AccumulatedLoss -= profit
			p.TotalQuantity -= quantity
			return &Tax{Tax: 0}
		}
		profit -= p.AccumulatedLoss
		p.AccumulatedLoss = 0
	}

	p.TotalQuantity -= quantity
	return &Tax{Tax: profit * TaxRate}
}
