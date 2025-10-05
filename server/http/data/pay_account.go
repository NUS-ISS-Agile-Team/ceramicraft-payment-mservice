package data

type UserPayAccount struct {
	UserId    int    `json:"user_id"`
	AccountNo string `json:"account_no"`
	Balance   int    `json:"balance"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type UserPayAccountTopUpRequest struct {
	RedeemCode string `json:"redeem_code" binding:"required"`
}

type UserPayAccountTopUpResult struct {
	TopUpAmount    int `json:"top_up_amount"`
	CurrentBalance int `json:"current_balance"`
}
