-- +goose Up
-- +goose StatementBegin
INSERT INTO setting VALUES ('EepD9r9nRHrIAXp', 'site_name', '');
ALTER TABLE product ADD COLUMN "seo" JSON DEFAULT '{}' NOT NULL;
ALTER TABLE page ADD COLUMN "seo" JSON DEFAULT '{}' NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE page DROP COLUMN "seo";
ALTER TABLE product DROP COLUMN "seo";
DELETE FROM setting WHERE id = 'EepD9r9nRHrIAXp';
-- +goose StatementEnd
