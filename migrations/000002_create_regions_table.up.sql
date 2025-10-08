CREATE TABLE regions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

INSERT INTO regions (name) VALUES
('Northern Region'),
('Western Region'),
('Eastern Division'),
('Southern Region');