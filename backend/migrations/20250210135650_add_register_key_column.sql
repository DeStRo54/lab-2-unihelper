-- +goose Up
-- +goose StatementBegin
ALTER TABLE groups ADD COLUMN register_key VARCHAR(10) NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE groups DROP COLUMN register_key;
-- +goose StatementEnd
