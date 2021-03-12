package max

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/nathan-tw/tanji/pkg/types"
)


func GetCurrPrice(symbol string) (*types.Price, error) {
	p := new(types.Price)
	resp, err := http.Get("https://max-api.maicoin.com/api/v2/tickers/" + symbol)
	if err != nil {
		return nil, err
	}
	
	temp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(temp), p)
	if err != nil {
		return nil, err
	}
	// strconv.ParseFloat()
	
	return p, nil
}