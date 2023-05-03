CREATE TYPE "frequency" AS ENUM (
  'yearly',
  'monthly',
  'weekly',
  'daily'
);

CREATE TABLE "expenses" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "date" DATE NOT NULL default CURRENT_DATE,
  "category" TEXT NOT NULL,
  "amount" DECIMAL(10, 2) NOT NULL,
  "description" TEXT,
  "user_id" int NOT NULL,
  "recurring_expense_id" int
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" TEXT NOT NULL,
  "wage" DECIMAL(10, 2) NOT NULL,
  "ratio" DECIMAL(10, 2) NOT NULL
);

CREATE TABLE "recurring_expenses" (
  "id" SERIAL PRIMARY KEY,
  "category" TEXT NOT NULL,
  "amount" DECIMAL(10, 2) NOT NULL,
  "frequency" frequency NOT NULL,
  "start_date" DATE NOT NULL default CURRENT_DATE,
  "end_date" DATE,
  "user_id" int NOT NULL
);

ALTER TABLE "expenses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("recurring_expense_id") REFERENCES "recurring_expenses" ("id");

ALTER TABLE "recurring_expenses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

