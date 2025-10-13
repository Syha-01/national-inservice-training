CREATE TABLE facilitator_ratings (
    id SERIAL PRIMARY KEY,
    session_enrollment_id INT NOT NULL,
    facilitator_id INT NOT NULL,
    score INT NOT NULL CHECK (score >= 1 AND score <= 5),
    comment TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (session_enrollment_id) REFERENCES session_enrollment(id) ON DELETE CASCADE,
    FOREIGN KEY (facilitator_id) REFERENCES facilitators(id) ON DELETE CASCADE,
    UNIQUE (session_enrollment_id, facilitator_id)
);

