-- +goose Up
-- +goose StatementBegin
INSERT INTO setting VALUES ('dUmMyPaYm3ntAcT', 'dummy_active', 'false');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM setting WHERE id = 'dUmMyPaYm3ntAcT';
-- +goose StatementEnd
