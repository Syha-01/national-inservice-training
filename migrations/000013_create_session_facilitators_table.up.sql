CREATE TABLE session_facilitators (
    session_id INT NOT NULL,
    facilitator_id INT NOT NULL,
    PRIMARY KEY (session_id, facilitator_id),
    FOREIGN KEY (session_id) REFERENCES training_sessions(id) ON DELETE CASCADE,
    FOREIGN KEY (facilitator_id) REFERENCES facilitators(id) ON DELETE CASCADE
);