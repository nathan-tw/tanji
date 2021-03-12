package main

import (
	"fmt"

	"github.com/nathan-tw/tanji/pkg/max"
)

func main() {
	// key := os.Getenv("MAX_API_KEY")
	// secret := os.Getenv("MAX_API_SECRET")

	p, err := max.GetCurrPrice("ethusdt")
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}




