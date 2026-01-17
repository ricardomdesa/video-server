CREATE TABLE courses (
  id   INTEGER PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE modules (
  id   INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  folder TEXT NOT NULL,
  course_id INTEGER NOT NULL,
  FOREIGN KEY (course_id) REFERENCES courses(id)
);

CREATE TABLE videos (
  id   INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  video_key TEXT NOT NULL,
  module_id INTEGER NOT NULL,
  FOREIGN KEY (module_id) REFERENCES modules(id)
);