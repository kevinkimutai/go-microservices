CREATE TABLE accounts (
  id INT AUTO_INCREMENT PRIMARY KEY,
  owner VARCHAR(255) NOT NULL,
  balance BIGINT NOT NULL,
  currency VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE entries (
  id INT AUTO_INCREMENT PRIMARY KEY,
  account_id BIGINT NOT NULL,
  amount BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transfers (
  id INT AUTO_INCREMENT PRIMARY KEY,
  from_account_id BIGINT NOT NULL,
  to_account_id BIGINT NOT NULL,
  amount BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_owner ON accounts (owner);
CREATE INDEX idx_account_id ON entries (account_id);
CREATE INDEX idx_from_account_id ON transfers (from_account_id);
CREATE INDEX idx_to_account_id ON transfers (to_account_id);
CREATE INDEX idx_from_to_account_id ON transfers (from_account_id, to_account_id);

ALTER TABLE entries ADD CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES accounts (id);
ALTER TABLE transfers ADD CONSTRAINT fk_from_account_id FOREIGN KEY (from_account_id) REFERENCES accounts (id);
ALTER TABLE transfers ADD CONSTRAINT fk_to_account_id FOREIGN KEY (to_account_id) REFERENCES accounts (id);

COMMENT ON COLUMN entries.amount IS 'can be positive/negative';
COMMENT ON COLUMN transfers.amount IS 'can be positive or negative';