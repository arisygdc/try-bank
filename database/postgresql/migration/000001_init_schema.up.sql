CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SEQUENCE register_number
   START 1000000
   INCREMENT 1;

CREATE TABLE cutomers (
    id UUID PRIMARY KEY NOT NULL,
    firstname VARCHAR(30) NOT NULL,
    lastname VARCHAR(30) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    email VARCHAR(100) NOT NULL UNIQUE,
    birth date NOT NULL,
    phone VARCHAR(13) NOT NULL UNIQUE
);

CREATE TABLE auth_info (
    id UUID PRIMARY KEY NOT NULL,
    registered_number INT DEFAULT nextval('register_number') NOT NULL,
    pin VARCHAR(150) NOT NULL
);

CREATE TABLE account_type (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    max_transfer FLOAT NOT NULL,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE wallets (
    id UUID PRIMARY KEY NOT NULL,
    balance FLOAT NOT NULL DEFAULT 0,
    last_update TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE accounts (
    id UUID PRIMARY KEY NOT NULL,
    cutomer_id UUID NOT NULL UNIQUE,
    auth_info_id UUID NOT NULL UNIQUE,
    wallet_id UUID NOT NULL UNIQUE,
    account_type_id UUID NOT NULL,
    
    CONSTRAINT cutomers
        FOREIGN KEY (cutomer_id)
        REFERENCES cutomers(id),

    CONSTRAINT auth_info
        FOREIGN KEY (auth_info_id)
        REFERENCES auth_info(id),

    CONSTRAINT wallets
        FOREIGN KEY (wallet_id)
        REFERENCES wallets(id),

    CONSTRAINT account_type
        FOREIGN KEY (account_type_id)
        REFERENCES account_type(id)
);

CREATE TABLE companies (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(13) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE virtual_account (
    id UUID PRIMARY KEY NOT NULL,
    authorization_key VARCHAR(150) NOT NULL UNIQUE,
    identity INT NOT NULL UNIQUE,
    callback_url  VARCHAR(120) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE companies_account (
    id UUID NOT NULL,
    company_id UUID NOT NULL UNIQUE,
    auth_info_id UUID NOT NULL UNIQUE,
    wallet_id UUID NOT NULL UNIQUE,
    virtual_account_id UUID UNIQUE,

    CONSTRAINT companies
        FOREIGN KEY (company_id)
        REFERENCES companies(id),
    
    CONSTRAINT wallets
        FOREIGN KEY (wallet_id)
        REFERENCES wallets(id),

    CONSTRAINT auth_info
        FOREIGN KEY (auth_info_id)
        REFERENCES auth_info(id),
    
    CONSTRAINT virtual_account
        FOREIGN KEY (virtual_account_id)
        REFERENCES virtual_account(id)
);

CREATE TABLE issued_payment (
    id UUID PRIMARY KEY NOT NULL,
    virtual_account_id UUID NOT NULL,
    virtual_account_number INT NOT NULL,
    payment_charge FLOAT NOT NULL,

    CONSTRAINT virtual_account 
        FOREIGN KEY (virtual_account_id)
        REFERENCES virtual_account(id)
);

CREATE TABLE va_payment (
    id UUID PRIMARY KEY NOT NULL,
    issued_payment_id UUID NOT NULL,
    paid_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT issued_payment
        FOREIGN KEY (issued_payment_id)
        REFERENCES issued_payment(id)
);

CREATE TABLE transfers (
    id UUID PRIMARY KEY NOT NULL,
    from_wallet UUID NOT NULL,
    to_wallet UUID NOT NULL,
    balance FLOAT NOT NULL,
    transfered_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE transfers ADD FOREIGN KEY (from_wallet) REFERENCES wallets(id);

ALTER TABLE transfers ADD FOREIGN KEY (to_wallet) REFERENCES wallets(id);

CREATE INDEX ON transfers (from_wallet);

CREATE INDEX ON transfers (to_wallet);

INSERT INTO account_type (name, max_transfer) VALUES ('silver', 5000000);
INSERT INTO account_type (name, max_transfer) VALUES ('gold', 15000000);
INSERT INTO account_type (name, max_transfer) VALUES ('platinum', 50000000);