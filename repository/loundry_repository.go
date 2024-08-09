package repository

import (
	"challenge-goapinew/model"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type loundryRepository struct {
	db *sql.DB
}

type LoundryRepository interface {
	CreateCustomer(customer model.Customer) (model.Customer, error)
	GetCustomerById(id int) (*model.Customer, error)
	UpdateCustomer(customer model.Customer) (model.Customer, error)
	DeleteCustomer(id int) error
}

func (l *loundryRepository) CreateCustomer(customer model.Customer) (model.Customer, error) {
	query := "INSERT INTO mst_customer (name, phone_number, address) VALUES (?, ?, ?)"

	result, err := l.db.Exec(query, customer.Name, customer.PhoneNumber, customer.Address)
	if err != nil {
		return model.Customer{}, fmt.Errorf("Failed to insert customer: %v", err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	customer.Id = strconv.Itoa(int(lastInsertId))

	return customer, nil
}

func (l *loundryRepository) GetCustomerById(id int) (*model.Customer, error) {
	query := "SELECT * FROM mst_customer WHERE id = ?"

	var customer model.Customer

	err := l.db.QueryRow(query, id).Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
	if err != nil {
		return nil, fmt.Errorf("Failed to get customer: %v", err)
	}

	return &customer, nil
}

func (l *loundryRepository) UpdateCustomer(customer model.Customer) (model.Customer, error) {
	query := "UPDATE mst_customer SET name = ?, phone_number = ?, address = ? WHERE id = ?"

	result, err := l.db.Exec(query, customer.Name, customer.PhoneNumber, customer.Address, customer.Id)
	if err != nil {
		fmt.Printf("Error querying query: %v\n", err)
		return customer, fmt.Errorf("Failed to update customer: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return customer, fmt.Errorf("Error: %v", err)
	}
	if rowsAffected == 0 {
		return customer, fmt.Errorf("Customer with id %s not found", customer.Id)
	}

	return customer, nil
}

func (l *loundryRepository) DeleteCustomer(id int) error {
	query := "DELETE FROM mst_customer WHERE id = ?"

	_, err := l.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete customer: %v", err)
	}

	return nil
}

func NewLoundryRepository(db *sql.DB) LoundryRepository {
	return &loundryRepository{db: db}
}
