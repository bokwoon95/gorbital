DROP VIEW IF EXISTS v_teams CASCADE;
CREATE OR REPLACE VIEW v_teams AS
WITH participants_full AS (
    SELECT uid, nusnetid, password, display_name, openid, email, cohort, team
    FROM users JOIN participants USING (uid)
)
SELECT
    p1.cohort
    ,t.team_project_level
    ,t.team_name AS team_name
    ,p1.display_name AS participant1
    ,p2.display_name AS participant2
    ,ua.display_name AS adviser
    ,um.display_name AS mentor
    ,t.ignition_pitch_poster
FROM
    teams t
    LEFT JOIN users AS ua ON ua.uid = t.adviser
    LEFT JOIN users AS um ON um.uid = t.mentor
    JOIN participants_full AS p1 ON p1.team = t.tid
    JOIN participants_full AS p2 ON p2.team = t.tid AND p1.display_name < p2.display_name
;

DROP VIEW IF EXISTS v_submissions CASCADE;
CREATE OR REPLACE VIEW v_submissions AS
WITH participants_full AS (
    SELECT uid, nusnetid, password, display_name, openid, email, cohort,team
    FROM users JOIN participants USING (uid)
)
SELECT
    s.milestone
    ,s.project_level
    ,s.project_name
    ,s.project_link
    ,s.project_readme
    ,s.project_poster
    ,s.project_video
    ,p1.cohort
    ,t.team_project_level
    ,t.team_name
    ,p1.display_name AS participant1
    ,p2.display_name AS participant2
    ,ua.display_name AS adviser
    ,um.display_name AS mentor
    ,t.ignition_pitch_poster
FROM
    teams t
    LEFT JOIN users AS ua ON ua.uid = t.adviser
    LEFT JOIN users AS um ON um.uid = t.mentor
    JOIN participants_full AS p1 ON p1.team = t.tid
    JOIN participants_full AS p2 ON p2.team = t.tid AND p1.display_name < p2.display_name
    JOIN submissions AS s ON s.team = t.tid
;
