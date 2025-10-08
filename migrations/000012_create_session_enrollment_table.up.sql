-- ========= JUNCTION TABLES (Many-to-Many Relationships) =========
CREATE TABLE session_enrollment (
    id SERIAL PRIMARY KEY,
    personnel_id INT NOT NULL,
    session_id INT NOT NULL,
    completion_date DATE,
    status VARCHAR(50) DEFAULT 'Enrolled' CHECK (status IN ('Enrolled', 'Completed', 'Failed', 'Withdrew')),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (personnel_id) REFERENCES personnel(id) ON DELETE CASCADE,
    FOREIGN KEY (session_id) REFERENCES training_sessions(id) ON DELETE CASCADE,
    UNIQUE(personnel_id, session_id)
);