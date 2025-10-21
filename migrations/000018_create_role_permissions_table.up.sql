CREATE TABLE IF NOT EXISTS role_permissions (
    role_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    permission_id INT NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

-- Assign all permissions to Administrator (role_id = 1)
INSERT INTO role_permissions (role_id, permission_id)
SELECT 1, id FROM permissions
ON CONFLICT DO NOTHING;

-- Assign limited permissions to Content Contributor (role_id = 2)
INSERT INTO role_permissions (role_id, permission_id)
SELECT 2, id FROM permissions
WHERE code IN (
    'officers:read', 'officers:write',
    'facilitators:read', 'facilitators:write',
    'courses:read', 'courses:write',
    'nits:read', 'nits:write'
)
ON CONFLICT DO NOTHING;

-- Assign read-only permissions to System User (role_id = 3)
INSERT INTO role_permissions (role_id, permission_id)
SELECT 3, id FROM permissions
WHERE code IN (
    'officers:read',
    'facilitators:read',
    'courses:read',
    'nits:read',
    'feedback:write'
)
ON CONFLICT DO NOTHING;
