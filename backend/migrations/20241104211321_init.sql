-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Groups (
      group_id BIGSERIAL PRIMARY KEY,
      group_name TEXT NOT NULL UNIQUE,
      course SMALLINT NOT NULL,
      ical_link TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Users (
     user_id BIGSERIAL PRIMARY KEY,
     name TEXT NOT NULL,
     surname TEXT,
     email TEXT NOT NULL UNIQUE,
     password TEXT NOT NULL,
     role int8 NOT NULL default 1 check ( role in (1, 2 ,3) ),
     group_id INTEGER not null,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

     CONSTRAINT group_fk FOREIGN KEY (group_id)REFERENCES Groups(group_id)
);


CREATE TABLE  IF NOT EXISTS Subjects (
     subject_id BIGSERIAL PRIMARY KEY,
     group_id integer not null,
     subject_name TEXT NOT NULL,


     CONSTRAINT group_fk FOREIGN KEY (group_id)REFERENCES Groups(group_id)
);

CREATE TABLE  IF NOT EXISTS Classes (
      class_id BIGSERIAL PRIMARY KEY,
      group_id INTEGER not null,
      subject_id INTEGER not null,
      start_time timestamp not null ,
      end_time timestamp not null,
      summary TEXT not null ,
      description TEXT not null ,
      class_sem_number  INTEGER NOT NULL,
      location TEXT,
      category int8 NOT NULL,

    CONSTRAINT group_fk FOREIGN KEY (group_id)REFERENCES Groups(group_id),
    CONSTRAINT subject_fk FOREIGN KEY (subject_id)REFERENCES Subjects(subject_id)
);

CREATE TABLE IF NOT EXISTS Homeworks (
    homework_id BIGSERIAL PRIMARY KEY,
    class_sem_number INTEGER,
    group_id INTEGER NOT NULL,
    subject_id INTEGER NOT NULL,
    homework_text TEXT NOT NULL,
    category int8,
    due_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT group_fk FOREIGN KEY (group_id)REFERENCES Groups(group_id),
    CONSTRAINT subject_fk FOREIGN KEY (subject_id)REFERENCES Subjects(subject_id)
);

CREATE TABLE IF NOT EXISTS HomeworkStatuses (
   id BIGSERIAL PRIMARY KEY,
   user_id INTEGER,
   homework_id INTEGER,

   CONSTRAINT user_fk FOREIGN KEY (user_id)REFERENCES Users(user_id),
   CONSTRAINT homework_fk FOREIGN KEY (homework_id)REFERENCES Homeworks(homework_id)

);

CREATE TABLE IF NOT EXISTS SubjectNotes (
    note_id BIGSERIAL PRIMARY KEY,
    subject_id INTEGER,
    group_id INTEGER,
    note_text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,



    CONSTRAINT group_fk FOREIGN KEY (group_id)REFERENCES Groups(group_id),
    CONSTRAINT subject_fk FOREIGN KEY (subject_id)REFERENCES Subjects(subject_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS Groups CASCADE;
DROP TABLE IF EXISTS Subjects CASCADE;
DROP TABLE IF EXISTS Classes CASCADE;
DROP TABLE IF EXISTS Homeworks CASCADE;
DROP TABLE IF EXISTS HomeworkStatuses CASCADE;
DROP TABLE IF EXISTS SubjectNotes CASCADE;
-- +goose StatementEnd
