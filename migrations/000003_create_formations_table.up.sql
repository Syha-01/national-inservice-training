CREATE TABLE formations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    region_id INT NOT NULL,
    FOREIGN KEY (region_id) REFERENCES regions(id) ON DELETE RESTRICT
);

-- Use subqueries to ensure the correct region_id is used for each formation.
INSERT INTO formations (name, region_id) VALUES
('Corozal Police Formation', (SELECT id FROM regions WHERE name = 'Northern Region')),
('Orange Walk Police Formation', (SELECT id FROM regions WHERE name = 'Northern Region')),
('Police Headquarters - Belmopan', (SELECT id FROM regions WHERE name = 'Western Region')),
('San Ignacio Police Formation', (SELECT id FROM regions WHERE name = 'Western Region')),
('Benque Viejo Police Formation', (SELECT id FROM regions WHERE name = 'Western Region')),
('Belmopan Police Formation', (SELECT id FROM regions WHERE name = 'Western Region')),
('Roaring Creek Police Sub-Formation', (SELECT id FROM regions WHERE name = 'Western Region')),
('Police Headquarters - Eastern Division', (SELECT id FROM regions WHERE name = 'Eastern Division')),
('Precinct 1', (SELECT id FROM regions WHERE name = 'Eastern Division')),
('Precinct 2', (SELECT id FROM regions WHERE name = 'Eastern Division')),
('Precinct 3', (SELECT id FROM regions WHERE name = 'Eastern Division')),
('Precinct 4', (SELECT id FROM regions WHERE name = 'Eastern Division')),
('Ladyville Police Sub-Formation', (SELECT id FROM regions WHERE name = 'Eastern Division')),
('Hattieville Police Sub-Formation', (SELECT id FROM regions WHERE name = 'Eastern Division')),
('Caye Caulker Police Formation', (SELECT id FROM regions WHERE name = 'Eastern Division')),
('San Pedro Police Formation', (SELECT id FROM regions WHERE name = 'Eastern Division')),
('Punta Gorda Police Formation', (SELECT id FROM regions WHERE name = 'Southern Region')),
('Intermediate Southern Formation', (SELECT id FROM regions WHERE name = 'Southern Region')),
('Placencia Police Sub-Formation', (SELECT id FROM regions WHERE name = 'Southern Region')),
('Seine Bight Police Sub-Formation', (SELECT id FROM regions WHERE name = 'Southern Region')),
('Hopkins Police Sub-Formation', (SELECT id FROM regions WHERE name = 'Southern Region')),
('Dangriga Police Formation', (SELECT id FROM regions WHERE name = 'Southern Region'));