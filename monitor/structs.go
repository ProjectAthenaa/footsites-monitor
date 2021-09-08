package monitor

type StockForm struct {
	ID    string `json:"id"`
	Style struct {
		Color             string        `json:"color"`
		Price             struct {
			ListPrice             float64     `json:"listPrice"`
		} `json:"price"`
	} `json:"style"`
	Sizes []struct {
		ProductWebKey   int    `json:"productWebKey"`
		Size            string `json:"size"`
	} `json:"sizes"`
	StyleVariants []struct {
		ProductWebKey int         `json:"productWebKey"`
		Color         string      `json:"color"`
		Size          string      `json:"size"`
		Price         struct {
			ListPrice             float64     `json:"listPrice"`
		} `json:"price"`
	} `json:"styleVariants"`
}