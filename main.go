package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	db "tpe/db/sqlc"
	"tpe/logic"
	"tpe/views"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// almacena las consultas a la base de datos
type Server struct {
	queries *db.Queries
}

// inicia la conexion a la base de datos, datos iniciales y servidor http
func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Falta la variable de entorno DATABASE_URL")
	}

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("No se pudo abrir la conexion: %v", err)
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	queries := db.New(conn)
	server := &Server{
		queries: queries,
	}
	log.Println("Conectado a la base de datos correctamente")

	// crea categorias iniciales si la tabla esta vacia
	ctx := context.Background()
	if err := server.asegurarCategoriasIniciales(ctx); err != nil {
		log.Printf("Advertencia al verificar categorias iniciales: %v\n", err)
	}

	mux := http.NewServeMux()

	// archivos estaticos: css, fuentes, etc.
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// pagina principal
	mux.HandleFunc("/", server.handleIndex)

	// productos: listado htmx o alta
	mux.HandleFunc("/productos", server.handleProductos)

	// operaciones sobre un producto especifico: borrado
	mux.HandleFunc("/productos/", server.handleProductoByID)

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Error al iniciar el servidor: %s", err)
	}
}

// inserta categorias de ejemplo si la tabla esta vacia
func (s *Server) asegurarCategoriasIniciales(ctx context.Context) error {
	categorias, err := s.queries.ListarCategorias(ctx)
	if err != nil {
		return err
	}
	if len(categorias) == 0 {
		categoriasBase := []struct {
			nombre string
			desc   string
		}{
			{"Almacén", "Comestibles y envasados"},
			{"Bebidas", "Gaseosas, jugos y aguas"},
			{"Lácteos", "Quesos, yogures y leches"},
			{"Limpieza", "Artículos de limpieza para el hogar"},
		}
		for _, cat := range categoriasBase {
			_, _ = s.queries.CrearCategoria(ctx, db.CrearCategoriaParams{
				Nombre:      cat.nombre,
				Descripcion: sql.NullString{String: cat.desc, Valid: true},
			})
		}
		log.Println("Categorias iniciales creadas exitosamente")
	}
	return nil
}

// consulta productos y categorias para renderizar el listado
func (s *Server) obtenerDatosListado(ctx context.Context) ([]db.Producto, []db.Categoria, map[int32]string, error) {
	productos, err := s.queries.ListarProductos(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	categorias, err := s.queries.ListarCategorias(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	catMap := make(map[int32]string, len(categorias))
	for _, c := range categorias {
		catMap[c.ID] = c.Nombre
	}
	return productos, categorias, catMap, nil
}

// muestra la pagina principal con el formulario y la lista de productos
func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	productos, categorias, catMap, err := s.obtenerDatosListado(ctx)
	if err != nil {
		log.Printf("error al obtener datos: %v", err)
		http.Error(w, "error al consultar datos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := views.IndexPage(productos, categorias, catMap).Render(ctx, w); err != nil {
		log.Printf("error al renderizar index: %v", err)
		http.Error(w, "error interno del servidor", http.StatusInternalServerError)
	}
}

// despacha peticiones sobre /productos
func (s *Server) handleProductos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.crearProductoHTMX(w, r)
	case http.MethodGet:
		// si es htmx devuelve solo la lista, sino redirige a la pagina principal
		if r.Header.Get("HX-Request") == "true" {
			ctx := r.Context()
			productos, _, catMap, err := s.obtenerDatosListado(ctx)
			if err != nil {
				log.Printf("error al obtener productos: %v", err)
				http.Error(w, "error al consultar productos", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			if err := views.EntityList(productos, catMap).Render(ctx, w); err != nil {
				log.Printf("error al renderizar lista: %v", err)
				http.Error(w, "error interno del servidor", http.StatusInternalServerError)
			}
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		http.Error(w, "metodo no permitido", http.StatusMethodNotAllowed)
	}
}

// envia un fragmento de error para htmx
func renderErrorHTMX(w http.ResponseWriter, r *http.Request, status int, msg string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if err := views.FeedbackPartial("error", msg).Render(r.Context(), w); err != nil {
		log.Printf("error al renderizar feedback: %v", err)
	}
}

// procesa el formulario de alta y devuelve la lista actualizada
func (s *Server) crearProductoHTMX(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		renderErrorHTMX(w, r, http.StatusBadRequest, "datos de formulario invalidos")
		return
	}

	nombre := strings.TrimSpace(r.FormValue("nombre"))
	descripcion := strings.TrimSpace(r.FormValue("descripcion"))
	precioStr := strings.TrimSpace(r.FormValue("precio"))
	stockStr := strings.TrimSpace(r.FormValue("stock"))
	catIDStr := strings.TrimSpace(r.FormValue("id_categoria"))

	precio, err := strconv.ParseFloat(precioStr, 64)
	if err != nil {
		renderErrorHTMX(w, r, http.StatusBadRequest, "el precio ingresado no es valido")
		return
	}

	stock, err := strconv.ParseFloat(stockStr, 64)
	if err != nil {
		renderErrorHTMX(w, r, http.StatusBadRequest, "el stock ingresado no es valido")
		return
	}

	idCategoria, err := strconv.Atoi(catIDStr)
	if err != nil || idCategoria <= 0 {
		renderErrorHTMX(w, r, http.StatusBadRequest, "debe seleccionar una categoria valida")
		return
	}

	prodLogic := logic.Producto{
		Nombre:      nombre,
		Descripcion: descripcion,
		Precio:      precio,
		Stock:       stock,
		IDCategoria: int32(idCategoria),
	}
	if err := logic.ValidateProduct(prodLogic); err != nil {
		renderErrorHTMX(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	if _, err := s.queries.GetCategoria(ctx, int32(idCategoria)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			renderErrorHTMX(w, r, http.StatusBadRequest, "la categoria seleccionada no existe")
			return
		}
		log.Printf("error al consultar categoria: %v", err)
		renderErrorHTMX(w, r, http.StatusInternalServerError, "error al verificar la categoria")
		return
	}

	var descNull sql.NullString
	if descripcion != "" {
		descNull = sql.NullString{String: descripcion, Valid: true}
	}

	_, err = s.queries.CrearProducto(ctx, db.CrearProductoParams{
		Nombre:      nombre,
		Descripcion: descNull,
		Stock:       strconv.FormatFloat(stock, 'f', 2, 64),
		Precio:      strconv.FormatFloat(precio, 'f', 2, 64),
		Foto:        sql.NullString{Valid: false},
		IDCategoria: int32(idCategoria),
	})
	if err != nil {
		log.Printf("error al insertar producto: %v", err)
		renderErrorHTMX(w, r, http.StatusInternalServerError, "no se pudo guardar el producto")
		return
	}

	productos, categorias, catMap, err := s.obtenerDatosListado(ctx)
	if err != nil {
		log.Printf("error al consultar lista actualizada: %v", err)
		renderErrorHTMX(w, r, http.StatusInternalServerError, "error al actualizar el listado")
		return
	}

	msg := fmt.Sprintf("Producto '%s' creado exitosamente", nombre)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := views.EntityListActualizada(productos, categorias, catMap, msg).Render(ctx, w); err != nil {
		log.Printf("error al renderizar lista actualizada: %v", err)
		http.Error(w, "error interno del servidor", http.StatusInternalServerError)
	}
}

// maneja operaciones sobre /productos/{id}
func (s *Server) handleProductoByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	path := strings.Trim(strings.TrimPrefix(r.URL.Path, "/productos/"), "/")
	if path == "" || strings.Contains(path, "/") {
		renderErrorHTMX(w, r, http.StatusBadRequest, "identificador de producto invalido")
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil || id <= 0 {
		renderErrorHTMX(w, r, http.StatusBadRequest, "identificador de producto invalido")
		return
	}

	ctx := r.Context()
	if _, err := s.queries.GetProducto(ctx, int32(id)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			renderErrorHTMX(w, r, http.StatusNotFound, "el producto no existe")
			return
		}
		log.Printf("error al obtener producto: %v", err)
		renderErrorHTMX(w, r, http.StatusInternalServerError, "error al verificar producto")
		return
	}

	if err := s.queries.BorrarProducto(ctx, int32(id)); err != nil {
		log.Printf("error al borrar producto: %v", err)
		renderErrorHTMX(w, r, http.StatusInternalServerError, "no se pudo eliminar el producto")
		return
	}

	// 200 ok con cuerpo vacio para que htmx elimine la fila
	w.WriteHeader(http.StatusOK)
}
