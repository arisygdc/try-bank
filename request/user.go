package request

type PostUser struct {
	Firstname string  `from:"firstname" json:"firstname" binding:"required"`
	Lastname  string  `from:"lastname" json:"lastname" binding:"required"`
	Email     string  `from:"email" json:"email" binding:"required"`
	Birth     string  `from:"birth" json:"birth" binding:"required"`
	TopUp     float64 `from:"deposit" json:"deposit" binding:"required"`
	Phone     string  `from:"phone" json:"phone" binding:"required"`
	Pin       string  `from:"pin" json:"pin" binding:"required"`
}
