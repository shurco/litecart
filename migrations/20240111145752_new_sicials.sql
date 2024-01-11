-- +goose Up
-- +goose StatementBegin
INSERT INTO setting VALUES ('44n6ydmyjgBBr5J', 'social_youtube', '');
INSERT INTO setting VALUES ('MI7Qa7SubLQBfQy', 'social_other', '');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM setting WHERE id = 'MI7Qa7SubLQBfQy';
DELETE FROM setting WHERE id = '44n6ydmyjgBBr5J';
-- +goose StatementEnd