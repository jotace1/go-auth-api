-- migrate:up
CREATE TABLE accounts (
   account_id    UUID NOT NULL,
   username       VARCHAR(100) NOT NULL,
   password       VARCHAR(100) NOT NULL,
   email       VARCHAR(100) NOT NULL,
   created_at           TIMESTAMP WITHOUT TIME ZONE,
   updated_at           TIMESTAMP WITHOUT TIME ZONE,

   CONSTRAINT pk_account PRIMARY KEY (account_id)
);

-- migrate:down
DROP TABLE accounts;