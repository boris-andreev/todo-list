-- +goose Up
-- +goose StatementBegin
INSERT INTO users (name, password)
VALUES 
    ('Ivan', '$2a$10$Gv1PNvmgCUC2gAtR0M9B5OXQoKtURKEGpsJEhQTHmm.z.5rzsMg4y'),
    ('Petr', '$2a$10$pwfu2YlWb3DqqY08yBV7XuPnuW4qTaf2YRt5jf0q7P.SUAl3corvi');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE name IN ('Ivan', 'Petr');
-- +goose StatementEnd
