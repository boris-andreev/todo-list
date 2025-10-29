-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR NOT NULL,
    password VARCHAR
);

CREATE UNIQUE INDEX idx_users_name_unique ON users(name);

CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    name VARCHAR NOT NULL,
    description TEXT,
    status INTEGER NOT NULL,
    user_id INTEGER,
    
    CONSTRAINT fk_task_user 
        FOREIGN KEY (user_id) 
        REFERENCES users(id) 
        ON DELETE CASCADE,
    
    CONSTRAINT chk_task_status 
        CHECK (status IN (1, 2, 4, 8))
);

CREATE INDEX idx_tasks_user_id ON tasks(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
