package converter

import (
	"currency_converter/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func Request() (models.Rates, error) {
	const url = "https://v6.exchangerate-api.com/v6/1d734d6c29003586488a43ac/latest/USD"

	resp, err := httpClient.Get(url)
	if err != nil {
		return models.Rates{}, fmt.Errorf("не удалось выполнить запрос к API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Rates{}, fmt.Errorf("API вернул статус: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Rates{}, fmt.Errorf("не удалось прочитать ответ API: %w", err)
	}

	var rates models.Rates
	if err := json.Unmarshal(body, &rates); err != nil {
		return models.Rates{}, fmt.Errorf("не удалось распарсить JSON: %w", err)
	}

	return rates, nil

}
