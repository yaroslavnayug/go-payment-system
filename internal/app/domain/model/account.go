package model

type Account struct {
	Id           uint64
	FirstName    string
	LastName     string
	PassportData string
	Phone        string
	Address      Address
}

type Address struct {
	Country string
	State   string
	City    string
	Street  string
}
