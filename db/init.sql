CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    course_id INT REFERENCES courses (id)
);

INSERT INTO courses (name) VALUES ('Progra3'), ('IngeSoft');

INSERT INTO
    students (name, course_id)
VALUES ('Alumno1', 1),
    ('Alumno2', 1),
    ('Alumno3', 2);