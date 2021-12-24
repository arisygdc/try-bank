package request

type PaymentVA struct {
	VirtualAccount string `from:"va" json:"va" binding:"required"`
}
