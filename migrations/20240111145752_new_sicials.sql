-- +goose Up
-- +goose StatementBegin
INSERT OR IGNORE INTO setting VALUES ('yLR1176FQj1BQks', 'social_facebook', '');
INSERT OR IGNORE INTO setting VALUES ('rKVq63So91kMuN7', 'social_instagram', '');
INSERT OR IGNORE INTO setting VALUES ('NVv27ea47Yo7gPm', 'social_twitter', '');
INSERT OR IGNORE INTO setting VALUES ('VjdMVG7LcUL274G', 'social_dribbble', '');
INSERT OR IGNORE INTO setting VALUES ('8sz9yVDNvNBa97b', 'social_github', '');
INSERT OR IGNORE INTO setting VALUES ('CoDDXfxF4GZxq6b', 'social_youtube', '');
INSERT OR IGNORE INTO setting VALUES ('AC3of7o9pS9HdB1', 'social_other', '');

-- Fix existing smtp_port values that might be '0'
UPDATE setting SET value = '' WHERE key = 'smtp_port' AND value = '0';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM setting WHERE id = 'CoDDXfxF4GZxq6b';
DELETE FROM setting WHERE id = 'AC3of7o9pS9HdB1';
DELETE FROM setting WHERE id = '8sz9yVDNvNBa97b';
DELETE FROM setting WHERE id = 'VjdMVG7LcUL274G';
DELETE FROM setting WHERE id = 'NVv27ea47Yo7gPm';
DELETE FROM setting WHERE id = 'rKVq63So91kMuN7';
DELETE FROM setting WHERE id = 'yLR1176FQj1BQks';
-- +goose StatementEnd