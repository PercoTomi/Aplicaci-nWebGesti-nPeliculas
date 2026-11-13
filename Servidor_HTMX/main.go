package main

import (
	// --- Paquetes Personalizados ---
    sqlc "TPEspecialParte6/internal/db/sqlc"
    "TPEspecialParte6/internal/views"

	// --- Paquetes Estándar ---
    "context"
    "database/sql"
    "log"
    "net/http"
    "strconv"
    "time"

	// --- Driver de Postgres ---
	_ "github.com/lib/pq"
)

func main() {
	// --- Conexión a la Base de Datos ---
	connStr := "user=user password=admin dbname=tpe_parte_6 host=localhost port=5432 sslmode=disable"
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer database.Close()

	// Inicializa sqlc
	queries := sqlc.New(database)
	ctx := context.Background()

	// --- Handler Principal (GET /) ---
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		
		movies, err := queries.ListMovies(ctx) 
		if err != nil {
			log.Printf("Error al listar películas: %v", err)
			http.Error(w, "Error al cargar datos", http.StatusInternalServerError)
			return
		}

		component := views.IndexPage("Mi catálogo de peliculas", movies)
		
		component.Render(ctx, w)
	})

    http.HandleFunc("POST /movies", func(w http.ResponseWriter, r *http.Request) {
        
		if err := r.ParseForm(); err != nil {
			log.Printf("Error al parsear formulario: %v", err)
			http.Error(w, "Datos de formulario inválidos", http.StatusBadRequest)
			return
		}

		duracion, _ := strconv.Atoi(r.PostForm.Get("duracion"))
        dateString := r.PostForm.Get("fecha") 
        fecha, err := time.Parse("2006-01-02", dateString) // fecha en formato YYYY-MM-DD
		
        if err != nil {
			log.Printf("Error al Parsear la fecha: %v", err)
			http.Error(w, "Fecha Invalida", http.StatusInternalServerError)
			return
		}

        // --- Generamos la pelicula ---
        params := sqlc.CreateMovieParams{
			Name:        r.PostForm.Get("nombre"),
			Description: r.PostForm.Get("descripcion"),
			ReleaseDate: fecha,
			Runtime:     int32(duracion),      
			Genre:       r.PostForm.Get("genero"),
			Director:    r.PostForm.Get("director"),
		}

        // --- Guardamos en BD ---

        if _, err := queries.CreateMovie(ctx, params); err != nil {
            log.Printf("Error al crear: %v", err)
            http.Error(w, "Error al guardar", 500)
            return
        }
        
        // --- Volvemos a consultar la lista actualizada ---
        movies, err := queries.ListMovies(ctx)
        if err != nil {
             http.Error(w, "Error al leer lista actualizada", 500)
             return
        }

        // --- En lugar de redirigir, devolvemos el componente HTML de la lista ---
        views.MovieList(movies).Render(ctx, w)
    })

    http.HandleFunc("DELETE /movies/{id}", func(w http.ResponseWriter, r *http.Request) {
        
        // --- Obtiene el ID de la URL ---
        idStr := r.PathValue("id") 
        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "ID inválido", http.StatusBadRequest)
            return
        }

        // --- Elimina de la Base de Datos ---
        err = queries.DeleteMovie(ctx, int32(id))
        if err != nil {
            log.Printf("Error eliminando: %v", err)
            http.Error(w, "Error al eliminar", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK) // Status 200
    })

	// --- Inicia Servidor ---
	log.Println("Servidor SSR iniciando en http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
