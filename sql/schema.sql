CREATE TABLE IF NOT EXISTS bank_accounts (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  name VARCHAR(200) NOT NULL,
  balance DECIMAL(10, 2) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id)
);
CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  bank_account_id INTEGER NOT NULL,
  amount DECIMAL(10, 2) NOT NULL,
  timestamp TIMESTAMP NOT NULL,
  description VARCHAR(200) NOT NULL,
  credit_debit BOOLEAN NOT NULL,
  FOREIGN KEY (bank_account_id) REFERENCES bank_accounts (id)
);
CREATE TABLE IF NOT EXISTS loans (
  id SERIAL PRIMARY KEY,
  bank_account_id INTEGER NOT NULL,
  amount DECIMAL(10, 2) NOT NULL,
  interest_rate DECIMAL(5, 2) NOT NULL,
  due_date DATE NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  FOREIGN KEY (bank_account_id) REFERENCES bank_accounts (id)
);
CREATE TABLE IF NOT EXISTS investments (
  id SERIAL PRIMARY KEY,
  bank_account_id INTEGER NOT NULL,
  amount DECIMAL(10, 2) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  FOREIGN KEY (bank_account_id) REFERENCES bank_accounts (id)
);
CREATE TABLE IF NOT EXISTS fixed_expenses (
  id SERIAL PRIMARY KEY,
  bank_account_id INTEGER NOT NULL,
  amount DECIMAL(10, 2) NOT NULL,
  description VARCHAR(200) NOT NULL,
  FOREIGN KEY (bank_account_id) REFERENCES bank_accounts (id)
);
CREATE TABLE IF NOT EXISTS bitcoin (
  id SERIAL PRIMARY KEY,
  bank_account_id INTEGER NOT NULL,
  purchase_price DECIMAL(10, 2) NOT NULL,
  quantity DECIMAL(10, 8) NOT NULL,
  purchase_date DATE NOT NULL,
  FOREIGN KEY (bank_account_id) REFERENCES bank_accounts (id)
);
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(200) NOT NULL,
  email VARCHAR(200) NOT NULL,
  password VARCHAR(200) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);