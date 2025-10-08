CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE -- e.g., 'Administrator', 'Content Contributor', 'System User'
);

INSERT INTO roles (name) VALUES
('Administrator'),
('Content Contributor'),
('System User');
