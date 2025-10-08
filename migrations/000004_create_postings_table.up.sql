CREATE TABLE postings (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

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
