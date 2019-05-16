INSERT INTO users (uid, nusnetid, password, display_name) VALUES (01,'e0031874','password','Admin01') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (02,'e0031875','password','Admin02') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (03,'e0031876','password','Facilitator03') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (04,'e0031877','password','Facilitator04') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (05,'e0031878','password','Mentor05') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (06,'e0031879','password','Mentor06') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (07,'e0031880','password','Adviser07') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (08,'e0031881','password','Adviser08') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (09,'e0031882','password','Tutor09') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (10,'e0031883','password','Tutor10') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (11,'e0031884','password','User11') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (12,'e0031885','password','User12') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (13,'e0031886','password','User13') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (14,'e0031887','password','User14') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (15,'e0031888','password','User15') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (16,'e0031889','password','User16') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (17,'e0031890','password','User17') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (18,'e0031891','password','User18') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (19,'e0031892','password','User19') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (20,'e0031893','password','User20') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (21,'e0031894','password','User21') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (22,'e0031895','password','User22') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (23,'e0031896','password','User23') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (24,'e0031897','password','User24') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (25,'e0031898','password','User25') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (26,'e0031899','password','User26') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (27,'e0031900','password','User27') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (28,'e0031901','password','User28') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (29,'e0031902','password','User29') ON CONFLICT DO NOTHING;
INSERT INTO users (uid, nusnetid, password, display_name) VALUES (30,'e0031903','password','User30') ON CONFLICT DO NOTHING;

INSERT INTO admins (uid) VALUES (1) ON CONFLICT DO NOTHING;
INSERT INTO admins (uid) VALUES (2) ON CONFLICT DO NOTHING;

INSERT INTO facilitators (uid) VALUES (3) ON CONFLICT DO NOTHING;
INSERT INTO facilitators (uid) VALUES (4) ON CONFLICT DO NOTHING;

INSERT INTO mentors (uid) VALUES (5) ON CONFLICT DO NOTHING;
INSERT INTO mentors (uid) VALUES (6) ON CONFLICT DO NOTHING;

INSERT INTO advisers (uid) VALUES (7) ON CONFLICT DO NOTHING;
INSERT INTO advisers (uid) VALUES (8) ON CONFLICT DO NOTHING;

INSERT INTO tutors (uid) VALUES (9) ON CONFLICT DO NOTHING;
INSERT INTO tutors (uid) VALUES (10) ON CONFLICT DO NOTHING;

INSERT INTO participants (uid) VALUES (11) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (12) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (13) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (14) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (15) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (16) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (17) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (18) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (19) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (20) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (21) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (22) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (23) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (24) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (25) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (26) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (27) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (28) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (29) ON CONFLICT DO NOTHING;
INSERT INTO participants (uid) VALUES (30) ON CONFLICT DO NOTHING;

INSERT INTO teams (tid, teamname, current_project_level, ignition_pitch_poster, adviser, mentor) VALUES (1,'PaperBack','vostok','https://image.ibb.co/n96hwJ/Initial_Idea.jpg',7,NULL) ON CONFLICT DO NOTHING;
INSERT INTO teams (tid, teamname, current_project_level, ignition_pitch_poster, adviser, mentor) VALUES (2,'MaveWonders','vostok','https://image.ibb.co/mwWd3y/Ignition_Spot_Sports.png',7,NULL) ON CONFLICT DO NOTHING;
INSERT INTO teams (tid, teamname, current_project_level, ignition_pitch_poster, adviser, mentor) VALUES (3,'1638','vostok','',7,NULL) ON CONFLICT DO NOTHING;
INSERT INTO teams (tid, teamname, current_project_level, ignition_pitch_poster, adviser, mentor) VALUES (4,'EatLah!','gemini','https://image.ibb.co/ioSomJ/Orbital_1.jpg',7, NULL) ON CONFLICT DO NOTHING;
INSERT INTO teams (tid, teamname, current_project_level, ignition_pitch_poster, adviser, mentor) VALUES (5,'1473','gemini','https://i.imgur.com/zdNULkO.jpg',8,5) ON CONFLICT DO NOTHING;
INSERT INTO teams (tid, teamname, current_project_level, ignition_pitch_poster, adviser, mentor) VALUES (6,'Garanguni','gemini','https://image.ibb.co/ctEhfd/Slide1.jpg',8,5) ON CONFLICT DO NOTHING;
INSERT INTO teams (tid, teamname, current_project_level, ignition_pitch_poster, adviser, mentor) VALUES (7,'Google','apollo','https://i.imgur.com/PxsuEwc.jpg',8,6) ON CONFLICT DO NOTHING;
INSERT INTO teams (tid, teamname, current_project_level, ignition_pitch_poster, adviser, mentor) VALUES (8,'1509','apollo','https://image.ibb.co/kQEjRJ/Screen_Shot_2018_05_15_at_2_30_22_PM.png',8,6) ON CONFLICT DO NOTHING;
INSERT INTO teams (tid, teamname, current_project_level, ignition_pitch_poster, adviser, mentor) VALUES (9,'KZMY','apollo','https://image.ibb.co/mCJm6J/Team_KZMY.jpg',8,6) ON CONFLICT DO NOTHING;

UPDATE participants SET team = 1 WHERE uid = 11;
UPDATE participants SET team = 1 WHERE uid = 12;
UPDATE participants SET team = 2 WHERE uid = 13;
UPDATE participants SET team = 2 WHERE uid = 14;
UPDATE participants SET team = 3 WHERE uid = 15;
UPDATE participants SET team = 3 WHERE uid = 16;
UPDATE participants SET team = 4 WHERE uid = 17;
UPDATE participants SET team = 4 WHERE uid = 18;
UPDATE participants SET team = 5 WHERE uid = 19;
UPDATE participants SET team = 5 WHERE uid = 20;
UPDATE participants SET team = 6 WHERE uid = 21;
UPDATE participants SET team = 6 WHERE uid = 22;
UPDATE participants SET team = 7 WHERE uid = 23;
UPDATE participants SET team = 7 WHERE uid = 24;
UPDATE participants SET team = 8 WHERE uid = 25;
UPDATE participants SET team = 8 WHERE uid = 26;
UPDATE participants SET team = 9 WHERE uid = 27;
UPDATE participants SET team = 9 WHERE uid = 28;

CREATE OR REPLACE PROCEDURE
insert_milestone(mid INT, phase TEXT, submission_offset_from_now INTERVAL, evaluation_offset_from_now INTERVAL) AS $$
DECLARE
    submission_date TIMESTAMPTZ := (SELECT now() + submission_offset_from_now);
    evaluation_date TIMESTAMPTZ := (SELECT now() + evaluation_offset_from_now);
BEGIN
    RAISE NOTICE 'INSERT INTO milestones (mid, phase, submission_deadline, evaluation_deadline) VALUES (%, %, %, %)', mid, phase, submission_date, evaluation_date;
    INSERT INTO milestones (mid, phase, submission_deadline, evaluation_deadline) VALUES (mid, phase, submission_date, evaluation_date) ON CONFLICT DO NOTHING;
END;
$$ LANGUAGE plpgsql;
CALL insert_milestone(1,'1','-1 month -2 weeks','-1 month');
CALL insert_milestone(2,'2','-1 week','1 week');
CALL insert_milestone(3,'3','1 month','1 month 2 weeks');

CREATE OR REPLACE PROCEDURE
insert_submission(team INT, milestone INT, project_level TEXT, project_name TEXT, project_link TEXT, project_readme TEXT, project_poster TEXT, project_video TEXT)
AS $$ DECLARE
    bigstring TEXT := (SELECT CONCAT(team, milestone, project_level, project_link, project_readme, project_poster, project_video));
    submission_hash TEXT := (SELECT encode(digest(bigstring, 'md5'), 'hex'));
BEGIN
    RAISE NOTICE 'INSERT INTO submissions (team, milestone, project_level, project_name, project_link, project_readme, project_poster, project_video, submission_hash)
    VALUES (%, %, %, %, %, %, %, %, %)', team, milestone, project_level, project_name, project_link, project_readme, project_poster, project_video, submission_hash;
    INSERT INTO submissions (team, milestone, project_level, project_name, project_link, project_readme, project_poster, project_video, submission_hash)
    VALUES (team, milestone, project_level, project_name, project_link, project_readme, project_poster, project_video, submission_hash) ON CONFLICT DO NOTHING;
END;
$$ LANGUAGE plpgsql;
CALL insert_submission(1,1,'vostok','PaperBack','https://github.com','https://github.com','https://ibb.co/e35LQ8','https://www.powtoon.com/online-presentation/elpuJ5cldKh/?mode=movie');
CALL insert_submission(2,1,'vostok','MaveWonders','https://github.com','https://github.com','https://drive.google.com/file/d/1f4xgcLr6I4A78lJSlSaMPt8itoT9KD3i/view','https://drive.google.com/file/d/1s2gl94G_uBgXTKTEgv7YCtmZV6_NQxNC/view');
CALL insert_submission(3,1,'vostok','Study Companion','https://github.com','https://github.com','https://imgur.com/OZBAhfX','https://www.youtube.com/watch?v=G_vs_XAtp08&feature=youtu.be');
CALL insert_submission(4,1,'gemini','EatLah!','https://github.com','https://github.com','https://1drv.ms/v/s!AvHQl1WYqEpSiLMCgf_2rtCS7NYXKw','https://drive.google.com/file/d/1kJJji6L6CcV5lROqBKrie6YxUbu3yTIg/view');
CALL insert_submission(5,1,'gemini','Medician','https://github.com','https://github.com','https://i.imgur.com/zdNULkO.jpg','https://youtu.be/eD7-GARUNds');
CALL insert_submission(6,1,'gemini','Call-a-garanguni','https://github.com','https://github.com', NULL,'https://www.youtube.com/watch?v=qyi6i0DXdPY&feature=youtu.be&t=572');
CALL insert_submission(7,1,'apollo','Scuttlebutt','https://github.com','https://github.com','https://imgur.com/CSPEDdx','https://www.youtube.com/watch?v=7FqG6LuHvGg&feature=youtu.be');
CALL insert_submission(8,1,'apollo','Object Builder','https://github.com','https://github.com','https://imgur.com/a/evRRshD','https://www.youtube.com/watch?v=Kul2qdbaieI');
CALL insert_submission(9,1,'apollo','KZMY','https://github.com','https://github.com','https://docs.google.com/document/d/1UhRvZ94q7FTMikgODaTk2UaIyDtcvIYLrdAzpQ7rNI4/edit','https://www.youtube.com/watch?v=pwt63gOzNnw&feature=youtu.be');

INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (1,2,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (1,3,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (2,1,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (2,3,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (3,1,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (3,2,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (4,5,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (4,6,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (5,4,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (5,6,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (6,4,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (6,5,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (7,8,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (7,9,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (8,7,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (8,9,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (9,7,NULL) ON CONFLICT DO NOTHING;
INSERT INTO peer_evaluationships (evaluator, evaluatee, evaluation) VALUES (9,8,NULL) ON CONFLICT DO NOTHING;
