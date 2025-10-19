CREATE TABLE facilitators (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE,
    personnel_id INT UNIQUE,
    version INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY (personnel_id) REFERENCES personnel(id) ON DELETE SET NULL
);
