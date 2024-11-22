package repository

import (
	"encoding/json"
	"os"
	"transaction_app/entities"
)

type MerchantRepository interface {
	ReadMerchants() ([]entities.Merchant, error)
	WriteMerchants([]entities.Merchant) error
}

type merchantRepository struct {
	merchantFile string
}

func NewMerchantRepository(merchantFile string) MerchantRepository {
	return &merchantRepository{
		merchantFile: merchantFile,
	}
}

func (m *merchantRepository) ReadMerchants() ([]entities.Merchant, error) {
	data, err := os.ReadFile(m.merchantFile)
	if err != nil {
		return nil, err
	}

	var merchants []entities.Merchant
	err = json.Unmarshal(data, &merchants)

	return merchants, err
}

func (m *merchantRepository) WriteMerchants(merchants []entities.Merchant) error {
	data, err := json.MarshalIndent(merchants, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(m.merchantFile, data, os.ModePerm)
}
