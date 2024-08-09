package usecase

import (
	"challenge-goapinew/model"
	"challenge-goapinew/repository"
	"fmt"
)

type loundryUsecase struct {
	repo repository.LoundryRepository
}

type LoundryUseCase interface {
	CreateCustomer(customer model.Customer) (model.Customer, error)
	GetCustomerById(id int) (*model.Customer, error)
	UpdateCustomer(customer model.Customer) (model.Customer, error)
	DeleteCustomer(id int) error
}

func (l *loundryUsecase) CreateCustomer(customer model.Customer) (model.Customer, error) {
	return l.repo.CreateCustomer(customer)
}

func (l *loundryUsecase) GetCustomerById(id int) (*model.Customer, error) {
	return l.repo.GetCustomerById(id)
}

func (l *loundryUsecase) UpdateCustomer(customer model.Customer) (model.Customer, error) {
	_, err := l.repo.UpdateCustomer(customer)
	if err != nil {
		fmt.Printf("Error in UpdateCustomer usecase: %v\n", err)
		return model.Customer{}, fmt.Errorf("Customer with id %s not found", customer.Id)
	}

	return customer, nil
}

func (l *loundryUsecase) DeleteCustomer(id int) error {
	_, err := l.repo.GetCustomerById(id)
	if err != nil {
		return fmt.Errorf("Customer with id %d not found", id)
	}

	return l.repo.DeleteCustomer(id)
}

func NewLoundryUseCase(repo repository.LoundryRepository) LoundryUseCase {
	return &loundryUsecase{repo: repo}
}
