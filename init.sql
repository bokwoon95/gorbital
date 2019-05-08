-- (Schema here) https://docs.google.com/spreadsheets/d/1soGXKMGHa9YAhCsl8QMVBCrWavu1tJWvSADyjT3Ftn8/edit#gid=0

-- Drop all tables
DO $$ DECLARE
  r RECORD;
BEGIN
  FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
    EXECUTE 'DROP TABLE ' || quote_ident(r.tablename) || ' CASCADE';
  END LOOP;
END $$;

CREATE TABLE users (
    uid INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    nusnetid TEXT NOT NULL,
    password TEXT NOT NULL,
    display_name TEXT NOT NULL UNIQUE,
    openid TEXT,
    email TEXT
);

CREATE TABLE admins (
    uid INT UNIQUE,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    PRIMARY KEY (uid, cohort)
);

CREATE TABLE advisers (
    uid INT UNIQUE,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    PRIMARY KEY (uid, cohort)
);

CREATE TABLE mentors (
    uid INT UNIQUE,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    PRIMARY KEY (uid, cohort)
);

CREATE TABLE participants (
    uid INT UNIQUE,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    team INT,
    PRIMARY KEY (uid, cohort)
);

CREATE TABLE tutors (
    uid INT UNIQUE,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    PRIMARY KEY (uid, cohort)
);

CREATE TABLE facilitators (
    uid INT UNIQUE,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    PRIMARY KEY (uid, cohort)
);

CREATE TABLE teams (
    tid INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    teamname TEXT NOT NULL,
    current_project_level text DEFAULT 'gemini',
    adviser INT,
    mentor INT
);

CREATE TABLE milestones (
    mid INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    phase INT NOT NULL,
    submission_deadline TIMESTAMPTZ NOT NULL,
    evaluation_deadline TIMESTAMPTZ NOT NULL
);

CREATE TABLE submissions (
    sid INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    team INT NOT NULL,
    milestone INT NOT NULL,
    project_level TEXT NOT NULL,
    project_link TEXT NOT NULL,
    project_readme TEXT NOT NULL,
    project_poster TEXT,
    project_video TEXT,
    submission_hash TEXT NOT NULL
);

CREATE table evaluation_templates (
    etid INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    template TEXT NOT NULL,
    formdown TEXT
);

CREATE TABLE evaluations (
    eid INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    evaluation_template INT NOT NULL,
    evaluation_content TEXT NOT NULL,
    evaluation_hash TEXT NOT NULL
);

CREATE TABLE peer_evaluationships (
    peid INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    evaluator INT,
    evaluatee INT,
    evaluation INT
);

CREATE TABLE adviser_evaluationships (
    aeid INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    adviser INT,
    evaluatee INT,
    evaluation INT
);

CREATE TABLE orbital_state (
    oid INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    data jsonb
);

-- data {
-- liftoff: [ ],
-- mission_control: {
--   webdev: ,
--   android: ,
--   ios: ,
-- },
-- milestone1_start: ,
-- milestone1_end: ,
-- milestone2_start: ,
-- milestone2_end: ,
-- milestone3_start: ,
-- milestone3_end: ,
-- splashdown: [ ],
-- }

ALTER TABLE admins ADD CONSTRAINT admins_uid_fkey FOREIGN KEY (uid) REFERENCES users (uid) ON DELETE CASCADE;

ALTER TABLE advisers ADD CONSTRAINT advisers_uid_fkey FOREIGN KEY (uid) REFERENCES users (uid) ON DELETE CASCADE;

ALTER TABLE mentors ADD CONSTRAINT mentors_uid_fkey FOREIGN KEY (uid) REFERENCES users (uid) ON DELETE CASCADE;

ALTER TABLE participants ADD CONSTRAINT participants_uid_fkey FOREIGN KEY (uid) REFERENCES users (uid) ON DELETE CASCADE;
ALTER TABLE participants ADD CONSTRAINT participants_team_fkey FOREIGN KEY (team) REFERENCES teams (tid) ON DELETE CASCADE;

ALTER TABLE tutors ADD CONSTRAINT tutors_uid_fkey FOREIGN KEY (uid) REFERENCES users (uid) ON DELETE CASCADE;

ALTER TABLE facilitators ADD CONSTRAINT facilitators_uid_fkey FOREIGN KEY (uid) REFERENCES users (uid) ON DELETE CASCADE;

ALTER TABLE teams ADD CONSTRAINT teams_adviser_fkey FOREIGN KEY (adviser) REFERENCES advisers (uid) ON DELETE CASCADE;
ALTER TABLE teams ADD CONSTRAINT teams_mentor_fkey FOREIGN KEY (mentor) REFERENCES mentors (uid) ON DELETE CASCADE;

ALTER TABLE submissions ADD CONSTRAINT submissions_team_fkey FOREIGN KEY (team) REFERENCES teams (tid) ON DELETE CASCADE;
ALTER TABLE submissions ADD CONSTRAINT submissions_milestone_fkey FOREIGN KEY (milestone) REFERENCES milestones (mid) ON DELETE CASCADE;

ALTER TABLE evaluations ADD CONSTRAINT evaluations_evaluation_template_fkey FOREIGN KEY (evaluation_template) REFERENCES evaluation_templates (etid) ON DELETE CASCADE;

ALTER TABLE peer_evaluationships ADD CONSTRAINT peer_evaluationships_evaluator_fkey FOREIGN KEY (evaluator) REFERENCES teams (tid) ON DELETE CASCADE;
ALTER TABLE peer_evaluationships ADD CONSTRAINT peer_evaluationships_evaluatee_fkey FOREIGN KEY (evaluatee) REFERENCES submissions (sid) ON DELETE CASCADE;
ALTER TABLE peer_evaluationships ADD CONSTRAINT peer_evaluationships_evaluation FOREIGN KEY (evaluation) REFERENCES evaluations (eid) ON DELETE CASCADE;

ALTER TABLE adviser_evaluationships ADD CONSTRAINT adviser_evaluationships_adviser_fkey FOREIGN KEY (adviser) REFERENCES advisers (uid) ON DELETE CASCADE;
ALTER TABLE adviser_evaluationships ADD CONSTRAINT adviser_evaluationships_evaluatee_fkey FOREIGN KEY (evaluatee) REFERENCES submissions (sid) ON DELETE CASCADE;
ALTER TABLE adviser_evaluationships ADD CONSTRAINT advisers_evaluation_fkey FOREIGN KEY (evaluation) REFERENCES evaluations (eid) ON DELETE CASCADE;
