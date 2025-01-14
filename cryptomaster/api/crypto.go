package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/nickemma/cryptomaster_and_web_server/model"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*model.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters required %d received", len(currency))
	}
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	// network or domain error
	if err != nil {
		return nil, err
	}
	var response DataResponse
	if res.StatusCode == http.StatusOK {
		bodyData, err := io.ReadAll(res.Body)

		// unexpected error on getting the response
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodyData, &response)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("internal server error %v", res.StatusCode)
	}
	rate := model.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}
