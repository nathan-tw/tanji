package types

import "time"

type Price struct {
	At       time.Duration `json:"at"`
	Buy      float64       `json:"buy,string"`
	Sell     float64       `json:"sell,string"`
	Open     float64       `json:"open,string"`
	Low      float64       `json:"low,string"`
	High     float64       `json:"high,string"`
	Last     float64       `json:"last,string"`
	Vol      float64       `json:"vol,string"`
	VolInBtc float64       `json:"vol_in_btc,string"`
}
