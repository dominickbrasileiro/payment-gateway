package usecase

type TransactionDtoInput struct {
	ID                 string `json:"id"`
	AccountID          string `json:"account_id"`
	Amount             int    `json:"amount"`
	CreditCardNumber   string `json:"credit_card_number"`
	CreditCardName     string `json:"credit_card_name"`
	CreditCardExpMonth int    `json:"credit_card_exp_month"`
	CreditCardExpYear  int    `json:"credit_card_exp_year"`
	CreditCardCvv      string `json:"credit_card_cvv"`
}

type TransactionDtoOutput struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}
