package converter

import (
	"currency_converter/internal/converter/models"
	"encoding/json"
	"io"
	"net/http"
)

func Request() models.Rates {
	response, err := http.Get("https://v6.exchangerate-api.com/v6/1d734d6c29003586488a43ac/latest/USD")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	var rates models.Rates
	err = json.Unmarshal(body, &rates)

	if err != nil {
		panic(err)
	}

	return rates

}
