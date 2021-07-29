CREATE TABLE IF NOT EXISTS parcel_status (
  id serial PRIMARY KEY,
  status_value TEXT
);

CREATE TABLE IF NOT EXISTS parcel (
  id serial PRIMARY KEY,
  user_id int,
  carrier_id int,
  created_at timestamp,
  updated_at timestamp,
  status int,
  source_address TEXT,
  destination_address TEXT,
  source_time timestamp,
  type TEXT,
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
    id serial PRIMARY KEY,
    parcel_id int,
    carrier_id int,
    status int,
    CONSTRAINT parcel_id
        FOREIGN KEY(parcel_id)
            REFERENCES parcel(id),
    CONSTRAINT status
            FOREIGN KEY(status)
            REFERENCES carrier_request_status(id)
);