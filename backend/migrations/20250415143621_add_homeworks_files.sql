-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS HomeworksFiles (
  file_id BIGSERIAL PRIMARY KEY,
  homework_id integer NOT NULL,
  group_id integer NOT NULL,
  file_name TEXT NOT NULL,
  file_url TEXT NOT NULL,
  key TEXT NOT NULL,
  uploaded_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT homework_fk FOREIGN KEY (homework_id) REFERENCES homeworks(homework_id),
  CONSTRAINT group_fk FOREIGN KEY (group_id) REFERENCES groups(group_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS HomeworksFiles CASCADE;
-- +goose StatementEnd
