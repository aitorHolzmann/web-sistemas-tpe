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
		mux.Handle("/categorias", http.HandlerFunc(server.categoriasHandler))
		mux.Handle("/categorias/", http.HandlerFunc(server.categoriasHandler))
		mux.Handle("/productos", http.HandlerFunc(server.productosHandler))
		mux.Handle("/productos/", http.HandlerFunc(server.productosHandler))
		//para clientes tengo en cuenta clientes y clientes/ por si piden listar clientes o si piden algo especifico clientes/...
		mux.Handle("/clientes", http.HandlerFunc(server.clientesHandler))
		mux.Handle("/clientes/",http.HandlerFunc(server.clientesHandler))
		mux.Handle("/", http.FileServer(http.Dir("./static")))

		port := ":8080"
		log.Printf("Servidor escuchando en http://localhost%s\n", port)
		if err := http.ListenAndServe(port, mux); err != nil {
			log.Fatalf("Error al iniciar el servidor: %s", err)
		}
	*/
}

//Como en la db tenemos atributos que pueden ser null, de esta forma puedo resolverlo para strings

func toNullString(s string) sql.NullString {
	if strings.TrimSpace(s) == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}

// -----------------------------------------------------------------------------------------------
// CLIENTE
// -----------------------------------------------------------------------------------------------

func (s *Server) clientesHandler(w http.ResponseWriter, r *http.Request) {
	// quitamos el prefijo "/clientes" y las barras sobrantes
	path := strings.TrimPrefix(r.URL.Path, "/clientes")
	path = strings.Trim(path, "/")

	// si el path quedó vacío, es la colección general y solo puede ser para listar o crear clientes
	if path == "" {
		switch r.Method {
		case http.MethodGet:
			s.listClientes(w, r)
		case http.MethodPost:
			s.createCliente(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
		return
	}

	// si contiene barras intermedias (ej. /clientes/algo/1), es una ruta inválida
	if strings.Contains(path, "/") {
		http.Error(w, "Ruta no encontrada", http.StatusNotFound)
		return
	}

	// si hay algo en el path, debe ser el ID numérico
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// despachamos según el verbo HTTP para /clientes/{id}
	switch r.Method {
	case http.MethodGet:
		s.getCliente(w, r, int32(id))
	case http.MethodPut:
		s.updateCliente(w, r, int32(id))
	case http.MethodDelete:
		s.deleteCliente(w, r, int32(id))
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func (s *Server) createCliente(w http.ResponseWriter, r *http.Request) {
	// decodificar el JSON del cuerpo de la petición
	var in logic.Cliente
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// validar reglas de negocio con la capa de lógica pura
	if err := logic.ValidateCliente(in.Nombre, in.Apellido, in.Email, in.Direccion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// persistir en la base de datos mediante sqlc
	c, err := s.queries.CrearCliente(r.Context(), db.CrearClienteParams{
		Nombre:    in.Nombre,
		Apellido:  in.Apellido,
		Email:     in.Email,
		Direccion: in.Direccion,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// responder con código 201 Created y el nuevo objeto en formato JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func (s *Server) listClientes(w http.ResponseWriter, r *http.Request) {
	//r.context permite al ususario cancelar la peticion o si se corta la conexion amitad de camino, el contexto se cancela y postgreSQL corta la consulta sin desperdiciar recursos
	lista, err := s.queries.ListarClientes(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//sqlc devuelve nil si no hay filas entonces devolvemos una lista vacia.
	if lista == nil {
		lista = []db.Cliente{}
	}
	// cabecera http fundamental
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lista)
}

func (s *Server) getCliente(w http.ResponseWriter, r *http.Request, id int32) {
	c, err := s.queries.GetCliente(r.Context(), id)
	if err != nil {
		// si no existe la fila, devolvemos 404 Not Found
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Cliente no encontrado", http.StatusNotFound)
			return
		}
		// cualquier otro error es del servidor (500)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}
func (s *Server) updateCliente(w http.ResponseWriter, r *http.Request, id int32) {
	// verificar si el cliente existe antes de actualizar
	if _, err := s.queries.GetCliente(r.Context(), id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Cliente no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// decodificar el JSON entrante
	var in logic.Cliente
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// validar con la lógica pura del negocio
	if err := logic.ValidateCliente(in.Nombre, in.Apellido, in.Email, in.Direccion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// actualizar en la base de datos (sqlc devuelve el cliente actualizado y el error)
	c, err := s.queries.ActualizarCliente(r.Context(), db.ActualizarClienteParams{
		ID:        id,
		Nombre:    in.Nombre,
		Apellido:  in.Apellido,
		Email:     in.Email,
		Direccion: in.Direccion,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// responder con el cliente actualizado en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func (s *Server) deleteCliente(w http.ResponseWriter, r *http.Request, id int32) {
	// verificar si el cliente existe
	if _, err := s.queries.GetCliente(r.Context(), id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Cliente no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// eliminar el registro en PostgreSQL
	if err := s.queries.BorrarCliente(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// responder 204 No Content (sin cuerpo de respuesta)
	w.WriteHeader(http.StatusNoContent)
}

// -----------------------------------------------------------------------------------------------
// CATEGORIAS
// -----------------------------------------------------------------------------------------------

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

	//Puede pasar que la consulta no devuelva ninguna fila.
	//Entonces recibo nil. Pero la consulta fue exitosa, debo responder 200 ok
	//Entonces devuelvo una lista vacia
	if lista == nil {
		lista = []db.Categoria{}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lista)
}

func (s *Server) createCategoria(w http.ResponseWriter, r *http.Request) {
	var categoria logic.Categoria
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		http.Error(w, "JSON invalido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := logic.ValidateCategoria(categoria.Nombre); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nueva_categoria, err := s.queries.CrearCategoria(r.Context(), db.CrearCategoriaParams{
		Nombre:      categoria.Nombre,
		Descripcion: toNullString(categoria.Descripcion),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nueva_categoria)
}

func (s *Server) getCategoria(w http.ResponseWriter, r *http.Request, id int32) {
	categoria, err := s.queries.GetCategoria(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Categoria no encontrada", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categoria)
}

func (s *Server) updateCategoria(w http.ResponseWriter, r *http.Request, id int32) {
	var categoria logic.Categoria
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		http.Error(w, "JSON invalido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := logic.ValidateCategoria(categoria.Nombre); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Si el id no existe, PostgreSQL no devuelve filas y sqlc genera sql.ErrNoRows,
	// por lo que no es necesario hacer un Get previo ni posterior para validar y responder.
	nueva_categoria, err := s.queries.ActualizarCategoria(r.Context(), db.ActualizarCategoriaParams{
		ID:          id,
		Nombre:      categoria.Nombre,
		Descripcion: toNullString(categoria.Descripcion),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Categoria no encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nueva_categoria)
}

func (s *Server) deleteCategoria(w http.ResponseWriter, r *http.Request, id int32) {
	if _, err := s.queries.GetCategoria(r.Context(), id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Categoria no encontrada", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
func aLogic(p db.Producto) (logic.Producto, error) {
	precio, err := strconv.ParseFloat(p.Precio, 64)
	if err != nil {
		return logic.Producto{}, err
	}
	stock, err := strconv.ParseFloat(p.Stock, 64)
	if err != nil {
		return logic.Producto{}, err
	}
	return logic.Producto{
		ID:          p.ID,
		IDCategoria: p.IDCategoria,
		Nombre:      p.Nombre,
		Descripcion: p.Descripcion.String,
		Stock:       stock,
		Precio:      precio,
	}, nil
}

func (s *Server) crearProducto(w http.ResponseWriter, r *http.Request) {
	var p logic.Producto
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "no se pudo interpretar el JSON recibido", http.StatusBadRequest)
		return
	}

	if err := logic.ValidateProduct(p); err != nil {
		http.Error(w, "No se pudo validar el producto", http.StatusBadRequest)
		return
	}

	creado, err := s.queries.CrearProducto(r.Context(), db.CrearProductoParams{
		Nombre:      p.Nombre,
		Descripcion: sql.NullString{String: p.Descripcion, Valid: p.Descripcion != ""},
		Stock:       strconv.FormatFloat(p.Stock, 'f', 2, 64),
		Precio:      strconv.FormatFloat(p.Precio, 'f', 2, 64),
		Foto:        sql.NullString{Valid: false},
		IDCategoria: p.IDCategoria,
	})
	if err != nil {
		http.Error(w, "no se pudo crear el producto", http.StatusInternalServerError)
		return
	}

	resp, err := aLogic(creado)
	if err != nil {
		http.Error(w, "error interno", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) getProducto(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "el id debe ser un número", http.StatusBadRequest)
		return
	}

	obtenido, err := s.queries.GetProducto(r.Context(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "producto no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "error interno", http.StatusInternalServerError)
		return
	}

	resp, err := aLogic(obtenido)
	if err != nil {
		http.Error(w, "error interno", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) listarProductos(w http.ResponseWriter, r *http.Request) {
	lista, err := s.queries.ListarProductos(r.Context())
	if err != nil {
		http.Error(w, "error interno", http.StatusInternalServerError)
		return
	}

	productos := make([]logic.Producto, 0, len(lista))
	for _, p := range lista {
		convertido, err := aLogic(p)
		if err != nil {
			http.Error(w, "error interno", http.StatusInternalServerError)
			return
		}
		productos = append(productos, convertido)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productos)
}
