-- +goose Up
CREATE TABLE quotes(
    id VARCHAR(60) PRIMARY KEY,
    quote_type VARCHAR(40) NOT NULL,
    customer_id VARCHAR(60) NOT NULL,
    source VARCHAR(100) NOT NULL,
    destination VARCHAR(100) NOT NULL ,
    door_pickup VARCHAR(20),
    door_address TEXT,
    door_delivery VARCHAR(20),
    delivery_address TEXT,
    liner_id VARCHAR(60),
    partner_id VARCHAR(60),
    validity VARCHAR(20),
    transmit_days VARCHAR(20),
    free_days VARCHAR(20),
    currency VARCHAR(10),
    buy VARCHAR(20),
    sell VARCHAR(20),
    partner_tax VARCHAR(20)
);

CREATE TABLE booking(
    id VARCHAR(60) PRIMARY KEY,
    booking_request_id VARCHAR(60) NOT NULL,
    booking_status VARCHAR(20),
    customer_id VARCHAR(60),
    task_id VARCHAR(60),
    quote_id VARCHAR(60),
    milestone_id VARCHAR(60),
    liner_id VARCHAR(60),
    source VARCHAR(100),
    destination VARCHAR(100),
    city VARCHAR(60)
);

CREATE TABLE partners(
    id VARCHAR(60) PRIMARY KEY,
    name VARCHAR(100) 
);

CREATE TABLE liners(
    id VARCHAR(60) PRIMARY KEY,
    name VARCHAR(100) 
);

CREATE TABLE account_details(
    id VARCHAR(60) PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    mobile VARCHAR(20),
    roles VARCHAR(30),
    city VARCHAR(60) 
);

CREATE TABLE tasks(
    id VARCHAR(60) PRIMARY KEY,
    name VARCHAR(100) 
);

CREATE TABLE milestones(
    id VARCHAR(60) PRIMARY KEY,
    name VARCHAR(100) 
);

CREATE TABLE booking_milestone(
    id VARCHAR(60) PRIMARY KEY,
    booking_id VARCHAR(60),
    milestone_status VARCHAR(20) ,
    created_at TEXT,
    completed_at TEXT
);

CREATE TABLE booking_task(
    id VARCHAR(60) PRIMARY KEY,
    booking_id VARCHAR(60),
    task_status VARCHAR(20) ,
    created_at TEXT,
    completed_at TEXT
);

ALTER TABLE
    quotes ADD CONSTRAINT quotes_liner_id_foreign FOREIGN KEY(liner_id) REFERENCES liners(id);

ALTER TABLE
    quotes ADD CONSTRAINT quotes_partner_id_foreign FOREIGN KEY(partner_id) REFERENCES partners(id);

ALTER TABLE
    quotes ADD CONSTRAINT quotes_customer_id_foreign FOREIGN KEY(customer_id) REFERENCES account_details(id);

ALTER TABLE
    booking_milestone ADD CONSTRAINT booking_milestone_booking_id_foreign FOREIGN KEY(booking_id) REFERENCES booking(id);

ALTER TABLE
    booking_task ADD CONSTRAINT booking_task_booking_id_foreign FOREIGN KEY(booking_id) REFERENCES booking(id);

-- +goose Down
DROP TABLE quotes CASCADE;
DROP TABLE booking CASCADE;
DROP TABLE partners CASCADE;
DROP TABLE liners CASCADE;
DROP TABLE account_details CASCADE;
DROP TABLE tasks CASCADE;
DROP TABLE milestones CASCADE;
DROP TABLE booking_milestone CASCADE;
DROP TABLE booking_task CASCADE;
