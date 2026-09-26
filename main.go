package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	db "tpe/db/sqlc"
	"tpe/logic"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Server struct {
	queries *db.Queries
}

func main() {

	// Creamos la conexion con la base de datos
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Falta la variable de entorno DATABASE_URL")
	}

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("No se pudo abrir la conexion: %v", err)
	}
	defer conn.Close()

	//HEALTH CHECK para verificar conexion con base de datos
	var pingErr error
	for i := 0; i < 30; i++ {
		if pingErr = conn.Ping(); pingErr == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if pingErr != nil {
		log.Fatalf("No se pudo conectar a la base de datos tras varios intentos: %v", pingErr)
	}

	//

	server := &Server{
		queries: db.New(conn),
	}
	_ = server
	log.Println("Conectado a la base de datos correctamente")
	log.Fatal(http.ListenAndServe(":8080", nil))

	/*====================================
	END-POINTS
	=====================================*/
	/*
		mux := http.NewServeMux()
		mux.Handle("/categorias", loggingMiddleware(http.HandlerFunc(server.categoriasHandler)))
		mux.Handle("/categorias/", loggingMiddleware(http.HandlerFunc(server.categoriasHandler)))
		mux.Handle("/productos", loggingMiddleware(http.HandlerFunc(server.productosHandler)))
		mux.Handle("/productos/", loggingMiddleware(http.HandlerFunc(server.productosHandler)))
		mux.Handle("/clientes", loggingMiddleware(http.HandlerFunc(server.clientesHandler)))
		mux.Handle("/clientes/", loggingMiddleware(http.HandlerFunc(server.clientesHandler)))
		mux.Handle("/", loggingMiddleware(http.FileServer(http.Dir("./static"))))

		port := ":8080"
		log.Printf("Servidor escuchando en http://localhost%s\n", port)
		if err := http.ListenAndServe(port, mux); err != nil {
			log.Fatalf("Error al iniciar el servidor: %s", err)
		}
	*/
}

// -----------------------------------------------------------------------------------------------
// -----------------------------------------------------------------------------------------------
// -----------------------------------------------------------------------------------------------
// -----------------------------------------------------------------------------------------------
// CATEGORIAS

func (s *Server) categoriasHandler(w http.ResponseWriter, r *http.Request) {
	/*
		Si llega /categorias/id/
		/id/
		id

		O puede llegar /categorias/algo/id/
		/algo/id/
		algo/id

		Entonces preguntamos si contiene /. De esa forma sabemos que la ruta esta mal
	*/

	path := strings.TrimPrefix(r.URL.Path, "/categorias")
	path = strings.Trim(path, "/")

	if path == "" {
		switch r.Method {
		case http.MethodGet:
			s.listCategorias(w, r)
		case http.MethodPost:
			s.createCategoria(w, r)
		default:
			http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		}
		return
	}

	if strings.Contains(path, "/") {
		http.Error(w, "Ruta no encontrada", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		s.getCategoria(w, r, int32(id))
	case http.MethodPut:
		s.updateCategoria(w, r, int32(id))
	case http.MethodDelete:
		s.deleteCategoria(w, r, int32(id))
	default:
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
	}
}

func (s *Server) listCategorias(w http.ResponseWriter, r *http.Request) {
	lista, err := s.queries.ListarCategorias(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if lista == nil {
		lista = []db.Categoria{}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lista)
}

func (s *Server) createCategoria(w http.ResponseWriter, r *http.Request) {
	var in CategoriaInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "JSON invalido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := logic.ValidateCategoria(in.Nombre); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c, err := s.queries.CrearCategoria(r.Context(), db.CrearCategoriaParams{
		Nombre:      in.Nombre,
		Descripcion: toNullString(in.Descripcion),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func (s *Server) getCategoria(w http.ResponseWriter, r *http.Request, id int32) {
	c, err := s.queries.GetCategoria(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Categoria no encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func (s *Server) updateCategoria(w http.ResponseWriter, r *http.Request, id int32) {
	if _, err := s.queries.GetCategoria(r.Context(), id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Categoria no encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var in CategoriaInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "JSON invalido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := logic.ValidateCategoria(in.Nombre); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := s.queries.ActualizarCategoria(r.Context(), db.ActualizarCategoriaParams{
		ID:          id,
		Nombre:      in.Nombre,
		Descripcion: toNullString(in.Descripcion),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c, err := s.queries.GetCategoria(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func (s *Server) deleteCategoria(w http.ResponseWriter, r *http.Request, id int32) {
	if _, err := s.queries.GetCategoria(r.Context(), id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Categoria no encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.queries.BorrarCategoria(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// -----------------------------------------------------------------------------------------------
// -----------------------------------------------------------------------------------------------
// -----------------------------------------------------------------------------------------------
// -----------------------------------------------------------------------------------------------
// PRODUCTOS

// traduce de la capa de datos a la capa de negocio
func aLogic(p db.Producto) (logic.producto, error) {
	precio, err := strconv.ParseFloat(p.Precio, 64)
	if err != nil {
		return logic.producto{}, err
	}
	stock, err := strconv.ParseFloat(p.Stock, 64)
	if err != nil {
		return logic.producto{}, err
	}
	return logic.producto{
		ID:          p.ID,
		IDCategoria: p.IDCategoria,
		Nombre:      p.Nombre,
		Descripcion: p.Descripcion.String,
		Stock:       stock,
		Precio:      precio,
	}, nil
}
