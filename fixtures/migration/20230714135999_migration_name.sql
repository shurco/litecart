-- +goose Up
-- +goose StatementBegin
INSERT INTO subdomain (id, name, desc) VALUES ('62zonbnpdrzeje0', 'name', 'description');

UPDATE setting SET value = 'true' WHERE key = 'installed';
UPDATE setting SET value = 'site.com' WHERE key = 'domain';
UPDATE setting SET value = 'user@mail.com' WHERE key = 'email';
UPDATE setting SET value = '$2a$04$pZgnMHAxGfXeAyWSAQELkemsW4AE2D9xDwHFQ95ROOgWt53S2Tk2i' WHERE key = 'password';
UPDATE setting SET value = 'd58ca30c8e5ca96695451fa27af949d9' WHERE key = 'jwt_secret';
UPDATE setting SET value = '48' WHERE key = 'jwt_secret_expire_hours';
UPDATE setting SET value = 'sk_test_000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000' WHERE key = 'stripe_secret_key';
UPDATE setting SET value = '' WHERE key = 'stripe_webhook_secret_key';
UPDATE setting SET value = 'username' WHERE key = 'social_facebook';
UPDATE setting SET value = 'username' WHERE key = 'social_dribbble';
UPDATE setting SET value = 'localhost' WHERE key = 'smtp_host';
UPDATE setting SET value = '1025' WHERE key = 'smtp_port';
UPDATE setting SET value = 'username' WHERE key = 'smtp_username';
UPDATE setting SET value = 'password' WHERE key = 'smtp_password';
UPDATE setting SET value = 'None' WHERE key = 'smtp_encryption';

UPDATE page SET content = '<h1><strong>Lorem Ipsum content</strong></h1><p>Lorem ipsum dolor sit amet. Non culpa fugiat Aut distinctio nam placeat fuga ut eveniet natus et blanditiis ullam qui velit accusamus. Ex omnis quos <em>Sed officia ut reprehenderit autem non optio recusandae</em> ut omnis atque et totam quas cum illum laudantium. Ex doloribus ipsum qui ipsa autemNam sapiente aut dignissimos voluptas sed autem voluptatem aut quis asperiores.</p><h2><strong>33 quam voluptatum sed quia suscipit in modi quia.</strong></h2><p>Quo quos reiciendis Non laborum aut corrupti nobis qui delectus velit est vero numquam et neque aperiam. Aut quaerat autemNon deleniti et assumenda alias ut officiis laboriosam. Qui tenetur assumenda <em>Cum molestiae</em> quo nihil aliquam cum galisum minima.</p><ol><li><p>Est error saepe a placeat galisum et tempore deleniti.</p></li><li><p>Et rerum obcaecati vel alias nobis ut illo incidunt.</p></li><li><p>Non nisi itaque non atque minima.</p></li><li><p>Eum dolor fuga eos magnam omnis.</p></li><li><p>33 iure enim aut commodi magni et quis sunt quo totam enim.</p></li></ol><blockquote><p>Sit culpa impedit aut commodi cupiditate vel alias voluptatem ad enim ipsa qui saepe odit aut vitae quas!</p></blockquote><h3><strong>Et nihil repellat qui ratione doloremque aut neque tempore.</strong></h3><p>Ea molestias tenetur <em>Ea nihil ea pariatur dolores qui odio nisi</em> quo ipsum nisi! Aut quia consequatur et earum voluptassed beatae! Eum eligendi quosNon accusamus est voluptas incidunt.</p><ul><li><p>Et mollitia quos 33 quia repudiandae ab beatae dignissimos 33 pariatur architecto?</p></li><li><p>Est officia enim aut dolor sapiente et aliquam amet?</p></li><li><p>Id atque consequatur eos similique iure sed eveniet accusantium.</p></li><li><p>Aut suscipit doloribus non consequuntur perferendis et velit quam.</p></li><li><p>At animi laudantium qui mollitia earum cum cupiditate dolor.</p></li></ul>' WHERE id = 'ig9jpCixAgAu31f';
UPDATE page SET content = '<h1><strong>Lorem Ipsum content</strong></h1><p>Lorem ipsum dolor sit amet. Sit animi repellat aut quas fugitQui repudiandae. Non error officiis <strong>Ea tempora et dicta quos ut assumenda aliquam non distinctio quae</strong>! Rem nobis vero <em>Aut quidem et quaerat adipisci sed voluptatem velit</em> quo sunt placeat sed voluptatem alias sit voluptas sequi. Non aliquam perferendis et voluptatem minusEt rerum est voluptatum atque.</p><h2><strong>Et repudiandae galisum et assumenda excepturi aut eius aspernatur.</strong></h2><p>Qui quas accusamus et officia molestiaequi ipsum id laudantium itaque. Eum temporibus itaqueaut reprehenderit id numquam quisquam.</p><ul><li><p>Ea fugiat dolore est consequatur repellat est nobis rerum qui nostrum ipsa.</p></li><li><p>Est officia molestias et cupiditate eligendi ea molestiae quasi et quia doloribus.</p></li><li><p>Ad iusto quam quo nostrum error sit natus labore aut suscipit accusantium.</p></li><li><p>Et repudiandae aperiam ut facere aliquam et neque minima sit impedit dolorem.</p></li><li><p>Nam minus velit qui iure omnis quo debitis voluptatibus.</p></li><li><p>Et nihil sunt At dolores sequi id dolorem itaque ea rerum quia?</p></li></ul><blockquote><p>Et mollitia voluptatem ut fugiat odio et natus nulla ut exercitationem ipsum eum fuga distinctio.</p></blockquote><h3><strong>Qui ullam libero sed nobis dolores.</strong></h3><p>Qui modi dignissimos sed eius nostrum Est harum et cumque iste et commodi illum qui dolorum maiores est harum assumenda. Id labore sint <em>Aut quas aut repudiandae quasi aut cupiditate quia</em> aut consectetur dolor qui ducimus distinctio sit aspernatur recusandae.</p><ol><li><p>Rem dignissimos rerum ut commodi iusto cum iste sequi sit consectetur autem.</p></li><li><p>Aut saepe culpa aut deserunt libero ea vitae quisquam aut dolore voluptas.</p></li><li><p>Sit quae pariatur vel explicabo omnis At sequi inventore.</p></li><li><p>Cum magnam consequuntur qui veniam modi est reprehenderit enim.</p></li><li><p>Sed voluptatem odit sit dolores minima.</p></li></ol>' WHERE id = 'sdH0wGM54e3mZC2';
UPDATE page SET content = '<h1><strong>Lorem Ipsum content</strong></h1><p>Lorem ipsum dolor sit amet. Et delectus quis ut sequi sunt Ut vitae ab tempore voluptatem in quia nesciunt qui vero aperiam! Ut natus magnam et dolore quaeratnon ducimus. At nesciunt enim in omnis iste <em>Et alias cum dolorem magnam aut dolorem voluptatem</em>.</p><h2><strong>Est provident dolorem sit voluptas molestias et dolore tempora.</strong></h2><p>Et quia deserunt non esse aliquamhic incidunt. Eum dolore animi et nisi ipsam Hic eveniet et itaque natus ut aliquam soluta. Est optio tempora sed nihil quae <em>A quia aut soluta explicabo et molestias dolores sed modi aliquid</em> et dolores mollitia. Sit aspernatur adipisciQui iure eum amet sint?</p><ul><li><p>Id inventore quam in eaque voluptatem in natus dicta.</p></li><li><p>Qui error aspernatur qui sequi recusandae sed consequuntur consequatur et aspernatur laborum.</p></li></ul><blockquote><p>Et rerum minima quo sapiente consequatur id doloribus facere vel aperiam mollitia et dolorem beatae et totam laboriosam.</p></blockquote><h3><strong>Ex voluptates rerum est galisum accusantium.</strong></h3><p>Ut doloribus tempore aut quaerat corporisaut iusto? Et error modi <em>Sit corporis aut facere itaque et voluptas aliquid</em> et optio ullam eos molestiae eveniet! Et quisquam galisum eum quod nobis Est perspiciatis et architecto quia et harum dolore.</p>' WHERE id = 'kFCjBnL25hNTRHk';

INSERT INTO product (id, name, desc, slug, metadata, attribute, amount, digital) VALUES 
('fv6c9s9cqzf36sc', 'name1', 'description1', 'url1', '[{"key":"key1", "value":"value1"}, {"key":"key2", "value":"value2"}, {"key":"key3", "value":"value3"}, {"key":"key4", "value":"value4"}, {"key":"key5", "value":"value5"}, {"key":"key6", "value":"value6"}]', '["atribute1" ,"atribute2", "atribute3"]', 2000, 'file'),
('xrtb1b919t2nuj9', 'name2', 'description2', 'url2', '[{"key":"key1", "value":"value1"}, {"key":"key2", "value":"value2"}, {"key":"key3", "value":"value3"}, {"key":"key4", "value":"value4"}, {"key":"key5", "value":"value5"}, {"key":"key6", "value":"value6"}]', '["atribute1" ,"atribute2", "atribute3"]', 2100, 'data'),
('7mweb67t8xv9pzx', 'name3', 'description3', 'url3', '[{"key":"key1", "value":"value1"}, {"key":"key2", "value":"value2"}, {"key":"key3", "value":"value3"}, {"key":"key4", "value":"value4"}, {"key":"key5", "value":"value5"}, {"key":"key6", "value":"value6"}]', '["atribute1" ,"atribute2", "atribute3"]', 2200, 'file'),
('k4pkxqhn4p0xhoc', 'name4', 'description4', 'url4', '[{"key":"key1", "value":"value1"}, {"key":"key2", "value":"value2"}, {"key":"key3", "value":"value3"}, {"key":"key4", "value":"value4"}, {"key":"key5", "value":"value5"}, {"key":"key6", "value":"value6"}]', '["atribute1" ,"atribute2", "atribute3"]', 2300, 'data'),
('2wdx6k7b3lywc2o', 'name5', 'description5', 'url5', '[{"key":"key1", "value":"value1"}, {"key":"key2", "value":"value2"}, {"key":"key3", "value":"value3"}, {"key":"key4", "value":"value4"}, {"key":"key5", "value":"value5"}, {"key":"key6", "value":"value6"}]', '["atribute1" ,"atribute2", "atribute3"]', 2400, 'file'),
('zlfpc6b17gte0ot', 'name6', 'description6', 'url6', '[{"key":"key1", "value":"value1"}, {"key":"key2", "value":"value2"}, {"key":"key3", "value":"value3"}, {"key":"key4", "value":"value4"}, {"key":"key5", "value":"value5"}, {"key":"key6", "value":"value6"}]', '["atribute1" ,"atribute2", "atribute3"]', 2500, 'data'),
('ktorsk0xj8w5zab', 'name7', 'description7', 'url7', '[{"key":"key1", "value":"value1"}, {"key":"key2", "value":"value2"}, {"key":"key3", "value":"value3"}, {"key":"key4", "value":"value4"}, {"key":"key5", "value":"value5"}, {"key":"key6", "value":"value6"}]', '["atribute1" ,"atribute2", "atribute3"]', 2600, 'file'),
('6bn739vrvfp6zaw', 'name8', 'description8', 'url8', '[{"key":"key1", "value":"value1"}, {"key":"key2", "value":"value2"}, {"key":"key3", "value":"value3"}, {"key":"key4", "value":"value4"}, {"key":"key5", "value":"value5"}, {"key":"key6", "value":"value6"}]', '["atribute1" ,"atribute2", "atribute3"]', 2700, 'data');
UPDATE product SET active = 0 WHERE id = 'zlfpc6b17gte0ot';

INSERT INTO digital_file (id, product_id, name, ext, orig_name) VALUES 
('QLYUrC7p3XuXRFC', 'fv6c9s9cqzf36sc', '1ca0a335-7cde-4ba1-a700-138cca9ca852', 'png', 'secret_image_1.png'),
('Z4tUs9CMeLGvJDr', 'fv6c9s9cqzf36sc', 'ff0b48d1-0a75-4d67-a0ac-e6243cfd6cec', 'png', 'secret_image_2.png');

INSERT INTO digital_data (id, product_id, content) VALUES 
('c0gog7a4zrwW4Vf', 'xrtb1b919t2nuj9', '0f8e7e98-1639-40a3-97f6-0aac15538d88'),
('utkHD3W9LudJomc', 'xrtb1b919t2nuj9', '1ca0a335-7cde-4ba1-a700-138cca9ca852'),
('CujBxoPx97C2GQf', 'xrtb1b919t2nuj9', 'ff0b48d1-0a75-4d67-a0ac-e6243cfd6cec'),
('qYUn8gM7s9KLtJ7', 'xrtb1b919t2nuj9', '165d4e99-ba1b-4d03-ba6c-3abfab65830e'),
('27tfwRp48wZ5jmt', 'xrtb1b919t2nuj9', '746becd7-59dc-4a00-aca9-e86e7290a54f');
UPDATE digital_data SET cart_id = 'iodz4ibf5h5zmov' WHERE id = 'CujBxoPx97C2GQf';

INSERT INTO product_image (id, product_id, name, ext, orig_name) VALUES 
('dj9bae53oob0ukj', 'fv6c9s9cqzf36sc', '0f8e7e98-1639-40a3-97f6-0aac15538d88', 'png', 'example_image_1.png'),
('jrnzt7lrh46xbct', 'xrtb1b919t2nuj9', '1ca0a335-7cde-4ba1-a700-138cca9ca852', 'png', 'example_image_2.png'),
('6bk6p3yvbn1lvw2', 'xrtb1b919t2nuj9', 'ff0b48d1-0a75-4d67-a0ac-e6243cfd6cec', 'png', 'example_image_3.png'),
('gshvsfj9i8d0qvn', 'xrtb1b919t2nuj9', '165d4e99-ba1b-4d03-ba6c-3abfab65830e', 'png', 'example_image_4.png'),
('325gfi7kgtik9kn', 'xrtb1b919t2nuj9', '746becd7-59dc-4a00-aca9-e86e7290a54f', 'png', 'example_image_5.png'),
('91gu06z0anxqpks', '7mweb67t8xv9pzx', 'f9b85683-25ee-40cd-b398-9c990d90b80b', 'png', 'example_image_6.png'),
('53i22q56oooiiit', '7mweb67t8xv9pzx', 'cae6dcde-9813-4ab2-9436-7bd4b2ccea36', 'png', 'example_image_7.png'),
('5i4w6byow9n4i25', '7mweb67t8xv9pzx', '32b0115f-27aa-4a9f-aebf-c7250d1a118e', 'png', 'example_image_8.png'),
('w8wwpp0vqjb4v8h', 'k4pkxqhn4p0xhoc', '76396b3e-5964-4f87-b80c-7909b2de9571', 'png', 'example_image_9.png'),
('q3idmpmokavhqql', '2wdx6k7b3lywc2o', 'aa322bd6-93de-42f1-a59d-43160e67e890', 'png', 'example_image_10.png'),
('ki2j3qyhboozw3c', 'zlfpc6b17gte0ot', 'd3f08f52-b290-430f-9fc7-45456fe3319f', 'png', 'example_image_11.png'),
('ecbxi01t5kulwnz', 'ktorsk0xj8w5zab', 'e827e0be-aaf6-4008-aacf-da35cf47952f', 'png', 'example_image_12.png'),
('tekaawq6bd9zakc', '6bn739vrvfp6zaw', 'ecd77e90-2b35-49eb-a810-a1ecf74c21a7', 'png', 'example_image_13.png');

INSERT INTO cart (id, email, name, amount_total, currency, payment_id, payment_status, cart, updated) VALUES 
('efzs4xayz43f226', NULL, NULL, 4200, 'usd', NULL, 'cancel', '[{"id":"7mweb67t8xv9pzx","quantity":1},{"id":"fv6c9s9cqzf36sc","quantity":1}]', NULL),
('iodz4ibf5h5zmov', 'user@gmail.com', 'User Name', 6300, 'usd', 'pi_3NpAmuBDuthUZlLW11fS8GrB', 'paid', '[{"id":"fv6c9s9cqzf36sc","quantity":1},{"id":"xrtb1b919t2nuj9","quantity":1},{"id":"7mweb67t8xv9pzx","quantity":1}]', datetime('now'));
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DELETE FROM cart;
DELETE FROM digital_file;
DELETE FROM digital_data;
DELETE FROM product_image;
DELETE FROM product;
UPDATE page SET content = '' WHERE id = 'ig9jpCixAgAu31f';
UPDATE page SET content = '' WHERE id = 'sdH0wGM54e3mZC2';
UPDATE page SET content = '' WHERE id = 'kFCjBnL25hNTRHk';
UPDATE setting SET value = 'false' WHERE key = 'installed';
UPDATE setting SET value = '' WHERE key = 'domain';
UPDATE setting SET value = '' WHERE key = 'email';
UPDATE setting SET value = '' WHERE key = 'password';
UPDATE setting SET value = 'secret' WHERE key = 'jwt_secret';
UPDATE setting SET value = '48' WHERE key = 'jwt_secret_expire_hours';
UPDATE setting SET value = '' WHERE key = 'stripe_secret_key';
UPDATE setting SET value = '' WHERE key = 'stripe_webhook_secret_key';
UPDATE setting SET value = '' WHERE key = 'social_facebook';
UPDATE setting SET value = '' WHERE key = 'social_dribbble';
UPDATE setting SET value = '' WHERE key = 'smtp_host';
UPDATE setting SET value = '' WHERE key = 'smtp_port';
UPDATE setting SET value = '' WHERE key = 'smtp_username';
UPDATE setting SET value = '' WHERE key = 'smtp_password';
UPDATE setting SET value = '' WHERE key = 'smtp_encryption';
DELETE FROM subdomain;
-- +goose StatementEnd
