-- +goose Up
-- +goose StatementBegin
INSERT INTO setting (id, key, value) VALUES ('EepD9r9nRHrIAXp', 'site_name', '');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM setting WHERE id = 'EepD9r9nRHrIAXp';
-- +goose StatementEnd
