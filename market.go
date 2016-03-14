package godatayes

import (
	"fmt"
)

func getMktEqud(ticker, beginDate, endDate, tradeDate string) (string, error) {
	var path string
	if tradeDate == "" {
		path = fmt.Sprintf("/api/market/getMktEqud.json?field=&beginDate=%s&endDate=%s&ticker=%s",
			beginDate, endDate, ticker)
	} else {
		path = fmt.Sprintf("/api/market/getMktEqud.json?field=&beginDate=%s&endDate=%s&secID=%s&ticker=%s&tradeDate=%s",
			beginDate, endDate, "", ticker, tradeDate)
	}
	data, err := getData(path)
	return data, err
}

// GetMktEqudByTicker get market equid by ticker
func GetMktEqudByTicker(ticker, beginDate, endDate string) (string, error) {
	path := fmt.Sprintf("/api/market/getMktEqud.json?field=&beginDate=%s&endDate=%s&secID=%s&ticker=%s&tradeDate=%s",
		beginDate, endDate, "", ticker, "")
	data, err := getData(path)
	return data, err
}
