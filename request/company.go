package request

// deprecated request
type VirtualAccount struct {
	FQDNCheck string `from:"fqdn_check" json:"fqdn_check" binding:"required"`
	FQDNPay   string `from:"fqdn_pay" json:"fqdn_pay" binding:"required"`
	Name      string `from:"name" json:"name" binding:"required"`
	Email     string `from:"email" json:"email" binding:"required,email"`
	Phone     string `from:"phone" json:"phone" binding:"required"`
	RegNum    int32  `from:"reg_num" json:"reg_num" binding:"required"`
}

type PostCompany struct {
	Name  string  `from:"name" json:"name" binding:"required"`
	Email string  `from:"email" json:"email" binding:"required,email"`
	Phone string  `from:"phone" json:"phone" binding:"required"`
	TopUp float64 `from:"top_up" json:"top_up" binding:"required"`
	Pin   string  `from:"pin" json:"pin" binding:"required"`
}
