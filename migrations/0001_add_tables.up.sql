CREATE TABLE IF NOT EXISTS parcel_status (
  id serial PRIMARY KEY,
  status_value TEXT
);

CREATE TABLE IF NOT EXISTS parcel (
  id serial PRIMARY KEY,
  user_id int not null,
  carrier_id int not null,
  created_at timestamp,
  updated_at timestamp,
  status int not null,
  source_address TEXT not null,
  destination_address TEXT not null,
  source_time timestamp,
  type TEXT not null,
  price float,
  carrier_fee float,
  company_fee float,
  CONSTRAINT status
      FOREIGN KEY(status)
	      REFERENCES parcel_status(id)
);

CREATE TABLE IF NOT EXISTS carrier_request_status (
    id serial PRIMARY KEY,
    status_value TEXT
);

CREATE TABLE IF NOT EXISTS carrier_request (
    PRIMARY KEY(parcel_id, carrier_id),
    parcel_id int not null,
    carrier_id int not null,
    status int not null,
    CONSTRAINT parcel_id
        FOREIGN KEY(parcel_id)
            REFERENCES parcel(id)
            ON DELETE CASCADE,
    CONSTRAINT status
            FOREIGN KEY(status)
            REFERENCES carrier_request_status(id)
);

