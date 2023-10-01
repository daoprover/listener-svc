-- +migrate Up
CREATE TABLE transactions (
      id SERIAL PRIMARY KEY,
      sender CHAR(42),
      recipient CHAR(42),
      hash VARCHAR(66) ,
      value_to numeric,
      timestamp_to timestamp,
      currency  VARCHAR(10),
)

-- +migrate Down
DROP TABLE transactions;