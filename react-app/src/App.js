import StudentList from './components/StudentList';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1>Lista de Alumnos</h1>
      </header>
      <StudentList courseId={1} />
    </div>
  );
}

export default App;
