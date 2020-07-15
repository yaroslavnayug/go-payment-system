package api

type CreateAccountRequest struct {
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	PassportData string  `json:"passport_data"`
	Phone        string  `json:"phone"`
	Address      Address `json:"address"`
}

type Address struct {
	Country string `json:"country"`
	Region  string `json:"region"`
	City    string `json:"city"`
	Street  string `json:"street"`
}
