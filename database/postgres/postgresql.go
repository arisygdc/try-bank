package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
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
	Queries *Queries
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
		Queries: New(sqlconn),
	}
	return
}

func (d DB) transaction(ctx context.Context, queryFunc func(*Queries) error) error {
	tx, err := d.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queriesTx := d.Queries.WithTx(tx)
	err = queryFunc(queriesTx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
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
		ID:     uuid.New(),
		VaKey:  util.RandString(32),
		Domain: req.Domain,
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

// Still using dummies balance
func (d DB) PaymentVA(ctx context.Context, req request.PaymentVA) error {
	var (
		targetAPIConsume = "http://192.168.1.235/web-kwh//?api=check-bayar"
		bodyType         = "application/json; charset=utf-8"
		bodyToJson       = make(map[string]string)
	)

	vaIdentity, err := strconv.Atoi(req.VirtualAccount[:len(req.VirtualAccount)-13])
	if err != nil {
		return err
	}

	checkRow, err := d.Queries.CheckVA(ctx, int32(vaIdentity))
	if err != nil {
		return err
	}

	vaNumber := req.VirtualAccount[len(req.VirtualAccount)-13:]
	bodyToJson["va_number"] = vaNumber

	response, err := callTarget(targetAPIConsume, bodyType, bodyToJson)
	if err != nil {
		return err
	}
	balance := UpdateBalanceParams{
		Balance: response.Data.Payment,
		ID:      checkRow.WalletID.UUID,
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
		if err := q.UpdateBalance(ctx, balance); err != nil {
			return err
		}
		return nil
	})
}

func callTarget(targetConsume string, bodyType string, bodyTemplate map[string]string) (response comsumeResponse, err error) {
	bodyJson, err := json.Marshal(bodyTemplate)
	if err != nil {
		return
	}
	resp, err := http.Post(targetConsume, bodyType, bytes.NewBuffer(bodyJson))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &response)
	return
}
