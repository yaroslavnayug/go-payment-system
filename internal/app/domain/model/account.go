package model

type Account struct {
	uid          uint64
	firstName    string
	lastName     string
	passportData string
	phone        string
	address      Address
}

type Address struct {
	country string
	state   string
	city    string
	street  string
}
