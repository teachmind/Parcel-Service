CREATE TABLE IF NOT EXISTS parcel_status (
  id SERIAL PRIMARY KEY,
  status_value TEXT
);

BEGIN;
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now(); 
    RETURN NEW;
END;
$$ language 'plpgsql';
COMMIT;

BEGIN;
CREATE TABLE IF NOT EXISTS parcel (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL CHECK(user_id > 0),
  carrier_id INT NOT NULL CHECK(carrier_id > 0),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  status INT DEFAULT 1,
  source_address TEXT NOT NULL CHECK(source_address != ''),
  destination_address TEXT NOT NULL CHECK(destination_address != ''),
  source_time TIMESTAMP,
  type TEXT NOT NULL CHECK(type != ''),
  price FLOAT,
  carrier_fee FLOAT,
  company_fee FLOAT,
  CONSTRAINT status
      FOREIGN KEY(status)
	      REFERENCES parcel_status(id)
);

CREATE TRIGGER user_timestamp BEFORE INSERT OR UPDATE ON parcel
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();
COMMIT;

CREATE TABLE IF NOT EXISTS carrier_request_status (
    id SERIAL PRIMARY KEY,
    status_value TEXT
);

CREATE TABLE IF NOT EXISTS carrier_request (
    PRIMARY KEY(parcel_id, carrier_id),
    parcel_id INT NOT NULL,
    carrier_id INT NOT NULL,
    status INT NOT NULL,
    CONSTRAINT parcel_id
        FOREIGN KEY(parcel_id)
            REFERENCES parcel(id)
                ON DELETE CASCADE,
    CONSTRAINT status
            FOREIGN KEY(status)
                REFERENCES carrier_request_status(id)
);