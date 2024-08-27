import React, { useEffect, useState } from 'react';

function StudentList({ courseId }) {
  const [students, setStudents] = useState([]);

  useEffect(() => {
    fetch(`http://localhost:8000/students/${courseId}`)
      .then(response => response.json())
      .then(data => setStudents(data));
  }, [courseId]);

  return (
    <div className="student-list">
      <h2>Alumnos del Curso {courseId}</h2>
      <ul>
        {students.map(student => (
          <li key={student.id}>
            <div className="student-card">
              <h3>{student.name}</h3>
              <p>ID: {student.id}</p>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default StudentList;
