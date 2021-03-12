package grid

import (
	"fmt"
	"math"

	"github.com/nathan-tw/tanji/pkg/max"
)

type Strategy struct {
	symbol      string
	UpperPrice  float64
	LowerPrice  float64
	Capital     float64
	TradeAmount float64
	GridNum     int
}

// func (s *Strategy) generateSellOrders() []*SubmitOrder{
// 	currentPrice :=
// }

func (s *Strategy) generateBuyOrders() error {
	currentPrice, err := max.GetCurrPrice(s.symbol)
	if err != nil {
		return err
	}
	priceRange := s.UpperPrice - s.LowerPrice

	if priceRange <= 0 {
		return fmt.Errorf("invalid upper and lower price")
	}
	gridInterval := priceRange / float64(s.GridNum)
	startPrice := math.Max(s.LowerPrice, currentPrice.Last+gridInterval)

	if startPrice > s.UpperPrice {
		return fmt.Errorf("current price %f exceed the upper price boundary %f", currentPrice.Last, s.UpperPrice)
	}
	return nil

}
