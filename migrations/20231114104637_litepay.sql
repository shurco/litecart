-- +goose Up
-- +goose StatementBegin
DELETE FROM setting WHERE "id" = 'hzmhlamxdwo6ca3';
UPDATE setting SET key = 'webhook_url' WHERE id = '7HkP2nYgR4sL8Qo';
UPDATE setting SET value = '{"subject":"Thank you for your purchase!","text":"Dear Customer,\n\nThank you for choosing our travel guide! We appreciate your support and hope that it will enhance your travel experience. \n\nAs per your order, here is a summary of your purchases:\n\n{{.Purchases}}\n\nWe believe that our guides will provide you with valuable insights and help you explore the places mentioned in them. \n\nIf you encounter any issues with downloading or accessing the files, please feel free to contact us through our feedback form {{.Admin_Email}}. We`ll be more than happy to assist you.\n\nOnce again, thank you for your purchase. We wish you an incredible journey filled with unforgettable moments!\n\nBest regards,","html":""}' WHERE id = 'CoDDXfxF4GZxq6b';
INSERT INTO setting VALUES ('Q6QTrA77q2QZkXh', 'stripe_active', 'false');
INSERT INTO setting VALUES ('cK2G9CgKwMaZ3Xx', 'spectrocoin_active', 'false');
INSERT INTO setting VALUES ('FxfwJFu42oDiE8v', 'spectrocoin_merchant_id', '');
INSERT INTO setting VALUES ('UXMw8hijBLb7K59', 'spectrocoin_project_id', '');
INSERT INTO setting VALUES ('cPw4L82dsCLFKxr', 'spectrocoin_private_key', '');
INSERT INTO setting VALUES ('pF1gvf3AAiTL8uN', 'mail_letter_payment', '{"subject":"New payment transaction","text":"Hello,\nWe would like to inform you about the registration of a new payment transaction on the [{{.Site_Name}}] website.\n\nTransaction information:\nAmount payment: {{.Amount_Payment}}\nPayment url: {{.Payment_URL}}\n\nBest regards,","html":""}');
ALTER TABLE cart DROP COLUMN "name";
ALTER TABLE cart ADD COLUMN "payment_system" TEXT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE cart DROP COLUMN "payment_system";
ALTER TABLE cart ADD COLUMN "name" TEXT DEFAULT NULL;
DELETE FROM setting WHERE id = 'pF1gvf3AAiTL8uN';
DELETE FROM setting WHERE id = 'cPw4L82dsCLFKxr';
DELETE FROM setting WHERE id = 'UXMw8hijBLb7K59';
DELETE FROM setting WHERE id = 'FxfwJFu42oDiE8v';
DELETE FROM setting WHERE id = 'cK2G9CgKwMaZ3Xx';
DELETE FROM setting WHERE id = 'Q6QTrA77q2QZkXh';
UPDATE setting SET value = '{"subject":"Thank you for your purchase!","text":"Dear {{.Customer_Name}},\n\nThank you for choosing our travel guide! We appreciate your support and hope that it will enhance your travel experience. \n\nAs per your order, here is a summary of your purchases:\n\n{{.Purchases}}\n\nWe believe that our guides will provide you with valuable insights and help you explore the places mentioned in them. \n\nIf you encounter any issues with downloading or accessing the files, please feel free to contact us through our feedback form {{.Admin_Email}}. We`ll be more than happy to assist you.\n\nOnce again, thank you for your purchase. We wish you an incredible journey filled with unforgettable moments!\n\nBest regards,","html":""}' WHERE id = 'CoDDXfxF4GZxq6b';
UPDATE setting SET key = 'payment_webhook_url' WHERE id = '7HkP2nYgR4sL8Qo';
INSERT INTO setting VALUES ('hzmhlamxdwo6ca3', 'stripe_webhook_secret_key', '');
-- +goose StatementEnd
