-- +goose Up
-- +goose StatementBegin
INSERT INTO setting VALUES ('nbyPHZ5roJt5Z2v', 'mail_sender_name', '');
INSERT INTO setting VALUES ('T3kP16of88quy3x', 'mail_sender_email', '');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM setting WHERE id = 'T3kP16of88quy3x';
DELETE FROM setting WHERE id = 'nbyPHZ5roJt5Z2v';
-- +goose StatementEnd