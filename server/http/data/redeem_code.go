package data

type RedeemCodeVO struct {
	Id        int    `json:"id"`
	Code      string `json:"code" binding:"required"`
	Amount    int    `json:"amount" binding:"required"`
	Used      bool   `json:"used"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type RedeemCodeGenRequest struct {
	Amount int `form:"amount" binding:"required"`
	Count  int `form:"count" binding:"required"`
}

type RedeemCodeGenResult struct {
	GenCount int `json:"gen_count"`
}

type RedeemCodeQuery struct {
	Code  string `form:"code"`
	Used  *bool  `form:"used"`
	Limit int    `form:"limit,default=10"`
}
