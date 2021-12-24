// Code generated by sqlc. DO NOT EDIT.

package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID       uuid.UUID `json:"id"`
	Users    uuid.UUID `json:"users"`
	AuthInfo uuid.UUID `json:"auth_info"`
	Wallet   uuid.UUID `json:"wallet"`
	Level    uuid.UUID `json:"level"`
}

type AuthInfo struct {
	ID               uuid.UUID `json:"id"`
	RegisteredNumber int32     `json:"registered_number"`
	Pin              string    `json:"pin"`
}

type CompaniesAccount struct {
	ID             uuid.UUID     `json:"id"`
	Company        uuid.UUID     `json:"company"`
	AuthInfo       uuid.UUID     `json:"auth_info"`
	Wallet         uuid.UUID     `json:"wallet"`
	VirtualAccount uuid.NullUUID `json:"virtual_account"`
}

type Company struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

type Level struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Transfer struct {
	ID         uuid.UUID `json:"id"`
	FromWallet uuid.UUID `json:"from_wallet"`
	ToWallet   uuid.UUID `json:"to_wallet"`
	Balance    float64   `json:"balance"`
	TransferAt time.Time `json:"transfer_at"`
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

type VaPayment struct {
	ID             uuid.UUID `json:"id"`
	VirtualAccount uuid.UUID `json:"virtual_account"`
	VaNumber       string    `json:"va_number"`
	RequestPayment float64   `json:"request_payment"`
	PaidAt         time.Time `json:"paid_at"`
}

type VirtualAccount struct {
	ID         uuid.UUID `json:"id"`
	VaKey      string    `json:"va_key"`
	Domain     string    `json:"domain"`
	VaIdentity int64     `json:"va_identity"`
	CreatedAt  time.Time `json:"created_at"`
}

type Wallet struct {
	ID         uuid.UUID `json:"id"`
	Balance    float64   `json:"balance"`
	LastUpdate time.Time `json:"last_update"`
}
