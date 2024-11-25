package usecase

import (
	"errors"
	"time"
	"transaction_app/entities"
	"transaction_app/repository"
	"transaction_app/services"
)

type HistoryUsecase interface {
	CreateTransaction(customerID, merchantID int, amount float64) error
}

type historyUsecase struct {
	customerRepo repository.CustomerRepository
	merchantRepo repository.MerchantRepository
	historyRepo  repository.HistoryRepository
}

func NewHistoryUsecase(
	customerRepo repository.CustomerRepository,
	merchantRepo repository.MerchantRepository,
	historyRepo repository.HistoryRepository,
) HistoryUsecase {
	return &historyUsecase{
		customerRepo: customerRepo,
		merchantRepo: merchantRepo,
		historyRepo:  historyRepo,
	}
}

func (h *historyUsecase) CreateTransaction(customerID, merchantID int, amount float64) error {
	customers, err := h.customerRepo.ReadCustomers()
	if err != nil {
		return err
	}

	var customer *entities.Customer
	for _, c := range customers {
		if c.ID == customerID {
			customer = &c
			break
		}
	}

	if customer == nil {
		return errors.New("customer not found")
	}

	if customer.Balance < amount {
		return errors.New("Insufficient balance")
	}

	merchants, err := h.merchantRepo.ReadMerchants()
	if err != nil {
		return err
	}

	var merchant *entities.Merchant
	for _, m := range merchants {
		if m.ID == merchantID {
			merchant = &m
			break
		}
	}

	if merchant == nil {
		return errors.New("merchant not found")
	}

	customer.Balance -= amount

	err = h.customerRepo.WriteCustomers(customers)
	if err != nil {
		return err
	}

	history := entities.History{
		ID:         services.GenerateID(),
		CustomerID: customerID,
		MerchantID: merchantID,
		Amount:     amount,
		Timestamp:  time.Now(),
	}

	return h.historyRepo.LogHistory(history)
}