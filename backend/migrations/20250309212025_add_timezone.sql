-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ALTER COLUMN created_at
    SET DATA TYPE TIMESTAMP WITH TIME ZONE
    USING created_at AT TIME ZONE 'Europe/Moscow';

ALTER TABLE classes ALTER COLUMN start_time
    SET DATA TYPE TIMESTAMP WITH TIME ZONE
    USING start_time AT TIME ZONE 'Europe/Moscow';

ALTER TABLE classes ALTER COLUMN end_time
    SET DATA TYPE TIMESTAMP WITH TIME ZONE
    USING end_time AT TIME ZONE 'Europe/Moscow';

ALTER TABLE homeworks ALTER COLUMN due_date
    SET DATA TYPE TIMESTAMP WITH TIME ZONE
    USING due_date AT TIME ZONE 'Europe/Moscow';

ALTER TABLE homeworks ALTER COLUMN created_at
    SET DATA TYPE TIMESTAMP WITH TIME ZONE
    USING created_at AT TIME ZONE 'Europe/Moscow';

ALTER TABLE subjectnotes ALTER COLUMN created_at
    SET DATA TYPE TIMESTAMP WITH TIME ZONE
    USING created_at AT TIME ZONE 'Europe/Moscow';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users ALTER COLUMN created_at
    SET DATA TYPE TIMESTAMP WITHOUT TIME ZONE
    USING created_at AT TIME ZONE 'UTC';

ALTER TABLE classes ALTER COLUMN start_time
    SET DATA TYPE TIMESTAMP WITHOUT TIME ZONE
    USING start_time AT TIME ZONE 'UTC';

ALTER TABLE classes ALTER COLUMN end_time
    SET DATA TYPE TIMESTAMP WITHOUT TIME ZONE
    USING end_time AT TIME ZONE 'UTC';

ALTER TABLE homeworks ALTER COLUMN due_date
    SET DATA TYPE TIMESTAMP WITHOUT TIME ZONE
    USING due_date AT TIME ZONE 'UTC';

ALTER TABLE homeworks ALTER COLUMN created_at
    SET DATA TYPE TIMESTAMP WITHOUT TIME ZONE
    USING created_at AT TIME ZONE 'UTC';

ALTER TABLE subjectnotes ALTER COLUMN created_at
    SET DATA TYPE TIMESTAMP WITHOUT TIME ZONE
    USING created_at AT TIME ZONE 'UTC';

-- +goose StatementEnd
