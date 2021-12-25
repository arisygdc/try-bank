package request

type PaymentVA struct {
	RegNum         int32  `from:"reg_num" json:"reg_num" binding:"required"`
	VirtualAccount string `from:"va" json:"va" binding:"required"`
}

type Transfer struct {
	FromRegNum    int32   `from:"from" json:"from" binding:"required"`
	ToRegNum      int32   `from:"to" json:"to" binding:"required"`
	TotalTransfer float64 `from:"total_transfer" json:"total_transfer" binding:"required"`
}
