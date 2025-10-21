CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL UNIQUE,
    description TEXT
);

-- Insert common permissions
INSERT INTO permissions (code, description) VALUES
    ('officers:read', 'View officers'),
    ('officers:write', 'Create and update officers'),
    ('officers:delete', 'Delete officers'),
    ('facilitators:read', 'View facilitators'),
    ('facilitators:write', 'Create and update facilitators'),
    ('facilitators:delete', 'Delete facilitators'),
    ('courses:read', 'View courses'),
    ('courses:write', 'Create and update courses'),
    ('courses:delete', 'Delete courses'),
    ('nits:read', 'View NITs'),
    ('nits:write', 'Create and update NITs'),
    ('nits:delete', 'Delete NITs'),
    ('feedback:read', 'Read feedback'),
    ('feedback:write', 'Write feedback'),
    ('admin:all', 'Full administrative access')
ON CONFLICT (code) DO NOTHING;
