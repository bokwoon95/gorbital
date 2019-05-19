-- (Schema here) https://docs.google.com/spreadsheets/d/1soGXKMGHa9YAhCsl8QMVBCrWavu1tJWvSADyjT3Ftn8/edit#gid=0

-- Drop all tables
DO $$ DECLARE
  r RECORD;
BEGIN
  FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
    EXECUTE 'DROP TABLE ' || quote_ident(r.tablename) || ' CASCADE';
  END LOOP;
END $$;

-- Drop stored procedures
DROP PROCEDURE IF EXISTS insert_milestone;
DROP PROCEDURE IF EXISTS insert_submission;

-- Drop extentions
DROP EXTENSION IF EXISTS pgcrypto;
CREATE EXTENSION pgcrypto;

CREATE TABLE users (
    uid SERIAL UNIQUE,
    nusnetid TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    display_name TEXT NOT NULL UNIQUE,
    openid TEXT,
    email TEXT
);

CREATE TABLE participants (
    uid INT,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    team INT,
    PRIMARY KEY (uid, cohort)
);

CREATE OR REPLACE VIEW v_participants AS
SELECT * FROM users JOIN participants USING (uid);

CREATE TABLE tutors (
    uid INT,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    PRIMARY KEY (uid, cohort)
);

CREATE OR REPLACE VIEW v_tutors AS
SELECT * FROM users JOIN tutors USING (uid);

CREATE TABLE advisers (
    uid INT,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    PRIMARY KEY (uid, cohort)
);

CREATE OR REPLACE VIEW v_advisers AS
SELECT * FROM users JOIN advisers USING (uid);

CREATE TABLE mentors (
    uid INT,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    PRIMARY KEY (uid, cohort)
);

CREATE OR REPLACE VIEW v_mentors AS
SELECT * FROM users JOIN mentors USING (uid);

CREATE TABLE admins (
    uid INT,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    PRIMARY KEY (uid, cohort)
);

CREATE OR REPLACE VIEW v_admins AS
SELECT * FROM users JOIN admins USING (uid);

CREATE TABLE facilitators (
    uid INT,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    PRIMARY KEY (uid, cohort)
);

CREATE OR REPLACE VIEW v_facilitators AS
SELECT * FROM users JOIN facilitators USING (uid);

CREATE TABLE teams (
    tid SERIAL UNIQUE,
    teamname TEXT NOT NULL,
    current_project_level text DEFAULT 'gemini',
    ignition_pitch_poster TEXT,
    adviser INT,
    mentor INT
);

CREATE TABLE milestones (
    mid SERIAL UNIQUE,
    cohort TEXT NOT NULL DEFAULT date_part('year', CURRENT_DATE),
    phase TEXT NOT NULL,
    submission_deadline TIMESTAMPTZ NOT NULL,
    evaluation_deadline TIMESTAMPTZ NOT NULL
);

CREATE TABLE submissions (
    sid SERIAL UNIQUE,
    team INT NOT NULL,
    milestone INT NOT NULL,
    project_level TEXT NOT NULL,
    project_name TEXT NOT NULL,
    project_link TEXT NOT NULL,
    project_readme TEXT NOT NULL,
    project_poster TEXT,
    project_video TEXT,
    submission_hash TEXT NOT NULL
);

CREATE table evaluation_templates (
    etid SERIAL UNIQUE,
    template TEXT NOT NULL,
    formdown TEXT
);

CREATE TABLE evaluations (
    eid SERIAL UNIQUE,
    evaluation_template INT NOT NULL,
    evaluation_content TEXT NOT NULL,
    evaluation_hash TEXT NOT NULL
);

CREATE TABLE peer_evaluationships (
    peid SERIAL UNIQUE,
    evaluator INT,
    evaluatee INT,
    evaluation INT
);

CREATE TABLE adviser_evaluationships (
    aeid SERIAL UNIQUE,
    adviser INT,
    evaluatee INT,
    evaluation INT
);

CREATE TABLE orbital_state (
    oid SERIAL UNIQUE,
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
-- splashdown: [ ],
-- }

ALTER TABLE admins ADD CONSTRAINT admins_partial_uid_fkey FOREIGN KEY (uid) REFERENCES users (uid) ON DELETE CASCADE;

ALTER TABLE advisers ADD CONSTRAINT advisers_partial_uid_fkey FOREIGN KEY (uid) REFERENCES users (uid) ON DELETE CASCADE;

ALTER TABLE mentors ADD CONSTRAINT mentors_uid_fkey FOREIGN KEY (uid) REFERENCES users (uid) ON DELETE CASCADE;

ALTER TABLE participants ADD CONSTRAINT participants_uid_fkey FOREIGN KEY (uid) REFERENCES users (uid) ON DELETE CASCADE;
ALTER TABLE participants ADD CONSTRAINT participants_team_fkey FOREIGN KEY (team) REFERENCES teams (tid) ON DELETE CASCADE;

ALTER TABLE tutors ADD CONSTRAINT tutors_uid_fkey FOREIGN KEY (uid) REFERENCES users (uid) ON DELETE CASCADE;

ALTER TABLE facilitators ADD CONSTRAINT facilitators_uid_fkey FOREIGN KEY (uid) REFERENCES users (uid) ON DELETE CASCADE;

ALTER TABLE teams ADD CONSTRAINT teams_adviser_fkey FOREIGN KEY (adviser) REFERENCES users (uid) ON DELETE CASCADE;
ALTER TABLE teams ADD CONSTRAINT teams_mentor_fkey FOREIGN KEY (mentor) REFERENCES users (uid) ON DELETE CASCADE;

ALTER TABLE milestones ADD CONSTRAINT milestones_submission_before_evaluation CHECK (submission_deadline < evaluation_deadline);
ALTER TABLE milestones ADD CONSTRAINT milestones_unique_cohort_phase UNIQUE (cohort, phase);

ALTER TABLE submissions ADD CONSTRAINT submissions_team_fkey FOREIGN KEY (team) REFERENCES teams (tid) ON DELETE CASCADE;
ALTER TABLE submissions ADD CONSTRAINT submissions_milestone_fkey FOREIGN KEY (milestone) REFERENCES milestones (mid) ON DELETE CASCADE;
ALTER TABLE submissions ADD CONSTRAINT submissions_unique_team_milestone UNIQUE (team, milestone);

ALTER TABLE evaluations ADD CONSTRAINT evaluations_evaluation_template_fkey FOREIGN KEY (evaluation_template) REFERENCES evaluation_templates (etid) ON DELETE CASCADE;

ALTER TABLE peer_evaluationships ADD CONSTRAINT peer_evaluationships_evaluator_fkey FOREIGN KEY (evaluator) REFERENCES teams (tid) ON DELETE CASCADE;
ALTER TABLE peer_evaluationships ADD CONSTRAINT peer_evaluationships_evaluatee_fkey FOREIGN KEY (evaluatee) REFERENCES submissions (sid) ON DELETE CASCADE;
ALTER TABLE peer_evaluationships ADD CONSTRAINT peer_evaluationships_evaluation FOREIGN KEY (evaluation) REFERENCES evaluations (eid) ON DELETE CASCADE;

ALTER TABLE adviser_evaluationships ADD CONSTRAINT adviser_evaluationships_adviser_fkey FOREIGN KEY (adviser) REFERENCES users (uid) ON DELETE CASCADE;
ALTER TABLE adviser_evaluationships ADD CONSTRAINT adviser_evaluationships_evaluatee_fkey FOREIGN KEY (evaluatee) REFERENCES submissions (sid) ON DELETE CASCADE;
ALTER TABLE adviser_evaluationships ADD CONSTRAINT advisers_evaluation_fkey FOREIGN KEY (evaluation) REFERENCES evaluations (eid) ON DELETE CASCADE;
