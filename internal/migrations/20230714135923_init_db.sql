-- +goose Up
-- +goose StatementBegin
CREATE TABLE "setting" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"key" varchar,
	"value" varchar
);
INSERT INTO  "setting" ("key", "value") VALUES 
('installed', 0),
('domain', ''),
('email', ''),
('password', ''),
('jwt_secret', 'secret'),
('jwt_secret_expire_hours', '48'),
('stripe_secret_key', ''),
('stripe_webhook_secret_key', '');

CREATE TABLE "session" (
	"key" varchar,
	"value" varchar,
	"expires" INTEGER
);

CREATE TABLE "subdomain" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"name" varchar,
	"desc" varchar
);

CREATE TABLE "product" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"name" varchar,
	"desc" varchar,
	"url" varchar,
	"created" TIMESTAMP DEFAULT (datetime('now')),
	"updated" TIMESTAMP
);

CREATE TABLE "product_price" (
	"stripe_id" varchar,
	"product_id" varchar,
	"currency" varchar,
	"amount" INTEGER,
	FOREIGN KEY ("product_id") REFERENCES "product"("id")
);

CREATE TABLE "product_image" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"product_id" varchar,
	"name" varchar,
	"ext" varchar,
	FOREIGN KEY ("product_id") REFERENCES "product"("id")
);

CREATE TABLE "product_metadata" (
	"id" INTEGER NOT NULL,
	"product_id" varchar,
	"key" varchar,
	"value" varchar,
	PRIMARY KEY (id),
	FOREIGN KEY ("product_id") REFERENCES "product"("id") 
);

CREATE TABLE "product_atribute" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"product_id" varchar,
	"name" varchar,
	FOREIGN KEY ("product_id") REFERENCES "product"("id")
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE "product_atribute";
DROP TABLE "product_metadata";
DROP TABLE "product_image";
DROP TABLE "product_price";
DROP TABLE "product";
DROP TABLE "subdomain";
DROP TABLE "session";
DROP TABLE "setting";
-- +goose StatementEnd
