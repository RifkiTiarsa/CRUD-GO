package controller

import (
	"challenge-goapinew/model"
	"challenge-goapinew/usecase"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type LoundryController struct {
	useCase usecase.LoundryUseCase
}

func (l *LoundryController) Route(router *mux.Router) {
	router.HandleFunc("/customers", l.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", l.GetCustomerById).Methods("GET")
	router.HandleFunc("/customers/{id}", l.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", l.DeleteCustomer).Methods("DELETE")
}

func (l *LoundryController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newCustomer, err := l.useCase.CreateCustomer(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCustomer)
}

func (l *LoundryController) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	customerId := mux.Vars(r)
	id, err := strconv.Atoi(customerId["id"])
	if err != nil {
		http.Error(w, "Invalid customer id", http.StatusBadRequest)
		return
	}

	customer, err := l.useCase.GetCustomerById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

func (l *LoundryController) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var customer model.Customer

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer.Id = id

	UpdatedCustomer, err := l.useCase.UpdateCustomer(customer)
	if err != nil {
		fmt.Printf("Received customer update request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Received customer update response: %v\n", err)

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(UpdatedCustomer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (l *LoundryController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerId := mux.Vars(r)
	id, err := strconv.Atoi(customerId["id"])
	if err != nil {
		http.Error(w, "Invalid customer id", http.StatusBadRequest)
		return
	}

	if err = l.useCase.DeleteCustomer(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted"))

}

func NewLoundryController(usecase usecase.LoundryUseCase) *LoundryController {
	return &LoundryController{useCase: usecase}
}
