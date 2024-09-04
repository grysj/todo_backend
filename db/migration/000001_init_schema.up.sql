CREATE TABLE "users" (
  "user_id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
  "username" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "lists" (
  "list_id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
  "created_by" integer NOT NULL,
  "title" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "permissions" (
  "permission_id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
  "list_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "list_points" (
  "point_id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
  "list_id" integer NOT NULL,
  "point" varchar[50] NOT NULL,
  "checked" bool NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "permissions" ADD FOREIGN KEY ("list_id") REFERENCES "lists" ("list_id");

ALTER TABLE "list_points" ADD FOREIGN KEY ("list_id") REFERENCES "lists" ("list_id");

ALTER TABLE "lists" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("user_id");

ALTER TABLE "permissions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
