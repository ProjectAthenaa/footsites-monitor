package monitor

import (
	"fmt"
	"github.com/json-iterator/go"
	"strconv"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func (tk *Task) iteration() error {
	req, err := tk.NewRequest("GET", tk.stockendpoint, nil)
	if err != nil {
		return err
	}
	req.Headers = tk.GenerateDefaultHeaders(fmt.Sprintf("https://www.%s.com", tk.site))

	res, err := tk.Do(req)
	if err != nil {
		return err
	}

	var stockform *StockForm
	if err = json.Unmarshal(res.Body, &stockform); err != nil {
		return err
	}

	for _, style := range stockform.Sizes{
		tk.SendStyles(strconv.Itoa(style.ProductWebKey),style.Size, stockform.Style.Color, strconv.FormatFloat(stockform.Style.Price.ListPrice, 'f', -1, 64))
	}

	for _, style := range stockform.StyleVariants{
		tk.SendStyles(strconv.Itoa(style.ProductWebKey),style.Size, style.Color, strconv.FormatFloat(style.Price.ListPrice, 'f', -1, 64))
	}

	return nil
}

func (tk *Task) SendStyles(variantid, size, color, price string){
	tk.Monitor.Channel <- map[string]interface{}{
		"pid":tk.pid,
		"variantid": variantid,
		"size" :size,
		"color": color,
		"price": price,
	}
}