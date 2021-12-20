// Code generated by sqlc. DO NOT EDIT.

package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID         uuid.UUID `json:"id"`
	Users      uuid.UUID `json:"users"`
	AuthInfo   uuid.UUID `json:"auth_info"`
	Wallet     uuid.UUID `json:"wallet"`
	Permission uuid.UUID `json:"permission"`
}

type AccountHaveCompany struct {
	Account       uuid.UUID `json:"account"`
	Company       uuid.UUID `json:"company"`
	CompanyWallet uuid.UUID `json:"company_wallet"`
}

type AuthInfo struct {
	ID               uuid.UUID `json:"id"`
	RegisteredNumber int32     `json:"registered_number"`
	Pin              string    `json:"pin"`
}

type CompaniesWallet struct {
	ID         uuid.UUID `json:"id"`
	Balance    float64   `json:"balance"`
	LastUpdate time.Time `json:"last_update"`
}

type Company struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	CompanyKey string    `json:"company_key"`
}

type CoustomerWallet struct {
	ID         uuid.UUID `json:"id"`
	Balance    float64   `json:"balance"`
	LastUpdate time.Time `json:"last_update"`
}

type PermissionLevel struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Transfer struct {
	ID          uuid.UUID `json:"id"`
	FromAccount uuid.UUID `json:"from_account"`
	ToAccount   uuid.UUID `json:"to_account"`
	Balance     float64   `json:"balance"`
	TransferAt  time.Time `json:"transfer_at"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	Birth     time.Time `json:"birth"`
	Phone     string    `json:"phone"`
}

type VirtualAccount struct {
	ID             uuid.UUID `json:"id"`
	CompanyID      uuid.UUID `json:"company_id"`
	RequestPayment float64   `json:"request_payment"`
	VaNumber       string    `json:"va_number"`
	PaidAt         time.Time `json:"paid_at"`
}
