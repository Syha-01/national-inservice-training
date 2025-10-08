-- ========= RATING TABLES =========
CREATE TABLE course_ratings (
    id SERIAL PRIMARY KEY,
    session_enrollment_id INT NOT NULL UNIQUE, -- Unique ensures one rating per enrollment
    score INT NOT NULL CHECK (score >= 1 AND score <= 5),
    comment TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (session_enrollment_id) REFERENCES session_enrollment(id) ON DELETE CASCADE
);