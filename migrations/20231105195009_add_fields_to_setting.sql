-- +goose Up
-- +goose StatementBegin

INSERT INTO setting (id, key, value) VALUES 
('h4ufqkz9nplde2g', 'stripe_webhook_id', ''),
('vjrxzwsp7qxwlb9', 'stripe_webhook_url', '');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DELETE FROM setting
WHERE (id = 'h4ufqkz9nplde2g')
   AND (id = 'vjrxzwsp7qxwlb9');

-- +goose StatementEnd
