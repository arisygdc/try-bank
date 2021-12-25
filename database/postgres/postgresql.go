package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"try-bank/request"
	"try-bank/util"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type DB struct {
	conn    *sql.DB
	queries *Queries
}

type comsumeResponse struct {
	Status  float64     `json:"status"`
	Data    consumeData `json:"data"`
	Message string      `json:"message"`
}

type consumeData struct {
	Payment float64 `json:"request_payment"`
	VaKey   int     `json:"va_key"`
}

func NewPostgres(dbdriver, dbsource string) (database *DB, err error) {
	sqlconn, err := sql.Open(dbdriver, dbsource)
	if err != nil {
		return
	}

	database = &DB{
		conn:    sqlconn,
		queries: New(sqlconn),
	}
	return
}

func (d DB) transaction(ctx context.Context, queryFunc func(*Queries) error) error {
	tx, err := d.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queriesTx := d.queries.WithTx(tx)
	err = queryFunc(queriesTx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}

func (d DB) CreateLevel(ctx context.Context, req request.PermissionReq) error {
	return d.queries.CreateLevel(ctx, CreateLevelParams{
		ID:   uuid.New(),
		Name: req.Name,
	})
}

func (d DB) CreateUser(ctx context.Context, req request.PostUser, permission string) error {
	t, err := time.Parse("2006-1-2", strings.Trim(req.Birth, " "))
	if err != nil {
		return err
	}

	user := CreateUserParams{
		ID:        uuid.New(),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Birth:     t.UTC(),
		Phone:     req.Phone,
	}

	authInfo := CreateAuthInfoParams{
		ID:               uuid.New(),
		RegisteredNumber: int32(t.Month()) + int32(req.Phone[9]+req.Phone[10]+req.Phone[11]),
		Pin:              req.Pin,
	}

	wallet := CreateWalletParams{
		ID:      uuid.New(),
		Balance: req.TopUp,
	}

	return d.transaction(ctx, func(query *Queries) error {
		if err := query.CreateUser(ctx, user); err != nil {
			return err
		}

		if err := query.CreateAuthInfo(ctx, authInfo); err != nil {
			return err
		}

		if err := query.CreateWallet(ctx, wallet); err != nil {
			return err
		}

		permID, err := query.GetLevelID(ctx, permission)
		if err != nil {
			return err
		}

		err = query.CreateAccount(ctx, CreateAccountParams{
			ID:       uuid.New(),
			Users:    user.ID,
			AuthInfo: authInfo.ID,
			Wallet:   wallet.ID,
			Level:    permID,
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func (d DB) CreateCompany(ctx context.Context, req request.PostCompany) error {
	company := CreateCompanyParams{
		ID:    uuid.New(),
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	wallet := CreateWalletParams{
		ID:      uuid.New(),
		Balance: req.TopUp,
	}

	authInfo := CreateAuthInfoParams{
		ID:               uuid.New(),
		RegisteredNumber: int32(req.Phone[9] + req.Phone[10] + req.Phone[11]),
		Pin:              req.Pin,
	}

	companyAccount := CreateCompanyAccountParams{
		ID:       uuid.New(),
		Company:  company.ID,
		AuthInfo: authInfo.ID,
		Wallet:   wallet.ID,
	}

	return d.transaction(ctx, func(q *Queries) error {
		if err := q.CreateCompany(ctx, company); err != nil {
			return err
		}

		if err := q.CreateWallet(ctx, wallet); err != nil {
			return err
		}

		if err := q.CreateAuthInfo(ctx, authInfo); err != nil {
			return err
		}

		if err := q.CreateCompanyAccount(ctx, companyAccount); err != nil {
			return err
		}
		return nil
	})
}

func (d DB) ActivateVA(ctx context.Context, req request.VirtualAccount) error {
	validateComp := ValidateCompanyParams{
		Name:             req.Name,
		Email:            req.Email,
		Phone:            req.Phone,
		RegisteredNumber: req.RegNum,
	}

	virtualAccount := CreateVirtualAccountParams{
		ID:                uuid.New(),
		VaKey:             util.RandString(32),
		FqdnDetailPayment: req.FQDNCheck,
		FqdnPay:           req.FQDNPay,
	}

	accountVA := UpdateVAstatusParams{
		VirtualAccount: uuid.NullUUID{
			UUID:  virtualAccount.ID,
			Valid: true,
		},
	}

	return d.transaction(ctx, func(q *Queries) error {
		vaID, err := q.ValidateCompany(ctx, validateComp)
		if err != nil {
			return err
		}

		accountVA.ID = vaID.UUID
		if err := q.CreateVirtualAccount(ctx, virtualAccount); err != nil {
			return err
		}

		if err := q.UpdateVAstatus(ctx, accountVA); err != nil {
			return err
		}
		return nil
	})
}

func (d DB) PaymentVA(ctx context.Context, req request.PaymentVA) error {
	var (
		bodyType   = "application/json; charset=utf-8"
		bodyToJson = make(map[string]string)
	)

	vaIdentity, err := strconv.Atoi(req.VirtualAccount[:len(req.VirtualAccount)-13])
	if err != nil {
		return err
	}

	checkRow, err := d.queries.CheckVA(ctx, int32(vaIdentity))
	if err != nil {
		return err
	}

	userWallet, err := d.queries.GetUserWalletFromAuthInfo(ctx, req.RegNum)
	if err != nil {
		return err
	}

	vaNumber := req.VirtualAccount[len(req.VirtualAccount)-13:]
	bodyToJson["va_number"] = vaNumber

	bodyBytes, err := ConsumeAPIPost(checkRow.FqdnDetailPayment, bodyType, bodyToJson)
	if err != nil {
		return err
	}

	var response comsumeResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return err
	}

	userBalance, err := d.queries.GetBalance(ctx, userWallet.UUID)
	if err != nil {
		return err
	}

	if response.Data.Payment < 10000 {
		return nil
	}

	if userBalance < response.Data.Payment {
		return errors.New("insufficient funds")
	}

	addCompanyBalance := AddBalanceParams{
		Balance: response.Data.Payment,
		ID:      checkRow.WalletID.UUID,
	}

	payerBalance := SubtractBalanceParams{
		Balance: response.Data.Payment,
		ID:      userWallet.UUID,
	}

	pay := PayVAParams{
		ID:             uuid.New(),
		VirtualAccount: checkRow.VaID,
		VaNumber:       vaNumber,
		RequestPayment: response.Data.Payment,
	}

	return d.transaction(ctx, func(q *Queries) error {
		if err := q.PayVA(ctx, pay); err != nil {
			return err
		}

		if err := q.SubtractBalance(ctx, payerBalance); err != nil {
			return err
		}

		respBody, err := ConsumeAPIPost(checkRow.FqdnPay, bodyType, map[string]string{"va_number": vaNumber})
		if err != nil {
			return err
		}

		var response struct {
			Message string `json:"message"`
		}
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return err
		}

		if err := q.AddBalance(ctx, addCompanyBalance); err != nil {
			return err
		}

		return nil
	})
}

func (d DB) Transfer(ctx context.Context, req request.Transfer) error {
	from, err := d.queries.GetUserWalletFromAuthInfo(ctx, req.FromRegNum)
	if err != nil {
		return err
	}

	to, err := d.queries.GetUserWalletFromAuthInfo(ctx, req.ToRegNum)
	if err != nil {
		return err
	}

	balance, _ := d.queries.GetBalance(ctx, from.UUID)
	if balance < req.TotalTransfer {
		return errors.New("insufficient funds")
	}

	fromArg := SubtractBalanceParams{
		Balance: req.TotalTransfer,
		ID:      from.UUID,
	}

	toArg := AddBalanceParams{
		Balance: req.TotalTransfer,
		ID:      to.UUID,
	}

	return d.transaction(ctx, func(q *Queries) error {
		if err := q.SubtractBalance(ctx, fromArg); err != nil {
			return err
		}

		if err := q.AddBalance(ctx, toArg); err != nil {
			return err
		}

		return nil
	})
}

func ConsumeAPIPost(targetConsume string, bodyType string, bodyTemplate map[string]string) (body []byte, err error) {
	bodyJson, err := json.Marshal(bodyTemplate)
	if err != nil {
		return
	}
	resp, err := http.Post(targetConsume, bodyType, bytes.NewBuffer(bodyJson))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	return
}
