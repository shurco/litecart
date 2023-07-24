-- +goose Up
-- +goose StatementBegin
INSERT INTO  "subdomain" ("name", "desc") VALUES ('name', 'description');

UPDATE "setting" SET "value" = 'true' WHERE "key" = 'installed';
UPDATE "setting" SET "value" = 'site.com' WHERE "key" = 'domain';
UPDATE "setting" SET "value" = 'user@site.com' WHERE "key" = 'email';
UPDATE "setting" SET "value" = '$2a$04$k.JZqE2LV81ThRR2tgNcceNNi9Ue6P079e6mrhOPyYN/ILcwag/7G' WHERE "key" = 'password';
UPDATE "setting" SET "value" = 'd58ca30c8e5ca96695451fa27af949d9' WHERE "key" = 'jwt_secret';
UPDATE "setting" SET "value" = '48' WHERE "key" = 'jwt_secret_expire_hours';
UPDATE "setting" SET "value" = 'sk_test_000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000' WHERE "key" = 'stripe_secret_key';
UPDATE "setting" SET "value" = '' WHERE "key" = 'stripe_webhook_secret_key';

INSERT INTO  "product" ("name", "desc", "url") VALUES 
('name1', 'description1', 'url1'),
('name2', 'description2', 'url2'),
('name3', 'description3', 'url3'),
('name4', 'description4', 'url4'),
('name5', 'description5', 'url5'),
('name6', 'description6', 'url6'),
('name7', 'description7', 'url7'),
('name8', 'description8', 'url8');

INSERT INTO  "product_price" ("stripe_id", "product_id", "currency", "amount") VALUES 
('price_1NU6CfBDuthUZlLWNncOPN53', 1, 'EUR', 2000),
('price_1NU6CfBDuthUZlLWNncOPN54', 2, 'USD', 2100),
('price_1NU6CfBDuthUZlLWNncOPN55', 3, 'EUR', 2200),
('price_1NU6CfBDuthUZlLWNncOPN56', 4, 'USD', 2300),
('price_1NU6CfBDuthUZlLWNncOPN57', 5, 'EUR', 2400),
('price_1NU6CfBDuthUZlLWNncOPN58', 6, 'USD', 2500),
('price_1NU6CfBDuthUZlLWNncOPN59', 7, 'EUR', 2600),
('price_1NU6CfBDuthUZlLWNncOPN50', 8, 'USD', 2700);


INSERT INTO  "product_image" ("product_id", "name", "ext") VALUES 
(1, '0f8e7e98-1639-40a3-97f6-0aac15538d88', 'webp'),
(2, '1ca0a335-7cde-4ba1-a700-138cca9ca852', 'webp'),
(3, '32b0115f-27aa-4a9f-aebf-c7250d1a118e', 'webp'),
(4, '76396b3e-5964-4f87-b80c-7909b2de9571', 'webp'),
(5, 'aa322bd6-93de-42f1-a59d-43160e67e890', 'webp'),
(6, 'd3f08f52-b290-430f-9fc7-45456fe3319f', 'webp'),
(7, 'e827e0be-aaf6-4008-aacf-da35cf47952f', 'webp'),
(8, 'ecd77e90-2b35-49eb-a810-a1ecf74c21a7', 'webp');


INSERT INTO  "product_metadata" ("product_id", "key", "value") VALUES 
(1, 'key1', 'value1'),(1, 'key2', 'value2'),(1, 'key3', 'value3'),(1, 'key4', 'value4'),(1, 'key5', 'value5'),(1, 'key6', 'value6'),
(2, 'key1', 'value1'),(2, 'key2', 'value2'),(2, 'key3', 'value3'),(2, 'key4', 'value4'),(2, 'key5', 'value5'),(2, 'key6', 'value6'),
(3, 'key1', 'value1'),(3, 'key2', 'value2'),(3, 'key3', 'value3'),(3, 'key4', 'value4'),(3, 'key5', 'value5'),(3, 'key6', 'value6'),
(4, 'key1', 'value1'),(4, 'key2', 'value2'),(4, 'key3', 'value3'),(4, 'key4', 'value4'),(4, 'key5', 'value5'),(4, 'key6', 'value6'),
(5, 'key1', 'value1'),(5, 'key2', 'value2'),(5, 'key3', 'value3'),(5, 'key4', 'value4'),(5, 'key5', 'value5'),(5, 'key6', 'value6'),
(6, 'key1', 'value1'),(6, 'key2', 'value2'),(6, 'key3', 'value3'),(6, 'key4', 'value4'),(6, 'key5', 'value5'),(6, 'key6', 'value6'),
(7, 'key1', 'value1'),(7, 'key2', 'value2'),(7, 'key3', 'value3'),(7, 'key4', 'value4'),(7, 'key5', 'value5'),(7, 'key6', 'value6'),
(8, 'key1', 'value1'),(8, 'key2', 'value2'),(8, 'key3', 'value3'),(8, 'key4', 'value4'),(8, 'key5', 'value5'),(8, 'key6', 'value6');

INSERT INTO  "product_atribute" ("product_id", "name") VALUES 
(1, 'atribute1'), (1, 'atribute2'), (1, 'atribute3'), 
(2, 'atribute1'), (2, 'atribute2'), (2, 'atribute3'), 
(3, 'atribute1'), (3, 'atribute2'), (3, 'atribute3'), 
(4, 'atribute1'), (4, 'atribute2'), (4, 'atribute3'), 
(5, 'atribute1'), (5, 'atribute2'), (5, 'atribute3'), 
(6, 'atribute1'), (6, 'atribute2'), (6, 'atribute3'), 
(7, 'atribute1'), (7, 'atribute2'), (7, 'atribute3'), 
(8, 'atribute1'), (8, 'atribute2'), (8, 'atribute3');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM "product_atribute";
DELETE FROM "product_metadata";
DELETE FROM "product_image";
DELETE FROM "product_price";
DELETE FROM "product";
DELETE FROM "setting";
DELETE FROM "subdomain";
-- +goose StatementEnd
