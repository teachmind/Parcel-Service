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

insert into parcel_status (status_value)
Values
('pending'),
('assigned'),
('picked'),
('reached'),
('delivered');

insert into parcel (user_id, carrier_id, status, source_address, destination_address, source_time, type, price, carrier_fee, company_fee)
Values
(1, 0, 1, 'dhaka', 'rajshahi', '2021-08-13 13:14', 'document', 200, 50, 150),
(1, 0, 1, 'vola', 'rajshahi', '2021-08-13 13:14', 'document', 200, 50, 150),
(1, 0, 1, 'mongla', 'rajshahi', '2021-08-13 13:14', 'document', 200, 50, 150),
(1, 0, 1, 'natore', 'rajshahi', '2021-08-13 13:14', 'document', 200, 50, 150);

insert into carrier_request_status (status_value)
Values
('pending'),
('accept'),
('rejected');

insert into carrier_request (parcel_id, carrier_id, status)
values
(1, 2, 1),
(1, 3, 1),
(1, 4, 1),
(1, 5, 1),
(1, 6, 1);
