package model

type Customer struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json :"address"`
}

type Employee struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json :"address"`
}

type Product struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Unit  string `json:"unit"`
}

type Transaction struct {
	Id          string       `json:"id,omitempty"`
	BillDate    string       `json:"billDate"`
	EntryDate   string       `json:"entryDate"`
	FinishDate  string       `json:"finishDate"`
	EmployeeID  string       `json:"employeeId"`
	CustomerID  string       `json:"customerId"`
	BillDetails []BillDetail `json:"billDetails"`
	TotalBill   int          `json:"totalBill"`
}

type BillDetail struct {
	Id           string `json:"id"`
	BillID       string `json:"billId"`
	ProductID    string `json:"productId"`
	ProductPrice int    `json:"productPrice"`
	Qty          int    `json:"qty"`
}
