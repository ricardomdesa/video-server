-- name: GetCourses :many
SELECT * FROM courses;

-- name: GetAllVideos :many
SELECT m.name, c.name, v.* FROM videos v
join modules m on m.id = v.module_id
join courses c on c.id = m.course_id
WHERE c.id = ?;


-- name: AddCourse :exec
INSERT INTO courses (name) VALUES (?);

-- name: AddModule :exec
INSERT INTO modules (name, folder, course_id) VALUES (?, ?, ?);

-- name: AddVideo :exec
INSERT INTO videos (name, video_key, module_id) VALUES (?, ?, ?);

-- name: GetAllModules :many
SELECT * FROM modules WHERE course_id = ?;
