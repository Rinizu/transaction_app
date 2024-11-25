package usecase

import (
	"errors"
	"transaction_app/entities"
	"transaction_app/repository"
	"transaction_app/services"

	"golang.org/x/crypto/bcrypt"
)

type CustomerUsecase interface {
	Login(email string, password string) (string, error)
	RegisterCustomer(name, email, password string) error
}

type customerUsecase struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerUsecase(customerRepo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
	}
}

func (c *customerUsecase) Login(email string, password string) (string, error) {
	customers, err := c.customerRepo.ReadCustomers()
	if err != nil {
		return "", errors.New("customer file not found")
	}

	for _, customer := range customers {
		if customer.Email == email && bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password)) == nil {
			token, err := services.GenerateJWT(customer.ID)
			if err != nil {
				return "", errors.New("failed to generate token")
			}
			return token, nil
		}
	}

	return "", errors.New("invalid email or password")
}

func (c *customerUsecase) RegisterCustomer(name, email, password string) error {
	customers, err := c.customerRepo.ReadCustomers()
	if err != nil {
		return err
	}

	for _, customer := range customers {
		if customer.Email == email {
			return errors.New("email already exists")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newCustomer := entities.Customer{
		ID:       services.GenerateID(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Balance:  0,
	}

	customers = append(customers, newCustomer)
	return c.customerRepo.WriteCustomers(customers)
}
