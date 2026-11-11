package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	db "tpe/db/sqlc"
	"tpe/logic"
	"tpe/views"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Server almacena las consultas a la base de datos generadas por sqlc
type Server struct {
	queries *db.Queries
}

// Inicia la conexion a la base de datos, inicializa datos base, configura rutas SSR y levanta el servidor HTTP
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

	// Sembrado preventivo de categorías para asegurar integridad referencial
	ctx := context.Background()
	if err := server.asegurarCategoriasIniciales(ctx); err != nil {
		log.Printf("Advertencia al verificar categorías iniciales: %v\n", err)
	}

	/*=============================================================================
	  RUTAS DE PRESENTACIÓN - SERVER-SIDE RENDERING CON TEMPL (TP5)
	  Siguiendo filmina 19 de 09-view-static.pdf y consigna TP5:
	  - Se abandonan los endpoints de API que devuelven JSON.
	  - GET / renderiza la página completa con Templ.
	  - POST /productos gestiona el formulario clásico con patrón PRG (HTTP 303).
	  - POST /productos/{id}/eliminar gestiona el borrado con patrón PRG (HTTP 303).
	  - /static/ sirve hojas de estilo y recursos estáticos mediante http.FileServer.
	=============================================================================*/
	mux := http.NewServeMux()

	// 1. Servidor de archivos estáticos (/static/styles.css, etc.)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// 2. Página principal con Server-Side Rendering
	mux.HandleFunc("/", server.handleIndex)

	// 3. Manejo de formulario de alta (patrón PRG)
	mux.HandleFunc("/productos", server.handleProductos)

	// 4. Manejo de borrado (patrón PRG)
	mux.HandleFunc("/productos/", server.handleProductoAcciones)

	port := ":8080"
	log.Printf("Servidor SSR con Templ escuchando en http://localhost%s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Error al iniciar el servidor: %s", err)
	}
}

// asegurarCategoriasIniciales inserta categorías de ejemplo si la tabla está vacía
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
		log.Println("Categorías iniciales creadas exitosamente")
	}
	return nil
}

// -----------------------------------------------------------------------------------------------
// HANDLERS SSR Y PATRÓN POST / REDIRECT / GET (PRG)
// -----------------------------------------------------------------------------------------------

// handleIndex atiende GET /:
// - Consulta los productos y categorías usando sqlc.
// - Construye el mapa de nombres de categorías.
// - Renderiza el componente IndexPage (que usa layout.templ, entity_form.templ y entity_list.templ).
func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	// Obtener lista de productos
	productos, err := s.queries.ListarProductos(ctx)
	if err != nil {
		log.Printf("Error al listar productos: %v", err)
		http.Error(w, "Error al consultar productos", http.StatusInternalServerError)
		return
	}

	// Obtener lista de categorías
	categorias, err := s.queries.ListarCategorias(ctx)
	if err != nil {
		log.Printf("Error al listar categorías: %v", err)
		http.Error(w, "Error al consultar categorías", http.StatusInternalServerError)
		return
	}

	// Mapa id -> nombre para resolver la categoría en la tabla
	catMap := make(map[int32]string, len(categorias))
	for _, c := range categorias {
		catMap[c.ID] = c.Nombre
	}

	// Mensajes de retroalimentación transferidos mediante query params por el patrón PRG
	errorMsg := r.URL.Query().Get("error")
	successMsg := r.URL.Query().Get("success")

	// Renderizar componente Templ directamente en el http.ResponseWriter
	component := views.IndexPage(productos, categorias, catMap, errorMsg, successMsg)
	if err := component.Render(ctx, w); err != nil {
		log.Printf("Error al renderizar IndexPage: %v", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
	}
}

// handleProductos despacha peticiones sobre /productos
func (s *Server) handleProductos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.crearProductoForm(w, r)
	default:
		// Redirige al inicio si navegan por GET /productos directamente
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// crearProductoForm implementa el patrón Post/Redirect/Get (PRG):
// 1. Parsea los datos del formulario (application/x-www-form-urlencoded).
// 2. Valida tipos y reglas de negocio con la capa pura `logic.ValidateProduct`.
// 3. Persiste en la base de datos usando sqlc (`queries.CrearProducto`).
// 4. Redirige a GET / con código 303 (See Other) para actualizar la vista sin riesgo de reenvío.
func (s *Server) crearProductoForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		redireccionarConError(w, r, "Datos de formulario inválidos")
		return
	}

	nombre := strings.TrimSpace(r.FormValue("nombre"))
	descripcion := strings.TrimSpace(r.FormValue("descripcion"))
	precioStr := strings.TrimSpace(r.FormValue("precio"))
	stockStr := strings.TrimSpace(r.FormValue("stock"))
	catIDStr := strings.TrimSpace(r.FormValue("id_categoria"))

	// Conversión y validación básica de números
	precio, err := strconv.ParseFloat(precioStr, 64)
	if err != nil {
		redireccionarConError(w, r, "El precio ingresado no es un número válido")
		return
	}

	stock, err := strconv.ParseFloat(stockStr, 64)
	if err != nil {
		redireccionarConError(w, r, "El stock ingresado no es un número válido")
		return
	}

	idCategoria, err := strconv.Atoi(catIDStr)
	if err != nil || idCategoria <= 0 {
		redireccionarConError(w, r, "Debe seleccionar una categoría válida")
		return
	}

	// Validar con la capa de lógica del negocio pura (Three-Tier Architecture)
	prodLogic := logic.Producto{
		Nombre:      nombre,
		Descripcion: descripcion,
		Precio:      precio,
		Stock:       stock,
		IDCategoria: int32(idCategoria),
	}
	if err := logic.ValidateProduct(prodLogic); err != nil {
		redireccionarConError(w, r, err.Error())
		return
	}

	// Verificar existencia de la categoría en la base de datos
	ctx := r.Context()
	if _, err := s.queries.GetCategoria(ctx, int32(idCategoria)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			redireccionarConError(w, r, "La categoría seleccionada no existe")
			return
		}
		log.Printf("Error al consultar categoría: %v", err)
		redireccionarConError(w, r, "Error al verificar la categoría")
		return
	}

	// Persistir con sqlc
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
		log.Printf("Error al insertar producto: %v", err)
		redireccionarConError(w, r, "No se pudo guardar el producto en la base de datos")
		return
	}

	// Patrón PRG: Redirección con HTTP 303 See Other
	msg := fmt.Sprintf("Producto '%s' creado exitosamente", nombre)
	http.Redirect(w, r, "/?success="+url.QueryEscape(msg), http.StatusSeeOther)
}

// handleProductoAcciones atiende rutas como POST /productos/{id}/eliminar o POST /productos/eliminar
func (s *Server) handleProductoAcciones(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Extraer ID del path (/productos/{id}/eliminar) o del formulario
	path := strings.TrimPrefix(r.URL.Path, "/productos/")
	partes := strings.Split(strings.Trim(path, "/"), "/")

	var id int
	var err error

	if len(partes) >= 1 && partes[0] != "" && partes[0] != "eliminar" {
		id, err = strconv.Atoi(partes[0])
	} else {
		// Intento alternativo leyendo campo oculto 'id'
		_ = r.ParseForm()
		id, err = strconv.Atoi(r.FormValue("id"))
	}

	if err != nil || id <= 0 {
		redireccionarConError(w, r, "Identificador de producto inválido")
		return
	}

	ctx := r.Context()

	// Verificar existencia
	prod, err := s.queries.GetProducto(ctx, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			redireccionarConError(w, r, "El producto que intenta eliminar no existe")
			return
		}
		log.Printf("Error al obtener producto para borrar: %v", err)
		redireccionarConError(w, r, "Error al verificar producto")
		return
	}

	// Borrar con sqlc
	if err := s.queries.BorrarProducto(ctx, int32(id)); err != nil {
		log.Printf("Error al borrar producto: %v", err)
		redireccionarConError(w, r, "No se pudo eliminar el producto")
		return
	}

	// Patrón PRG tras eliminación
	msg := fmt.Sprintf("Producto '%s' (#%d) eliminado correctamente", prod.Nombre, id)
	http.Redirect(w, r, "/?success="+url.QueryEscape(msg), http.StatusSeeOther)
}

// redireccionarConError redirige al usuario a la página principal con el mensaje de error escapado
func redireccionarConError(w http.ResponseWriter, r *http.Request, msg string) {
	http.Redirect(w, r, "/?error="+url.QueryEscape(msg), http.StatusSeeOther)
}
