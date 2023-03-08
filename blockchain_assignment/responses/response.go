package responses

type CreateAccountRequest struct {
	Name string `json:"name,omitempty" validate:"required"`
	Pin  string `json:"pin,omitempty" binding:"required"`
}

type DepositRequest struct {
	AccountNo string  `json:"account_number,omitempty" binding:required`
	Pin       string  `json:"pin,omitempty" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}

type WithdrawRequest struct {
	AccountNo string  `json:"account_number,omitempty" binding:"required"`
	Pin       string  `json:"pin,omitempty" binding:"required"`
	Amount    float64 `json:"amount,omitempty" binding:"required"`
}

type TransferRequest struct {
	FromAccount string  `json:"from_account,omitempty" binding:"required"`
	FromPin     string  `json:"from_pin,omitempty" binding:"required"`
	ToAccount   string  `json:"to_account,omitempty" binding:"required"`
	Amount      float64 `json:"amount,omitempty" binding:"required"`
}

type PinRequest struct {
	AccountNo string `json:"account_number,omitempty" binding:"required"`
	OldPin    string `json:"old_pin,omitempty" binding:"required"`
	NewPin    string `json:"new_pin,omitempty" binding:"required"`
}
