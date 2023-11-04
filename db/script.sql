CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    document_number varchar(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS operation_types (
   id SERIAL PRIMARY KEY,
   description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    account_id INTEGER NOT NULL,
    operation_id INTEGER NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    event_date TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_account
    FOREIGN KEY(account_id)
    REFERENCES accounts(id),
    CONSTRAINT fk_operation
    FOREIGN KEY(operation_id)
    REFERENCES operation_types(id)
);