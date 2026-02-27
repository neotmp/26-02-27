# 26-02-27

CREATE TABLE accounts (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer,
    balance numeric,
    CONSTRAINT accounts_id_not_null 
);

-- Indices -------------------------------------------------------

CREATE UNIQUE INDEX accounts_pkey ON accounts(id int4_ops);


CREATE TABLE transactions (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer,
    amount numeric,
    currency character varying,
    status integer,
    hash character varying UNIQUE,
    destination integer DEFAULT 1,
    CONSTRAINT transactions_id_not_null 
);

-- Indices -------------------------------------------------------

CREATE UNIQUE INDEX transactions_pkey ON transactions(id int4_ops);
CREATE UNIQUE INDEX transactions_hash_key ON transactions(hash text_ops);
