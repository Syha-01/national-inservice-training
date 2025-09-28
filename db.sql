CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "created_at" "timestamp(0) with time zone" NOT NULL DEFAULT (now()),
  "officer_reg_number" text UNIQUE,
  "email" citext UNIQUE NOT NULL,
  "password_hash" bytea NOT NULL,
  "activated" bool NOT NULL DEFAULT false,
  "role" text NOT NULL,
  "version" integer NOT NULL DEFAULT 1
);

CREATE TABLE "officers" (
  "regulation_number" text PRIMARY KEY,
  "created_at" "timestamp(0) with time zone" NOT NULL DEFAULT (now()),
  "updated_at" "timestamp(0) with time zone" NOT NULL DEFAULT (now()),
  "first_name" text NOT NULL,
  "last_name" text NOT NULL,
  "sex" text NOT NULL,
  "rank_id" bigint NOT NULL,
  "formation_id" bigint NOT NULL,
  "posting_id" bigint NOT NULL,
  "is_active" bool NOT NULL DEFAULT true,
  "version" integer NOT NULL DEFAULT 1
);

CREATE TABLE "attendance" (
  "id" bigserial PRIMARY KEY,
  "session_id" bigint NOT NULL,
  "officer_reg_number" text NOT NULL
);

CREATE TABLE "training_sessions" (
  "id" bigserial PRIMARY KEY,
  "course_id" bigint NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "facilitator_details" text
);

CREATE TABLE "training_courses" (
  "id" bigserial PRIMARY KEY,
  "name" text UNIQUE NOT NULL,
  "category" text NOT NULL,
  "default_credit_hours" integer NOT NULL
);

CREATE TABLE "ranks" (
  "id" bigserial PRIMARY KEY,
  "name" text UNIQUE NOT NULL
);

CREATE TABLE "postings" (
  "id" bigserial PRIMARY KEY,
  "name" text UNIQUE NOT NULL
);

CREATE TABLE "formations" (
  "id" bigserial PRIMARY KEY,
  "name" text UNIQUE NOT NULL,
  "region_id" bigint NOT NULL
);

CREATE TABLE "regions" (
  "id" bigserial PRIMARY KEY,
  "name" text UNIQUE NOT NULL
);

CREATE UNIQUE INDEX ON "attendance" ("session_id", "officer_reg_number");

COMMENT ON COLUMN "users"."role" IS 'Must be ''Administrator'', ''Content Contributor'', or ''System User''';

COMMENT ON COLUMN "officers"."sex" IS 'Must be ''Male'' or ''Female''';

COMMENT ON COLUMN "training_courses"."category" IS 'Must be ''Mandatory'' or ''Elective''';

ALTER TABLE "users" ADD FOREIGN KEY ("officer_reg_number") REFERENCES "officers" ("regulation_number");

ALTER TABLE "officers" ADD FOREIGN KEY ("rank_id") REFERENCES "ranks" ("id");

ALTER TABLE "officers" ADD FOREIGN KEY ("formation_id") REFERENCES "formations" ("id");

ALTER TABLE "officers" ADD FOREIGN KEY ("posting_id") REFERENCES "postings" ("id");

ALTER TABLE "attendance" ADD FOREIGN KEY ("session_id") REFERENCES "training_sessions" ("id");

ALTER TABLE "attendance" ADD FOREIGN KEY ("officer_reg_number") REFERENCES "officers" ("regulation_number");

ALTER TABLE "training_sessions" ADD FOREIGN KEY ("course_id") REFERENCES "training_courses" ("id");

ALTER TABLE "formations" ADD FOREIGN KEY ("region_id") REFERENCES "regions" ("id");
