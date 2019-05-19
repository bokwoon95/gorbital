DROP VIEW IF EXISTS vvv_teams CASCADE;
DROP VIEW IF EXISTS vv_participant_pairs CASCADE;

CREATE OR REPLACE VIEW vv_participant_pairs AS
SELECT
    a.team
    ,a.cohort
    ,a.uid AS uid1
    ,a.display_name AS participant1
    ,b.uid AS uid2
    ,b.display_name AS participant2
FROM
    v_participants a
    JOIN v_participants b ON a.team = b.team AND a.display_name < b.display_name
;

CREATE OR REPLACE VIEW vvv_teams AS
SELECT
    vv_pp.cohort
    ,t.current_project_level
    ,t.teamname
    ,vv_pp.participant1
    ,vv_pp.participant2
    ,v_a.display_name AS adviser_name
    ,v_m.display_name AS mentor_name
    ,t.ignition_pitch_poster
FROM
    teams t
    LEFT JOIN v_advisers AS v_a ON v_a.uid = t.adviser
    LEFT JOIN v_mentors AS v_m ON v_m.uid = t.mentor
    JOIN vv_participant_pairs AS vv_pp ON vv_pp.team = t.tid
;
