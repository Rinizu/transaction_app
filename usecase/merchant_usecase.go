package usecase

import (
	"transaction_app/entities"
	"transaction_app/repository"
)

type MerchantUsecase interface {
	GetMerchants() ([]entities.Merchant, error)
}

type merchantUsecase struct {
	merchantRepo repository.MerchantRepository
}

func NewMerchantUsecase(merchantRepo repository.MerchantRepository) MerchantUsecase {
	return &merchantUsecase{
		merchantRepo: merchantRepo,
	}
}

func (m *merchantUsecase) GetMerchants() ([]entities.Merchant, error) {
	return m.merchantRepo.ReadMerchants()
}
