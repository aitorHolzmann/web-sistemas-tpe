# Guía Completa de Handlers Implementados (TP3 - Programación Web UNICEN)

Este archivo contiene todos los handlers implementados en [`main.go`](file:///home/esteban/Downloads/ArquiWeb/tpe/tpe3/main.go) para las 3 tablas (`cliente`, `categoria`, `producto`), junto con el middleware de logging y las rutas en `main()`. Todos incluyen comentarios breves en minúscula antes de cada función.

---

## 1. Configuración en `main()` y Funciones Auxiliares (`main.go`)

```go
// inicia la conexion a la base de datos, configura las rutas y levanta el servidor http
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

	server := &Server{
		queries: db.New(conn),
	}
	log.Println("Conectado a la base de datos correctamente")

	/*====================================
	END-POINTS
	=====================================*/
	mux := http.NewServeMux()
	mux.Handle("/categorias", loggingMiddleware(http.HandlerFunc(server.categoriasHandler)))
	mux.Handle("/categorias/", loggingMiddleware(http.HandlerFunc(server.categoriasHandler)))
	mux.Handle("/productos", loggingMiddleware(http.HandlerFunc(server.productosHandler)))
	mux.Handle("/productos/", loggingMiddleware(http.HandlerFunc(server.productosHandler)))
	//para clientes tengo en cuenta clientes y clientes/ por si piden listar clientes o si piden algo especifico clientes/...
	mux.Handle("/clientes", loggingMiddleware(http.HandlerFunc(server.clientesHandler)))
	mux.Handle("/clientes/", loggingMiddleware(http.HandlerFunc(server.clientesHandler)))
	mux.Handle("/", loggingMiddleware(http.FileServer(http.Dir("./static"))))

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Error al iniciar el servidor: %s", err)
	}
}

// registra en consola el metodo, la ruta y la ip de cada peticion antes de pasarla al handler
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s desde %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// como en la db tenemos atributos que pueden ser null, de esta forma puedo resolverlo para strings
func toNullString(s string) sql.NullString {
	if strings.TrimSpace(s) == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}
```

---

## 2. Handlers CRUD de `Cliente` (`main.go`)

```go
// analiza la ruta de clientes y deriva la peticion segun el metodo http y si tiene id
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

// decodifica un cliente del json, lo valida y lo guarda en la base de datos
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

// obtiene todos los clientes de la base de datos y los devuelve como array json
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

// busca un cliente por su id y devuelve 404 si no existe
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

// valida los nuevos datos y actualiza un cliente existente por su id
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

// verifica que el cliente exista, lo elimina por id y devuelve 204 sin contenido
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
```

---

## 3. Handlers CRUD de `Producto` (`main.go`)

```go
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

// analiza la ruta de productos y deriva la peticion segun el metodo http y si tiene id
func (s *Server) productosHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/productos")
	path = strings.Trim(path, "/")

	if path == "" {
		switch r.Method {
		case http.MethodGet:
			s.listarProductos(w, r)
		case http.MethodPost:
			s.crearProducto(w, r)
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
		s.getProducto(w, r, int32(id))
	case http.MethodPut:
		s.updateProducto(w, r, int32(id))
	case http.MethodDelete:
		s.deleteProducto(w, r, int32(id))
	default:
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
	}
}

// decodifica un producto del json, valida que exista su categoria y lo guarda en la db
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

	if _, err := s.queries.GetCategoria(r.Context(), p.IDCategoria); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "la categoria especificada no existe", http.StatusBadRequest)
			return
		}
		http.Error(w, "error interno", http.StatusInternalServerError)
		return
	}

	creado, err := s.queries.CrearProducto(r.Context(), db.CrearProductoParams{
		Nombre:      p.Nombre,
		Descripcion: toNullString(p.Descripcion),
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

// busca un producto por su id, lo traduce a la capa de logica y devuelve 404 si no existe
func (s *Server) getProducto(w http.ResponseWriter, r *http.Request, id int32) {
	obtenido, err := s.queries.GetProducto(r.Context(), id)
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

// obtiene todos los productos de la db, los traduce a logica y los devuelve como array json
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

// valida los nuevos datos de un producto y lo actualiza en la base de datos por su id
func (s *Server) updateProducto(w http.ResponseWriter, r *http.Request, id int32) {
	var p logic.Producto
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "no se pudo interpretar el JSON recibido", http.StatusBadRequest)
		return
	}

	if err := logic.ValidateProduct(p); err != nil {
		http.Error(w, "No se pudo validar el producto", http.StatusBadRequest)
		return
	}

	if _, err := s.queries.GetCategoria(r.Context(), p.IDCategoria); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "la categoria especificada no existe", http.StatusBadRequest)
			return
		}
		http.Error(w, "error interno", http.StatusInternalServerError)
		return
	}

	actualizado, err := s.queries.ActualizarProducto(r.Context(), db.ActualizarProductoParams{
		ID:          id,
		Nombre:      p.Nombre,
		Descripcion: toNullString(p.Descripcion),
		Stock:       strconv.FormatFloat(p.Stock, 'f', 2, 64),
		Precio:      strconv.FormatFloat(p.Precio, 'f', 2, 64),
		Foto:        sql.NullString{Valid: false},
		IDCategoria: p.IDCategoria,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "producto no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "no se pudo actualizar el producto", http.StatusInternalServerError)
		return
	}

	resp, err := aLogic(actualizado)
	if err != nil {
		http.Error(w, "error interno", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// verifica que el producto exista, lo elimina por su id y responde 204 sin contenido
func (s *Server) deleteProducto(w http.ResponseWriter, r *http.Request, id int32) {
	if _, err := s.queries.GetProducto(r.Context(), id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "producto no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "error interno", http.StatusInternalServerError)
		return
	}

	if err := s.queries.BorrarProducto(r.Context(), id); err != nil {
		http.Error(w, "no se pudo eliminar el producto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
```
