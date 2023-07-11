package coin_cap

import "time"

type AssetsResponse struct {
	Data      []assetData   `json:"data"`
	Timestamp time.Duration `json:"timestamp"`
}

type AssetResponse struct {
	Data      assetData     `json:"data"`
	Timestamp time.Duration `json:"timestamp"`
}

type assetData struct {
	ID                string  `json:"id"`
	Rank              int     `json:"rank,string"`
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	Supply            string  `json:"supply"`
	MaxSupply         string  `json:"maxSupply"`
	MarketCapUsd      string  `json:"marketCapUsd"`
	VolumeUsd24Hr     string  `json:"volumeUsd24Hr"`
	PriceUsd          float64 `json:"priceUsd,string"`
	ChangePercent24Hr string  `json:"changePercent24Hr"`
	VWAP24Hr          string  `json:"vwap24Hr"`
	Explorer          string  `json:"explorer"`
}
