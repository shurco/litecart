-- +goose Up
-- +goose StatementBegin
INSERT INTO setting VALUES ('8qADdU1QvVluLa7', 'paypal_active', 'false');
INSERT INTO setting VALUES ('NPnCcAhHgG26p1a', 'paypal_client_id', '');
INSERT INTO setting VALUES ('BCY0A3hbWwH0cay', 'paypal_secret_key', '');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM setting WHERE id = 'BCY0A3hbWwH0cay';
DELETE FROM setting WHERE id = 'NPnCcAhHgG26p1a';
DELETE FROM setting WHERE id = '8qADdU1QvVluLa7';
-- +goose StatementEnd