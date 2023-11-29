-- +goose Up
-- +goose StatementBegin
ALTER TABLE product ADD COLUMN "brief" TEXT NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE product DROP COLUMN "brief";
-- +goose StatementEnd
