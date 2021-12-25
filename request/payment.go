package request

type PaymentVA struct {
	RegNum         int32  `from:"reg_num" json:"reg_num" binding:"required"`
	VirtualAccount string `from:"va" json:"va" binding:"required"`
}
