package pgrepo

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"try-bank/database/postgres"
	"try-bank/request"
	"try-bank/util"

	"github.com/google/uuid"
)

type comsumeResponse struct {
	Status  float64     `json:"status"`
	Data    consumeData `json:"data"`
	Message string      `json:"message"`
}

type consumeData struct {
	Payment float64 `json:"request_payment"`
	VaKey   int     `json:"va_key"`
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

	bodyBytes, err := util.ConsumeAPIPost(checkRow.FqdnDetailPayment, bodyType, bodyToJson)
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

	addCompanyBalance := postgres.AddBalanceParams{
		Balance: response.Data.Payment,
		ID:      checkRow.WalletID.UUID,
	}

	payerBalance := postgres.SubtractBalanceParams{
		Balance: response.Data.Payment,
		ID:      userWallet.UUID,
	}

	pay := postgres.PayVAParams{
		ID:             uuid.New(),
		VirtualAccount: checkRow.VaID,
		VaNumber:       vaNumber,
		RequestPayment: response.Data.Payment,
	}

	return d.transaction(ctx, func(q *postgres.Queries) error {
		if err := q.PayVA(ctx, pay); err != nil {
			return err
		}

		if err := q.SubtractBalance(ctx, payerBalance); err != nil {
			return err
		}

		respBody, err := util.ConsumeAPIPost(checkRow.FqdnPay, bodyType, map[string]string{"va_number": vaNumber})
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

	fromArg := postgres.SubtractBalanceParams{
		Balance: req.TotalTransfer,
		ID:      from.UUID,
	}

	toArg := postgres.AddBalanceParams{
		Balance: req.TotalTransfer,
		ID:      to.UUID,
	}

	return d.transaction(ctx, func(q *postgres.Queries) error {
		if err := q.SubtractBalance(ctx, fromArg); err != nil {
			return err
		}

		if err := q.AddBalance(ctx, toArg); err != nil {
			return err
		}

		return nil
	})
}
