-- ===================================================================
--                  TEST DATA INSERTION SCRIPT
-- ===================================================================
-- Purpose: To populate the database with a variety of sample data
--          to test all relationships and constraints.
-- Prerequisite: All tables have been created and lookup tables
--               (regions, ranks, postings, formations, roles)
--               have been populated.
-- ===================================================================

-- Clear previous test data to run this script multiple times
-- NOTE: Use with caution. This deletes data.
-- DELETE FROM course_ratings;
-- DELETE FROM facilitator_ratings;
-- DELETE FROM session_enrollment;
-- DELETE FROM session_facilitators;
-- DELETE FROM users;
-- DELETE FROM facilitators;
-- DELETE FROM personnel;
-- DELETE FROM training_sessions;
-- DELETE FROM courses;


-- ========= 1. INSERT Personnel =========
-- We create a few police officers with different roles.
-- Using subqueries like (SELECT id FROM...) makes the script robust.

INSERT INTO personnel (regulation_number, first_name, last_name, sex, rank_id, formation_id, posting_id) VALUES
('101', 'Kenroy', 'Elijo', 'Male',
    (SELECT id FROM ranks WHERE name = 'Inspector of Police'),
    (SELECT id FROM formations WHERE name = 'Police Headquarters - Belmopan'),
    (SELECT id FROM postings WHERE name = 'Staff Duties')),
('202', 'Jane', 'Smith', 'Female',
    (SELECT id FROM ranks WHERE name = 'Constable'),
    (SELECT id FROM formations WHERE name = 'Precinct 2'),
    (SELECT id FROM postings WHERE name = 'Relief')),
('303', 'John', 'Doe', 'Male',
    (SELECT id FROM ranks WHERE name = 'Sergeant'),
    (SELECT id FROM formations WHERE name = 'San Ignacio Police Formation'),
    (SELECT id FROM postings WHERE name = 'Crimes Investigations Branch [CIB]')),
('404', 'Maria', 'Garcia', 'Female',
    (SELECT id FROM ranks WHERE name = 'Corporal'),
    (SELECT id FROM formations WHERE name = 'San Pedro Police Formation'),
    (SELECT id FROM postings WHERE name = 'Tourism Police Unit [TPU]'));

-- ========= 2. INSERT Facilitators =========
-- One facilitator IS a police officer (Kenroy Elijo)
-- One facilitator IS NOT a police officer (Dr. Annabelle Crane), so personnel_id is NULL.

INSERT INTO facilitators (first_name, last_name, email, personnel_id) VALUES
('Kenroy', 'Elijo', 'kelijo@police.bz', (SELECT id FROM personnel WHERE regulation_number = '101')),
('Annabelle', 'Crane', 'acrane.consulting@gmail.com', NULL);


-- ========= 3. INSERT Users =========
-- Create users with different roles.
-- One admin user is not an officer.
-- Kenroy Elijo is a 'Content Contributor' linked to his personnel record.

INSERT INTO users (email, password_hash, role_id, personnel_id) VALUES
('admin@police.bz', 'some_secure_password_hash_here', (SELECT id FROM roles WHERE name = 'Administrator'), NULL),
('kelijo@police.bz', 'another_secure_password_hash', (SELECT id FROM roles WHERE name = 'Content Contributor'), (SELECT id FROM personnel WHERE regulation_number = '101'));


-- ========= 4. INSERT Courses =========
-- Create one mandatory and two elective courses.

INSERT INTO courses (title, description, category, credit_hours) VALUES
('Annual Firearms Requalification', 'Mandatory handgun and shotgun requalification course.', 'Mandatory', 8.00),
('Advanced Report Writing', 'A course on writing clear, concise, and legally defensible reports.', 'Elective', 16.00),
('Community Policing Strategies', 'Workshop on modern community engagement tactics.', 'Elective', 12.50);


-- ========= 5. INSERT Training Sessions =========
-- Schedule specific instances of the courses above.

INSERT INTO training_sessions (course_id, start_date, end_date, location) VALUES
((SELECT id FROM courses WHERE title = 'Annual Firearms Requalification'), '2025-11-03', '2025-11-03', 'Belmopan Firing Range'),
((SELECT id FROM courses WHERE title = 'Advanced Report Writing'), '2025-11-10', '2025-11-11', 'Police Training Academy, Classroom 3'),
((SELECT id FROM courses WHERE title = 'Community Policing Strategies'), '2025-11-17', '2025-11-18', 'San Ignacio Town Hall');


-- ========= 6. LINK Facilitators to Sessions =========
-- Assign facilitators to the sessions.
-- The Advanced Report Writing course will be taught by BOTH facilitators.

INSERT INTO session_facilitators (session_id, facilitator_id) VALUES
((SELECT id FROM training_sessions WHERE start_date = '2025-11-03'), (SELECT id FROM facilitators WHERE email = 'kelijo@police.bz')),
((SELECT id FROM training_sessions WHERE start_date = '2025-11-10'), (SELECT id FROM facilitators WHERE email = 'kelijo@police.bz')),
((SELECT id FROM training_sessions WHERE start_date = '2025-11-10'), (SELECT id FROM facilitators WHERE email = 'acrane.consulting@gmail.com'));


-- ========= 7. ENROLL Personnel in Sessions =========
-- This is a critical step to test different scenarios.

-- Scenario A: Everyone attends firearms training. Two complete, one fails.
INSERT INTO session_enrollment (personnel_id, session_id, status, completion_date) VALUES
    ((SELECT id FROM personnel WHERE regulation_number = '202'), (SELECT id FROM training_sessions WHERE start_date = '2025-11-03'), 'Completed', '2025-11-03'),
    ((SELECT id FROM personnel WHERE regulation_number = '303'), (SELECT id FROM training_sessions WHERE start_date = '2025-11-03'), 'Completed', '2025-11-03'),
    ((SELECT id FROM personnel WHERE regulation_number = '404'), (SELECT id FROM training_sessions WHERE start_date = '2025-11-03'), 'Failed', '2025-11-03');

-- Scenario B: Jane Smith and John Doe attend the writing course. Jane completes, John withdraws.
INSERT INTO session_enrollment (personnel_id, session_id, status, completion_date) VALUES
    ((SELECT id FROM personnel WHERE regulation_number = '202'), (SELECT id FROM training_sessions WHERE start_date = '2025-11-10'), 'Completed', '2025-11-11'),
    ((SELECT id FROM personnel WHERE regulation_number = '303'), (SELECT id FROM training_sessions WHERE start_date = '2025-11-10'), 'Withdrew', NULL);

-- Scenario C: Maria Garcia is enrolled in a future course.
INSERT INTO session_enrollment (personnel_id, session_id, status) VALUES
    ((SELECT id FROM personnel WHERE regulation_number = '404'), (SELECT id FROM training_sessions WHERE start_date = '2025-11-17'), 'Enrolled');


-- ========= 8. INSERT Ratings =========
-- Only officers who have 'Completed' a session can leave a rating.

-- Jane Smith rates the Firearms course she completed.
INSERT INTO course_ratings (session_enrollment_id, score, comment) VALUES
    ((SELECT id FROM session_enrollment WHERE personnel_id = (SELECT id FROM personnel WHERE regulation_number = '202') AND session_id = (SELECT id FROM training_sessions WHERE start_date = '2025-11-03')), 5, 'Excellent instruction and very well-organized range day.');

-- Jane Smith also rates Kenroy Elijo for that same Firearms course.
INSERT INTO facilitator_ratings (session_enrollment_id, facilitator_id, score, comment) VALUES
    ((SELECT id FROM session_enrollment WHERE personnel_id = (SELECT id FROM personnel WHERE regulation_number = '202') AND session_id = (SELECT id FROM training_sessions WHERE start_date = '2025-11-03')),
    (SELECT id FROM facilitators WHERE email = 'kelijo@police.bz'), 5, 'Inspector Elijo is a knowledgeable and patient instructor.');

-- Now, Jane Smith rates the Advanced Report Writing course AND both of its facilitators.
INSERT INTO course_ratings (session_enrollment_id, score, comment) VALUES
    ((SELECT id FROM session_enrollment WHERE personnel_id = (SELECT id FROM personnel WHERE regulation_number = '202') AND session_id = (SELECT id FROM training_sessions WHERE start_date = '2025-11-10')), 4, 'The content was very relevant, but a lot to cover in two days.');

INSERT INTO facilitator_ratings (session_enrollment_id, facilitator_id, score, comment) VALUES
    ((SELECT id FROM session_enrollment WHERE personnel_id = (SELECT id FROM personnel WHERE regulation_number = '202') AND session_id = (SELECT id FROM training_sessions WHERE start_date = '2025-11-10')),
    (SELECT id FROM facilitators WHERE email = 'kelijo@police.bz'), 4, 'Great real-world examples from his experience.'),
    ((SELECT id FROM session_enrollment WHERE personnel_id = (SELECT id FROM personnel WHERE regulation_number = '202') AND session_id = (SELECT id FROM training_sessions WHERE start_date = '2025-11-10')),
    (SELECT id FROM facilitators WHERE email = 'acrane.consulting@gmail.com'), 5, 'Dr. Crane provided fantastic insight into the legal aspects. Very clear communicator.');

-- ===================================================================
--                  END OF TEST DATA SCRIPT
-- ===================================================================