-- +goose Up
-- +goose StatementBegin
INSERT INTO subdomain (id, name, desc) VALUES ('62zonbnpdrzeje0', 'name', 'description');

UPDATE setting SET value = 'true' WHERE key = 'installed';
UPDATE setting SET value = 'site.com' WHERE key = 'domain';
UPDATE setting SET value = 'user@site.com' WHERE key = 'email';
UPDATE setting SET value = '$2a$04$k.JZqE2LV81ThRR2tgNcceNNi9Ue6P079e6mrhOPyYN/ILcwag/7G' WHERE key = 'password';
UPDATE setting SET value = 'd58ca30c8e5ca96695451fa27af949d9' WHERE key = 'jwt_secret';
UPDATE setting SET value = '48' WHERE key = 'jwt_secret_expire_hours';
UPDATE setting SET value = 'sk_test_000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000' WHERE key = 'stripe_secret_key';
UPDATE setting SET value = '' WHERE key = 'stripe_webhook_secret_key';

INSERT INTO product (id, name, desc, url, metadata, attribute, stripe) VALUES 
('fv6c9s9cqzf36sc', 'name1', 'description1', 'url1', '{"key1":"value1", "key2":"value2", "key3":"value3", "key4":"value4", "key5":"value5", "key6":"value6"}', '["atribute1" ,"atribute2", "atribute3"]', '{"product":{"id":"prod_OGdTsykDrQSkmA", "valid": 0},"price":{"id":"", "currency":"EUR", "amount":"2000"}}'),
('xrtb1b919t2nuj9', 'name2', 'description2', 'url2', '{"key1":"value1", "key2":"value2", "key3":"value3", "key4":"value4", "key5":"value5", "key6":"value6"}', '["atribute1" ,"atribute2", "atribute3"]', '{"product":{"id":"", "valid": 0},"price":{"id":"", "currency":"USD", "amount":"2000"}}'),
('7mweb67t8xv9pzx', 'name3', 'description3', 'url3', '{"key1":"value1", "key2":"value2", "key3":"value3", "key4":"value4", "key5":"value5", "key6":"value6"}', '["atribute1" ,"atribute2", "atribute3"]', '{"product":{"id":"prod_OGdTsykDrQSkmb", "valid": 0},"price":{"id":"", "currency":"EUR", "amount":"2200"}}'),
('k4pkxqhn4p0xhoc', 'name4', 'description4', 'url4', '{"key1":"value1", "key2":"value2", "key3":"value3", "key4":"value4", "key5":"value5", "key6":"value6"}', '["atribute1" ,"atribute2", "atribute3"]', '{"product":{"id":"", "valid": 0},"price":{"id":"", "currency":"USD", "amount":"2300"}}'),
('2wdx6k7b3lywc2o', 'name5', 'description5', 'url5', '{"key1":"value1", "key2":"value2", "key3":"value3", "key4":"value4", "key5":"value5", "key6":"value6"}', '["atribute1" ,"atribute2", "atribute3"]', '{"product":{"id":"", "valid": 0},"price":{"id":"", "currency":"EUR", "amount":"2400"}}'),
('zlfpc6b17gte0ot', 'name6', 'description6', 'url6', '{"key1":"value1", "key2":"value2", "key3":"value3", "key4":"value4", "key5":"value5", "key6":"value6"}', '["atribute1" ,"atribute2", "atribute3"]', '{"product":{"id":"", "valid": 0},"price":{"id":"", "currency":"USD", "amount":"2500"}}'),
('ktorsk0xj8w5zab', 'name7', 'description7', 'url7', '{"key1":"value1", "key2":"value2", "key3":"value3", "key4":"value4", "key5":"value5", "key6":"value6"}', '["atribute1" ,"atribute2", "atribute3"]', '{"product":{"id":"", "valid": 0},"price":{"id":"", "currency":"EUR", "amount":"2600"}}'),
('6bn739vrvfp6zaw', 'name8', 'description8', 'url8', '{"key1":"value1", "key2":"value2", "key3":"value3", "key4":"value4", "key5":"value5", "key6":"value6"}', '["atribute1" ,"atribute2", "atribute3"]', '{"product":{"id":"", "valid": 0},"price":{"id":"", "currency":"USD", "amount":"2700"}}');

UPDATE product SET active = 0 WHERE id = 'zlfpc6b17gte0ot';

INSERT INTO product_image (id, product_id, name, ext) VALUES 
('dj9bae53oob0ukj', 'fv6c9s9cqzf36sc', '0f8e7e98-1639-40a3-97f6-0aac15538d88', 'png'),
('jrnzt7lrh46xbct', 'xrtb1b919t2nuj9', '1ca0a335-7cde-4ba1-a700-138cca9ca852', 'png'),
('6bk6p3yvbn1lvw2', 'xrtb1b919t2nuj9', 'ff0b48d1-0a75-4d67-a0ac-e6243cfd6cec', 'png'),
('gshvsfj9i8d0qvn', 'xrtb1b919t2nuj9', '165d4e99-ba1b-4d03-ba6c-3abfab65830e', 'png'),
('325gfi7kgtik9kn', 'xrtb1b919t2nuj9', '746becd7-59dc-4a00-aca9-e86e7290a54f', 'png'),
('91gu06z0anxqpks', '7mweb67t8xv9pzx', 'f9b85683-25ee-40cd-b398-9c990d90b80b', 'png'),
('53i22q56oooiiit', '7mweb67t8xv9pzx', 'cae6dcde-9813-4ab2-9436-7bd4b2ccea36', 'png'),
('5i4w6byow9n4i25', '7mweb67t8xv9pzx', '32b0115f-27aa-4a9f-aebf-c7250d1a118e', 'png'),
('w8wwpp0vqjb4v8h', 'k4pkxqhn4p0xhoc', '76396b3e-5964-4f87-b80c-7909b2de9571', 'png'),
('q3idmpmokavhqql', '2wdx6k7b3lywc2o', 'aa322bd6-93de-42f1-a59d-43160e67e890', 'png'),
('ki2j3qyhboozw3c', 'zlfpc6b17gte0ot', 'd3f08f52-b290-430f-9fc7-45456fe3319f', 'png'),
('ecbxi01t5kulwnz', 'ktorsk0xj8w5zab', 'e827e0be-aaf6-4008-aacf-da35cf47952f', 'png'),
('tekaawq6bd9zakc', '6bn739vrvfp6zaw', 'ecd77e90-2b35-49eb-a810-a1ecf74c21a7', 'png');
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DELETE FROM product_image;
DELETE FROM product;
DELETE FROM setting;
DELETE FROM subdomain;
-- +goose StatementEnd
