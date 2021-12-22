CREATE TABLE users (
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
    registered_number INT NOT NULL,
    pin VARCHAR(150) NOT NULL
);

CREATE TABLE permission_level (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE coustomer_wallet (
    id UUID PRIMARY KEY NOT NULL,
    balance FLOAT NOT NULL DEFAULT 0,
    last_update TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE accounts (
    id UUID PRIMARY KEY NOT NULL,
    users UUID NOT NULL UNIQUE,
    auth_info UUID NOT NULL UNIQUE,
    wallet UUID NOT NULL UNIQUE,
    permission UUID NOT NULL,
    
    CONSTRAINT users
        FOREIGN KEY (users)
        REFERENCES users(id),

    CONSTRAINT auth_info
        FOREIGN KEY (auth_info)
        REFERENCES auth_info(id),

    CONSTRAINT wallet
        FOREIGN KEY (wallet)
        REFERENCES coustomer_wallet(id),

    CONSTRAINT permission_level
        FOREIGN KEY (permission)
        REFERENCES permission_level(id)
);

CREATE TABLE companies (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    company_key VARCHAR(128) NOT NULL,
    domain VARCHAR(150) NOT NULL
);

CREATE TABLE companies_wallet (
    id UUID PRIMARY KEY NOT NULL,
    balance FLOAT NOT NULL DEFAULT 0,
    last_update TIMESTAMP NOT NULL DEFAULT NOW()
);  

CREATE TABLE account_have_company (
    account UUID PRIMARY KEY NOT NULL,
    company UUID NOT NULL UNIQUE,
    company_wallet UUID NOT NULL UNIQUE,

    CONSTRAINT accounts
        FOREIGN KEY (account)
        REFERENCES accounts(id),

    CONSTRAINT companies
        FOREIGN KEY (company)
        REFERENCES companies(id),
    
    CONSTRAINT companies_wallet
        FOREIGN KEY(company_wallet)
        REFERENCES companies_wallet(id)
);

CREATE TABLE virtual_account (
    id UUID PRIMARY KEY NOT NULL,
    company_id UUID NOT NULL,
    request_payment FLOAT NOT NULL,
    va_number VARCHAR(15) NOT NULL,
    paid_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE transfers (
    id UUID PRIMARY KEY NOT NULL,
    from_account UUID NOT NULL,
    to_account UUID NOT NULL,
    balance FLOAT NOT NULL,
    transfer_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE transfers ADD FOREIGN KEY (from_account) REFERENCES accounts (id);

ALTER TABLE transfers ADD FOREIGN KEY (to_account) REFERENCES accounts (id);

CREATE INDEX ON transfers (from_account);

CREATE INDEX ON transfers (to_account);