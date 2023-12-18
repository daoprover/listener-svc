-- +migrate Up
CREATE TABLE dataset (
      id SERIAL PRIMARY KEY,
      token_name                VARCHAR(100),
      txs_number                int,
      timestamp                 time,
      average_txs_count 		float,
      average_google_sites	    float,
      is_there_github 		    boolean,
      token_description 		VARCHAR(200),
      number_of_user_group 	    int,
      percent_token_handlers	int
);


create FUNCTION  calculate_average_txs_count(token_name varchar(100)) returns int as '' LANGUAGE sql;
create FUNCTION  calculate_txs_count(token_name varchar(100)) returns int as '' LANGUAGE sql;


-- +migrate Down
DROP TABLE dataset;