package requests

type CreateAccountRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PassportData string `json:"passport_data"`
	Phone        string `json:"phone"`
	Address      struct {
		Country string `json:"country"`
		State   string `json:"state"`
		City    string `json:"city"`
		Street  string `json:"street"`
	} `json:"address"`
}
