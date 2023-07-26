-- +goose Up
-- +goose StatementBegin
CREATE TABLE "setting" (
	"id" 		TEXT PRIMARY KEY NOT NULL,
	"key" 	TEXT UNIQUE NOT NULL,
	"value" TEXT DEFAULT NULL
);
INSERT INTO  "setting" ("id", "key", "value") VALUES 
('fkzjyd1p4z866mj', 'installed', 0),
('j3j2kaq67n0v9op', 'domain', ''),
('vlr2rtp82fewd1o', 'email', ''),
('zg7kdyrm9c9ivi5', 'password', ''),
('0oh908z9r28133g', 'jwt_secret', 'secret'),
('6o20io9hb2qt8c2', 'jwt_secret_expire_hours', '48'),
('2r59p9nudtykndd', 'stripe_secret_key', ''),
('hzmhlamxdwo6ca3', 'stripe_webhook_secret_key', '');

CREATE TABLE "session" (
	"key" 		TEXT UNIQUE NOT NULL,
	"value" 	TEXT DEFAULT NULL,
	"expires" INTEGER
);

CREATE TABLE "subdomain" (
	"id" 		TEXT PRIMARY KEY NOT NULL,
	"name" 	TEXT UNIQUE NOT NULL,
	"desc" 	TEXT DEFAULT NULL
);

CREATE TABLE "product" (
	"id" 			TEXT PRIMARY KEY NOT NULL,
	"name" 		TEXT NOT NULL,
	"desc" 		TEXT NOT NULL,
	"url" 		TEXT UNIQUE NOT NULL,
	"created" TIMESTAMP DEFAULT (datetime('now')),
	"updated" TIMESTAMP DEFAULT (datetime('now'))
);

CREATE TABLE "product_price" (
	"id" 		    	TEXT PRIMARY KEY NOT NULL,
	"stripe_id" 	TEXT NOT NULL,
	"product_id" 	TEXT NOT NULL,
	"currency" 		TEXT NOT NULL,
	"amount" 			INTEGER,
	FOREIGN KEY ("product_id") REFERENCES "product"("id") ON UPDATE CASCADE
);

CREATE TABLE "product_image" (
	"id" 						TEXT PRIMARY KEY NOT NULL,
	"product_id" 		TEXT NOT NULL,
	"name" 					TEXT NOT NULL,
	"ext" 					TEXT NOT NULL,
	FOREIGN KEY ("product_id") REFERENCES "product"("id") ON UPDATE CASCADE
);

CREATE TABLE "product_metadata" (
	"id" 					TEXT PRIMARY KEY NOT NULL,
	"product_id" 	TEXT NOT NULL,
	"key" 				TEXT NOT NULL,
	"value" 			TEXT NOT NULL,
	FOREIGN KEY ("product_id") REFERENCES "product"("id") ON UPDATE CASCADE
);

CREATE TABLE "product_attribute" (
	"id" 					TEXT PRIMARY KEY NOT NULL,
	"product_id" 	TEXT NOT NULL,
	"name" 				TEXT NOT NULL,
	FOREIGN KEY ("product_id") REFERENCES "product"("id") ON UPDATE CASCADE
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE "product_attribute";
DROP TABLE "product_metadata";
DROP TABLE "product_image";
DROP TABLE "product_price";
DROP TABLE "product";
DROP TABLE "subdomain";
DROP TABLE "session";
DROP TABLE "setting";
-- +goose StatementEnd
