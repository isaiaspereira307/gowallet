-- name: CreateBankAccount :exec
INSERT INTO bank_accounts (id, user_id, name, balance, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6);

-- name: CreateTransaction :exec
INSERT INTO transactions (id, bank_account_id, amount, timestamp, description, credit_debit) VALUES ($1,$2,$3,$4,$5,$6);

-- name: CreateLoan :exec
INSERT INTO loans (id, bank_account_id, amount, interest_rate, due_date, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7);

-- name: CreateInvestment :exec
INSERT INTO investments (id, bank_account_id, amount, created_at, updated_at) VALUES ($1,$2,$3,$4,$5);

-- name: CreateFixedExpense :exec
INSERT INTO fixed_expenses (id, bank_account_id, amount, description) VALUES ($1,$2,$3,$4);

-- name: CreateBitcoin :exec
INSERT INTO bitcoin (id, bank_account_id, purchase_price, quantity, purchase_date) VALUES ($1,$2,$3,$4,$5);

-- name: CreateUser :exec
INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6);

-- name: GetBankAccountById :one
SELECT * FROM bank_accounts WHERE id = $1;

-- name: GetBankAccountsByUserId :many
SELECT * FROM bank_accounts WHERE user_id = $1;

-- name: GetTransactionById :one
SELECT * FROM transactions WHERE id = $1;

-- name: GetTransactionsByBankAccountId :many
SELECT * FROM transactions WHERE bank_account_id = $1;

-- name: GetLoanById :one
SELECT * FROM loans WHERE id = $1;

-- name: GetLoansByBankAccountId :many
SELECT * FROM loans WHERE bank_account_id = $1;

-- name: GetInvestmentById :one
SELECT * FROM investments WHERE id = $1;

-- name: GetInvestmentsByBankAccountId :many
SELECT * FROM investments WHERE bank_account_id = $1;

-- name: GetFixedExpenseById :one
SELECT * FROM fixed_expenses WHERE id = $1;

-- name: GetFixedExpensesByBankAccountId :many
SELECT * FROM fixed_expenses WHERE bank_account_id = $1;

-- name: GetBitcoinById :one
SELECT * FROM bitcoin WHERE id = $1;

-- name: GetBitcoinsByBankAccountId :many
SELECT * FROM bitcoin WHERE bank_account_id = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: UpdateBankAccountBalance :exec
UPDATE bank_accounts SET balance = $1 WHERE id = $2;

-- name: UpdateBankAccount :exec
UPDATE bank_accounts SET name = $1, balance = $2, updated_at = $3 WHERE id = $4;

-- name: UpdateTransaction :exec
UPDATE transactions SET amount = $1, timestamp = $2, description = $3, credit_debit = $4 WHERE id = $5;

-- name: UpdateLoan :exec
UPDATE loans SET amount = $1, interest_rate = $2, due_date = $3, updated_at = $4 WHERE id = $5;

-- name: UpdateInvestment :exec
UPDATE investments SET amount = $1, updated_at = $2 WHERE id = $3;

-- name: UpdateFixedExpense :exec
UPDATE fixed_expenses SET amount = $1, description = $2 WHERE id = $3;

-- name: UpdateBitcoin :exec
UPDATE bitcoin SET purchase_price = $1, quantity = $2, purchase_date = $3 WHERE id = $4;

-- name: UpdateUser :exec
UPDATE users SET name = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5;

-- name: DeleteBankAccount :exec
DELETE FROM bank_accounts WHERE id = $1;

-- name: DeleteTransaction :exec
DELETE FROM transactions WHERE id = $1;

-- name: DeleteLoan :exec
DELETE FROM loans WHERE id = $1;

-- name: DeleteInvestment :exec
DELETE FROM investments WHERE id = $1;

-- name: DeleteFixedExpense :exec
DELETE FROM fixed_expenses WHERE id = $1;

-- name: DeleteBitcoin :exec
DELETE FROM bitcoin WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
