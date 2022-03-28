package data

type WithdrawRequest struct {
	UserId int `json:"userId"`
	Amount int `json:"amount"`
}

type DepositRequest struct {
	UserId int `json:"userId"`
	Amount int `json:"amount"`
}

type TransferRequest struct {
	SenderId   int `json:"senderId"`
	RecieverId int `json:"recieverId"`
	Amount     int `json:"amount"`
}
