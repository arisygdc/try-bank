-- name: IssuePaymentVA :exec
INSERT INTO issued_payment (id, virtual_account_id, virtual_account_number, payment_charge) VALUES ($1, $2, $3, $4);