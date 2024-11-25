package repository

import (
	"encoding/json"
	"errors"
	"os"
	"transaction_app/entities"
)

type CustomerRepository interface {
	ReadCustomers() ([]entities.Customer, error)
	WriteCustomers([]entities.Customer) error
}

type customerRepository struct {
	customerFile string
}

func NewCustomerRepository(customerFile string) CustomerRepository {
	return &customerRepository{
		customerFile: customerFile,
	}
}

func (c *customerRepository) ReadCustomers() ([]entities.Customer, error) {
	data, err := os.ReadFile(c.customerFile)
	if err != nil {
		return nil, err
	}

	var customers []entities.Customer
	err = json.Unmarshal(data, &customers)

	return customers, err
}

func (c *customerRepository) WriteCustomers(customers []entities.Customer) error {
	data, err := json.MarshalIndent(customers, "", "  ")
	if err != nil {
		return errors.New("failed to marshal customer file")
	}

	return os.WriteFile(c.customerFile, data, os.ModePerm)
}
