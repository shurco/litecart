-- +goose Up
-- +goose StatementBegin
INSERT INTO setting (id, key, value) VALUES 
('7HkP2nYgR4sL8Qo', 'payment_webhook_url', '');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM setting WHERE id = '7HkP2nYgR4sL8Qo';
-- +goose StatementEnd
