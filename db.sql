-- ========= LOOKUP TABLES =========
-- These tables store lists of options to ensure data consistency.

CREATE TABLE regions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE formations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    region_id INT NOT NULL,
    FOREIGN KEY (region_id) REFERENCES regions(id) ON DELETE RESTRICT
);

CREATE TABLE postings (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE ranks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    abbreviation VARCHAR(20) UNIQUE
);

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE -- e.g., 'Administrator', 'Content Contributor', 'System User'
);


-- ========= CORE DATA TABLES =========

CREATE TABLE personnel (
    id SERIAL PRIMARY KEY,
    regulation_number VARCHAR(50) NOT NULL UNIQUE,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    sex VARCHAR(10) NOT NULL CHECK (sex IN ('Male', 'Female')),
    rank_id INT,
    formation_id INT,
    posting_id INT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (rank_id) REFERENCES ranks(id) ON DELETE SET NULL,
    FOREIGN KEY (formation_id) REFERENCES formations(id) ON DELETE SET NULL,
    FOREIGN KEY (posting_id) REFERENCES postings(id) ON DELETE SET NULL
);

CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    category VARCHAR(20) NOT NULL CHECK (category IN ('Mandatory', 'Elective')),
    credit_hours DECIMAL(5, 2) NOT NULL CHECK (credit_hours > 0),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE training_sessions (
    id SERIAL PRIMARY KEY,
    course_id INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    location VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
);

CREATE TABLE facilitators (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE,
    personnel_id INT UNIQUE,
    FOREIGN KEY (personnel_id) REFERENCES personnel(id) ON DELETE SET NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    personnel_id INT UNIQUE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE RESTRICT,
    FOREIGN KEY (personnel_id) REFERENCES personnel(id) ON DELETE SET NULL
);


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

CREATE TABLE session_facilitators (
    session_id INT NOT NULL,
    facilitator_id INT NOT NULL,
    PRIMARY KEY (session_id, facilitator_id),
    FOREIGN KEY (session_id) REFERENCES training_sessions(id) ON DELETE CASCADE,
    FOREIGN KEY (facilitator_id) REFERENCES facilitators(id) ON DELETE CASCADE
);


-- ========= RATING TABLES =========

CREATE TABLE course_ratings (
    id SERIAL PRIMARY KEY,
    session_enrollment_id INT NOT NULL UNIQUE, -- Unique ensures one rating per enrollment
    score INT NOT NULL CHECK (score >= 1 AND score <= 5),
    comment TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (session_enrollment_id) REFERENCES session_enrollment(id) ON DELETE CASCADE
);

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

-- ========= DATA INSERTION FOR LOOKUP TABLES =========
-- This section populates the lookup tables with the complete data from your documents.

INSERT INTO regions (name) VALUES
('Northern Region'),
('Western Region'),
('Eastern Division'),
('Southern Region');

INSERT INTO roles (name) VALUES
('Administrator'),
('Content Contributor'),
('System User');

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

INSERT INTO postings (name) VALUES
('Relief'),
('Staff Duties'),
('Station Manager'),
('Crimes Investigations Branch [CIB]'),
('Special Branch [SB]'),
('Quick Response Team [QRT]'),
('Prosecution Branch'),
('Gang Intelligence, Interdiction & Investigation [GI³]'),
('Anti-Narcotics Unit [ANU]'),
('Special Patrol Unit [SPU]'),
('Tourism Police Unit [TPU]'),
('Major Crimes Unit [MCU]'),
('Mobile Interdiction Unit [MIU]'),
('K-9 Unit'),
('Professional Standards Branch [PSB]'),
('Deputy Officer Commanding/Deputy Commander'),
('Officer Commanding/Commander'),
('Regional Commander'),
('Special Assignment'),
('Other');

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