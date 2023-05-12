CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "todolists" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar,
  "activity" varchar,
  "due_date" varchar NOT NULL,
  "status" integer NOT NULL,
  "created_at" timestamp
);

ALTER TABLE "todolists" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
