-- +goose Up
-- +goose StatementBegin
-- Fix existing smtp_port values that might be '0' to empty string
UPDATE setting SET value = '' WHERE key = 'smtp_port' AND value = '0';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- No down migration needed for this fix
-- +goose StatementEnd 