package utils

import (
	"currency_converter/internal/converter/models"
	"encoding/json"
	"os"
)

func SaveFile(r models.Rates) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}
	err = os.WriteFile("D:/Programming/Go/currency_converter/internal/converter/data/rates.json", data, 0644)
	return err
}

func LoadRates() (models.Rates, error) {
	var r models.Rates
	data, err := os.ReadFile("D:/Programming/Go/currency_converter/internal/converter/data/rates.json")
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(data, &r)
	return r, err
}
