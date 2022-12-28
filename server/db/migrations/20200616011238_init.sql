-- +goose Up
CREATE TABLE "quotes"(
    "id" VARCHAR(60) NOT NULL,
    "type" VARCHAR(40) NOT NULL,
    "customer_id" VARCHAR(60) NOT NULL,
    "source" VARCHAR(100) NOT NULL,
    "destination" VARCHAR(100) NOT NULL,
    "door_pickup" BOOLEAN,
    "door_address" TEXT ,
    "door_delivery" BOOLEAN ,
    "delivery_address" TEXT,
    "liner_id" VARCHAR(60) ,
    "partner_id" VARCHAR(60) ,
    "validity" DATE ,
    "transmit_days" INTEGER ,
    "free_days" INTEGER ,
    "currency" VARCHAR(10) ,
    "buy" INTEGER ,
    "sell" INTEGER ,
    "partner_tax" INTEGER 
);
ALTER TABLE
    "quotes" ADD PRIMARY KEY("id");
CREATE TABLE "booking"(
    "id" VARCHAR(60) NOT NULL,
    "booking_request_id" VARCHAR(60) ,
    "status" VARCHAR(20),
    "customer_id" VARCHAR(60),
    "task_id" VARCHAR(60),
    "quote_id" VARCHAR(60) ,
    "milestone_id" VARCHAR(60) ,
    "liner_id" VARCHAR(60) ,
    "source" VARCHAR(100) ,
    "destination" VARCHAR(100) ,
    "city" VARCHAR(60) ,
);
ALTER TABLE
    "booking" ADD PRIMARY KEY("id");
CREATE TABLE "partner"(
    "id" VARCHAR(60) NOT NULL,
    "name" VARCHAR(100) 
);
ALTER TABLE
    "partner" ADD PRIMARY KEY("id");
CREATE TABLE "liners"(
    "id" VARCHAR(60) NOT NULL,
    "name" VARCHAR(100) 
);
ALTER TABLE
    "liners" ADD PRIMARY KEY("id");
CREATE TABLE "account_details"(
    "id" VARCHAR(60) NOT NULL,
    "name" VARCHAR(100) ,
    "email" VARCHAR(100) ,
    "mobile" VARCHAR(20) ,
    "roles" VARCHAR(30) ,
    "city" VARCHAR(60) 
);
ALTER TABLE
    "account_details" ADD PRIMARY KEY("id");
CREATE TABLE "tasks"(
    "id" VARCHAR(60) NOT NULL,
    "name" VARCHAR(100) 
);
ALTER TABLE
    "tasks" ADD PRIMARY KEY("id");
CREATE TABLE "milestones"(
    "id" VARCHAR(60) NOT NULL,
    "name" VARCHAR(100) 
);
ALTER TABLE
    "milestones" ADD PRIMARY KEY("id");
CREATE TABLE "booking_milestone"(
    "id" VARCHAR(60) NOT NULL,
    "booking_id" VARCHAR(60) ,
    "status" VARCHAR(20) ,
    "crated_at" DATE ,
    "completed_at" DATE 
);
ALTER TABLE
    "booking_milestone" ADD PRIMARY KEY("id");
CREATE TABLE "booking_task"(
    "id" VARCHAR(60) NOT NULL,
    "booking_id" VARCHAR(60) ,
    "status" VARCHAR(20) ,
    "created_at" DATE ,
    "completed_at" DATE 
);
ALTER TABLE
    "booking_task" ADD PRIMARY KEY("id");
ALTER TABLE
    "booking" ADD CONSTRAINT "booking_customer_id_foreign" FOREIGN KEY("customer_id") REFERENCES "quotes"("id");
ALTER TABLE
    "quotes" ADD CONSTRAINT "quotes_liner_id_foreign" FOREIGN KEY("liner_id") REFERENCES "liners"("id");
ALTER TABLE
    "quotes" ADD CONSTRAINT "quotes_partner_id_foreign" FOREIGN KEY("partner_id") REFERENCES "partner"("id");
ALTER TABLE
    "quotes" ADD CONSTRAINT "quotes_customer_id_foreign" FOREIGN KEY("customer_id") REFERENCES "account_details"("id");
ALTER TABLE
    "booking_milestone" ADD CONSTRAINT "booking_milestone_booking_id_foreign" FOREIGN KEY("booking_id") REFERENCES "booking"("id");
ALTER TABLE
    "booking_task" ADD CONSTRAINT "booking_task_booking_id_foreign" FOREIGN KEY("booking_id") REFERENCES "booking"("id");
-- +goose Down
DROP TABLE `quotes`;
DROP TABLE `booking`;
DROP TABLE `partner`;
DROP TABLE `liners`;
DROP TABLE `account_details`;
DROP TABLE `tasks`;
DROP TABLE `milestones`;
DROP TABLE `booking_milestone`;
DROP TABLE `booking_task`;


