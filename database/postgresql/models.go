// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package postgresql

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
	ID               uuid.UUID     `json:"id"`
	CompanyID        uuid.UUID     `json:"company_id"`
	AuthInfoID       uuid.UUID     `json:"auth_info_id"`
	WalletID         uuid.UUID     `json:"wallet_id"`
	VirtualAccountID uuid.NullUUID `json:"virtual_account_id"`
}

type Company struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

type IssuedPayment struct {
	ID                   uuid.UUID `json:"id"`
	VirtualAccountID     uuid.UUID `json:"virtual_account_id"`
	VirtualAccountNumber int32     `json:"virtual_account_number"`
	PaymentCharge        float64   `json:"payment_charge"`
}

type Level struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Transfer struct {
	ID           uuid.UUID `json:"id"`
	FromWallet   uuid.UUID `json:"from_wallet"`
	ToWallet     uuid.UUID `json:"to_wallet"`
	Balance      float64   `json:"balance"`
	TransferedAt time.Time `json:"transfered_at"`
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
	ID              uuid.UUID `json:"id"`
	IssuedPaymentID uuid.UUID `json:"issued_payment_id"`
	PaidAt          time.Time `json:"paid_at"`
}

type VirtualAccount struct {
	ID               uuid.UUID `json:"id"`
	AuthorizationKey string    `json:"authorization_key"`
	Identity         int32     `json:"identity"`
	CallbackUrl      string    `json:"callback_url"`
	CreatedAt        time.Time `json:"created_at"`
}

type Wallet struct {
	ID         uuid.UUID `json:"id"`
	Balance    float64   `json:"balance"`
	LastUpdate time.Time `json:"last_update"`
}