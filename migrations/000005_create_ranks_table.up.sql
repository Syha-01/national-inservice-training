CREATE TABLE ranks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    abbreviation VARCHAR(20) UNIQUE
);

INSERT INTO ranks (name, abbreviation) VALUES
('Special Constable', 'SC'),
('Constable', 'PC'),
('Corporal', 'CPL'),
('Sergeant', 'SGT'),
('Inspector of Police', 'INSP'),
('Assistant Superintendent of Police', 'ASP'),
('Superintendent of Police', 'SUPT'),
('Senior Superintendent of Police', 'Sr. SUPT'),
('Assistant Commissioner of Police', 'ACP'),
('Deputy Commissioner of Police', 'DCP'),
('Commissioner of Police', 'COMPOL'),
('Not Applicable', 'N/A');