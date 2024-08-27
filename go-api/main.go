package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "yourpassword"
	dbname   = "matriculasdb"
)

func main() {
	fmt.Println("Iniciando conexión a PostgreSQL...")

	// Conectar a PostgreSQL
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println("Intentando conectar a la base de datos con la siguiente cadena de conexión:", psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error al abrir la conexión a la base de datos: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error al hacer ping a la base de datos: %v", err)
	}

	fmt.Println("Conexión exitosa a PostgreSQL")

	r := mux.NewRouter()

	// Definir rutas
	r.HandleFunc("/students/{course_id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request recibido para obtener estudiantes")
		vars := mux.Vars(r)
		courseID := vars["course_id"]

		// Obtener estudiantes del curso
		rows, err := db.Query("SELECT id, name FROM students WHERE course_id=$1", courseID)
		if err != nil {
			http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		students := []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{}

		for rows.Next() {
			var student struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			}
			err := rows.Scan(&student.ID, &student.Name)
			if err != nil {
				http.Error(w, "Error al escanear resultados", http.StatusInternalServerError)
				return
			}
			students = append(students, student)
		}

		json.NewEncoder(w).Encode(students)
	}).Methods("GET")

	// Habilitar CORS para todas las rutas
	corsObj := handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:3000"}))

	fmt.Println("Iniciando el servidor en el puerto 8000")
	// Iniciar el servidor
	log.Fatal(http.ListenAndServe(":8000", corsObj(r)))
}
