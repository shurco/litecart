-- +goose Up
-- +goose StatementBegin
INSERT INTO  "subdomain" ("id", "name", "desc") VALUES ('62zonbnpdrzeje0', 'name', 'description');

UPDATE "setting" SET "value" = 'true' WHERE "key" = 'installed';
UPDATE "setting" SET "value" = 'site.com' WHERE "key" = 'domain';
UPDATE "setting" SET "value" = 'user@site.com' WHERE "key" = 'email';
UPDATE "setting" SET "value" = '$2a$04$k.JZqE2LV81ThRR2tgNcceNNi9Ue6P079e6mrhOPyYN/ILcwag/7G' WHERE "key" = 'password';
UPDATE "setting" SET "value" = 'd58ca30c8e5ca96695451fa27af949d9' WHERE "key" = 'jwt_secret';
UPDATE "setting" SET "value" = '48' WHERE "key" = 'jwt_secret_expire_hours';
UPDATE "setting" SET "value" = 'sk_test_000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000' WHERE "key" = 'stripe_secret_key';
UPDATE "setting" SET "value" = '' WHERE "key" = 'stripe_webhook_secret_key';

INSERT INTO  "product" ("id", "name", "desc", "url") VALUES 
('fv6c9s9cqzf36sc', 'name1', 'description1', 'url1'),
('xrtb1b919t2nuj9', 'name2', 'description2', 'url2'),
('7mweb67t8xv9pzx', 'name3', 'description3', 'url3'),
('k4pkxqhn4p0xhoc', 'name4', 'description4', 'url4'),
('2wdx6k7b3lywc2o', 'name5', 'description5', 'url5'),
('zlfpc6b17gte0ot', 'name6', 'description6', 'url6'),
('ktorsk0xj8w5zab', 'name7', 'description7', 'url7'),
('6bn739vrvfp6zaw', 'name8', 'description8', 'url8');

INSERT INTO  "product_price" ("id", "stripe_id", "product_id", "currency", "amount") VALUES 
('ohrozxu47cwnaup', 'price_1NU6CfBDuthUZlLWNncOPN53', 'fv6c9s9cqzf36sc', 'EUR', 2000),
('nvhdb52ogpzc4q6', 'price_1NU6CfBDuthUZlLWNncOPN54', 'xrtb1b919t2nuj9', 'USD', 2100),
('k0k2ulegp96s34t', 'price_1NU6CfBDuthUZlLWNncOPN55', '7mweb67t8xv9pzx', 'EUR', 2200),
('eyv2sngy7l5kc87', 'price_1NU6CfBDuthUZlLWNncOPN56', 'k4pkxqhn4p0xhoc', 'USD', 2300),
('peepx4wzppbq1v0', 'price_1NU6CfBDuthUZlLWNncOPN57', '2wdx6k7b3lywc2o', 'EUR', 2400),
('vry3k3za3t5zswq', 'price_1NU6CfBDuthUZlLWNncOPN58', 'zlfpc6b17gte0ot', 'USD', 2500),
('joyc7vc5bc2o8tj', 'price_1NU6CfBDuthUZlLWNncOPN59', 'ktorsk0xj8w5zab', 'EUR', 2600),
('0xtws4xeiet56ze', 'price_1NU6CfBDuthUZlLWNncOPN50', '6bn739vrvfp6zaw', 'USD', 2700);

INSERT INTO  "product_image" ("id", "product_id", "name", "ext") VALUES 
('dj9bae53oob0ukj', 'fv6c9s9cqzf36sc', '0f8e7e98-1639-40a3-97f6-0aac15538d88', 'webp'),
('jrnzt7lrh46xbct', 'xrtb1b919t2nuj9', '1ca0a335-7cde-4ba1-a700-138cca9ca852', 'webp'),
('5i4w6byow9n4i25', '7mweb67t8xv9pzx', '32b0115f-27aa-4a9f-aebf-c7250d1a118e', 'webp'),
('w8wwpp0vqjb4v8h', 'k4pkxqhn4p0xhoc', '76396b3e-5964-4f87-b80c-7909b2de9571', 'webp'),
('q3idmpmokavhqql', '2wdx6k7b3lywc2o', 'aa322bd6-93de-42f1-a59d-43160e67e890', 'webp'),
('ki2j3qyhboozw3c', 'zlfpc6b17gte0ot', 'd3f08f52-b290-430f-9fc7-45456fe3319f', 'webp'),
('ecbxi01t5kulwnz', 'ktorsk0xj8w5zab', 'e827e0be-aaf6-4008-aacf-da35cf47952f', 'webp'),
('tekaawq6bd9zakc', '6bn739vrvfp6zaw', 'ecd77e90-2b35-49eb-a810-a1ecf74c21a7', 'webp');

INSERT INTO  "product_metadata" ("id", "product_id", "key", "value") VALUES 
('zt6v26lmyee1gvv', 'fv6c9s9cqzf36sc', 'key1', 'value1'),('h7icc0iuvj6jr5x', 'fv6c9s9cqzf36sc', 'key2', 'value2'),('gj91gpvcr0sixz7', 'fv6c9s9cqzf36sc', 'key3', 'value3'),('3mq67jx6mhqb2tw', 'fv6c9s9cqzf36sc', 'key4', 'value4'),('oh563wybdvfu446', 'fv6c9s9cqzf36sc', 'key5', 'value5'),('fiq4n9258nizymk', 'fv6c9s9cqzf36sc', 'key6', 'value6'),
('x8mrgtrk7s05h9z', 'xrtb1b919t2nuj9', 'key1', 'value1'),('cw8l48cgommvptg', 'xrtb1b919t2nuj9', 'key2', 'value2'),('pbgfktrulpzjima', 'xrtb1b919t2nuj9', 'key3', 'value3'),('74bhmgdv1quxhs9', 'xrtb1b919t2nuj9', 'key4', 'value4'),('yn7h26yeq8uoe21', 'xrtb1b919t2nuj9', 'key5', 'value5'),('htzphqb5lsmroiu', 'xrtb1b919t2nuj9', 'key6', 'value6'),
('t9xknw0ui0uxkgg', '7mweb67t8xv9pzx', 'key1', 'value1'),('y9yct1jeghaogr5', '7mweb67t8xv9pzx', 'key2', 'value2'),('xbnpdlhr1payke4', '7mweb67t8xv9pzx', 'key3', 'value3'),('oq3qem7w531v5ru', '7mweb67t8xv9pzx', 'key4', 'value4'),('2p1oa6pkl62c88x', '7mweb67t8xv9pzx', 'key5', 'value5'),('onckbjh3yotpaja', '7mweb67t8xv9pzx', 'key6', 'value6'),
('jcxjcfbfcdph7fv', 'k4pkxqhn4p0xhoc', 'key1', 'value1'),('jrdi5g11v0285or', 'k4pkxqhn4p0xhoc', 'key2', 'value2'),('46x81tfbr9ap0e0', 'k4pkxqhn4p0xhoc', 'key3', 'value3'),('scjmtb1nyzawdvz', 'k4pkxqhn4p0xhoc', 'key4', 'value4'),('yuahtbloti0hxj7', 'k4pkxqhn4p0xhoc', 'key5', 'value5'),('norwv0i8ni0jlnl', 'k4pkxqhn4p0xhoc', 'key6', 'value6'),
('67m57ss2wetpgfy', '2wdx6k7b3lywc2o', 'key1', 'value1'),('t2f4fq7zz1v5ksw', '2wdx6k7b3lywc2o', 'key2', 'value2'),('zwkt86igspejpuh', '2wdx6k7b3lywc2o', 'key3', 'value3'),('6v2xufq5cjb92pg', '2wdx6k7b3lywc2o', 'key4', 'value4'),('9ekxw6reuzu0cb3', '2wdx6k7b3lywc2o', 'key5', 'value5'),('vmm1uc5gtpf16r1', '2wdx6k7b3lywc2o', 'key6', 'value6'),
('063ryo4sd8jf2aa', 'zlfpc6b17gte0ot', 'key1', 'value1'),('nvbgd613aeh9aqg', 'zlfpc6b17gte0ot', 'key2', 'value2'),('v8g3t61ocwt1jd2', 'zlfpc6b17gte0ot', 'key3', 'value3'),('14ebaza4zqjg2cq', 'zlfpc6b17gte0ot', 'key4', 'value4'),('g9xa9euq52m396i', 'zlfpc6b17gte0ot', 'key5', 'value5'),('5f8kxut0nlf7ikb', 'zlfpc6b17gte0ot', 'key6', 'value6'),
('yj35lgpdanrdxon', 'ktorsk0xj8w5zab', 'key1', 'value1'),('nguodxo6vnkmbh4', 'ktorsk0xj8w5zab', 'key2', 'value2'),('h3xqgc7o0ks2two', 'ktorsk0xj8w5zab', 'key3', 'value3'),('8iamc5q76q3n4t5', 'ktorsk0xj8w5zab', 'key4', 'value4'),('hjjvjg1wqgo9xoo', 'ktorsk0xj8w5zab', 'key5', 'value5'),('mvctgiy4r29ys7n', 'ktorsk0xj8w5zab', 'key6', 'value6'),
('mw26utpx9ho602r', '6bn739vrvfp6zaw', 'key1', 'value1'),('r9i4urxzkcn9jfp', '6bn739vrvfp6zaw', 'key2', 'value2'),('7qzji1llq0tj8vg', '6bn739vrvfp6zaw', 'key3', 'value3'),('mqsd5ihner42cft', '6bn739vrvfp6zaw', 'key4', 'value4'),('ke128cyd7kd5fsm', '6bn739vrvfp6zaw', 'key5', 'value5'),('l5ix9iqmlidyks4', '6bn739vrvfp6zaw', 'key6', 'value6');

INSERT INTO  "product_attribute" ("id", "product_id", "name") VALUES 
('tbluns4k4ckd6br', 'fv6c9s9cqzf36sc', 'atribute1'), ('smtduvncjx72e94', 'fv6c9s9cqzf36sc', 'atribute2'), ('evfcn53551vke38', 'fv6c9s9cqzf36sc', 'atribute3'), 
('8qihprkwzdk9zbg', 'xrtb1b919t2nuj9', 'atribute1'), ('d06xyilf4wvntch', 'xrtb1b919t2nuj9', 'atribute2'), ('t1uxvw33clal4k2', 'xrtb1b919t2nuj9', 'atribute3'), 
('zf6mciqeeo9mujo', '7mweb67t8xv9pzx', 'atribute1'), ('6va59toeiwilqnh', '7mweb67t8xv9pzx', 'atribute2'), ('brsvtc4yiwzoyza', '7mweb67t8xv9pzx', 'atribute3'), 
('t8mn0ragowrsbvg', 'k4pkxqhn4p0xhoc', 'atribute1'), ('qh003vb31bqusl8', 'k4pkxqhn4p0xhoc', 'atribute2'), ('rgt8jjteqf90ut3', 'k4pkxqhn4p0xhoc', 'atribute3'), 
('tkty3zteimcpsd0', '2wdx6k7b3lywc2o', 'atribute1'), ('wjize1u7b13qy8p', '2wdx6k7b3lywc2o', 'atribute2'), ('u2k5nn906bxa3t3', '2wdx6k7b3lywc2o', 'atribute3'), 
('992tws7udqhu2pa', 'zlfpc6b17gte0ot', 'atribute1'), ('48tqcpqk5acodi5', 'zlfpc6b17gte0ot', 'atribute2'), ('p3lcl2wg1oadkti', 'zlfpc6b17gte0ot', 'atribute3'), 
('w5f4dmiw262bi3v', 'ktorsk0xj8w5zab', 'atribute1'), ('5zlzo1mjldwa98i', 'ktorsk0xj8w5zab', 'atribute2'), ('1f7hs3869cxabfu', 'ktorsk0xj8w5zab', 'atribute3'), 
('r44p3e4m6bnxec6', '6bn739vrvfp6zaw', 'atribute1'), ('wwkzdmz1ludk22v', '6bn739vrvfp6zaw', 'atribute2'), ('2hwv6ynr1gl3678', '6bn739vrvfp6zaw', 'atribute3');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM "product_attribute";
DELETE FROM "product_metadata";
DELETE FROM "product_image";
DELETE FROM "product_price";
DELETE FROM "product";
DELETE FROM "setting";
DELETE FROM "subdomain";
-- +goose StatementEnd
