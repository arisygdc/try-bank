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

CREATE TABLE levels (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE wallets (
    id UUID PRIMARY KEY NOT NULL,
    balance FLOAT NOT NULL DEFAULT 0,
    last_update TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE accounts (
    id UUID PRIMARY KEY NOT NULL,
    users UUID NOT NULL UNIQUE,
    auth_info UUID NOT NULL UNIQUE,
    wallet UUID NOT NULL UNIQUE,
    level UUID NOT NULL,
    
    CONSTRAINT users
        FOREIGN KEY (users)
        REFERENCES users(id),

    CONSTRAINT auth_info
        FOREIGN KEY (auth_info)
        REFERENCES auth_info(id),

    CONSTRAINT wallets
        FOREIGN KEY (wallet)
        REFERENCES wallets(id),

    CONSTRAINT levels
        FOREIGN KEY (level)
        REFERENCES levels(id)
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
    va_key  VARCHAR(150) NOT NULL,
    domain  VARCHAR(100) NOT NULL,
    va_identity SERIAL NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE companies_account (
    id UUID NOT NULL,
    company UUID NOT NULL UNIQUE,
    auth_info UUID NOT NULL UNIQUE,
    wallet UUID NOT NULL UNIQUE,
    virtual_account UUID UNIQUE,

    CONSTRAINT companies
        FOREIGN KEY (company)
        REFERENCES companies(id),
    
     CONSTRAINT wallets
        FOREIGN KEY (wallet)
        REFERENCES wallets(id),

    CONSTRAINT auth_info
        FOREIGN KEY (auth_info)
        REFERENCES auth_info(id),
    
    CONSTRAINT virtual_account
        FOREIGN KEY (virtual_account)
        REFERENCES virtual_account(id)
);

CREATE TABLE va_payment (
    id UUID PRIMARY KEY NOT NULL,
    virtual_account UUID NOT NULL,
    va_number VARCHAR(13) NOT NULL,
    request_payment FLOAT NOT NULL,
    paid_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT virtual_account
        FOREIGN KEY (virtual_account)
        REFERENCES virtual_account(id)
);

CREATE TABLE transfers (
    id UUID PRIMARY KEY NOT NULL,
    from_wallet UUID NOT NULL,
    to_wallet UUID NOT NULL,
    balance FLOAT NOT NULL,
    transfer_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE transfers ADD FOREIGN KEY (from_wallet) REFERENCES wallets(id);

ALTER TABLE transfers ADD FOREIGN KEY (to_wallet) REFERENCES wallets(id);

CREATE INDEX ON transfers (from_wallet);

CREATE INDEX ON transfers (to_wallet);