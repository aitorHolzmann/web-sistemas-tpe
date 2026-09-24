---
name: web-unicen-presentacion-ssr-htmx
description: >-
  Guía y patrones de la cátedra de Programación Web (UNICEN) para la capa de Presentación.
  Usar para TP1, manejo de formularios HTML, plantillas con html/template, renderizado en servidor (SSR),
  y dinamismo en el cliente con Vanilla JS fetch y HTMX.
---

# Capa de Presentación: SSR, Plantillas y HTMX (Cátedra Web UNICEN)

Esta skill reúne los conceptos y patrones oficiales de la cátedra para la capa de Presentación, basados en:
- `teoria/01-intro.md`: Modelo cliente-servidor, HTTP, HTML y formularios/CGI.
- `teoria/08-view.md`: DOM, manipulación con JavaScript, SOP y CORS.
- `teoria/09-view-static.md`: SSR tradicional con `html/template` vs `templ`.
- `teoria/10-view-dynamic.md`: Actualizaciones parciales del DOM con Fetch y HTMX.

---

## 1. Servidor de Archivos Estáticos y Formularios Básicos (Estilo TP1)

```go
// Servir directorio estático completo:
fileServer := http.FileServer(http.Dir("./static"))
http.Handle("/static/", http.StripPrefix("/static/", fileServer))

// Procesar formulario POST clásico:
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// 1. Parsear datos del formulario
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al parsear formulario", http.StatusBadRequest)
		return
	}

	// 2. Extraer campos
	nombre := r.FormValue("nombre")
	// Procesar y responder...
}
```

---

## 2. Server-Side Rendering con `html/template`

> [!NOTE]
> La cátedra enfatiza usar `html/template` en lugar de `text/template` debido al escape contextual automático de caracteres especiales que previene vulnerabilidades de Cross-Site Scripting (XSS).

```go
func renderHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/content.html")
	if err != nil {
		http.Error(w, "Error al cargar plantilla", http.StatusInternalServerError)
		return
	}

	data := struct {
		Titulo  string
		Items   []string
	}{
		Titulo: "Listado de Productos",
		Items:  []string{"Teclado", "Mouse", "Monitor"},
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)
}
```

---

## 3. Dinamismo en Cliente con HTMX

La cátedra propone HTMX como la alternativa declarativa "HTML-first" para actualizar fragmentos del DOM sin recargar la página completa.

### Principales atributos HTMX enseñados:
- `hx-get="/endpoint"` / `hx-post="/endpoint"`: Dispara la solicitud HTTP.
- `hx-target="#contenedor"`: Elemento del DOM que se actualizará con la respuesta.
- `hx-swap="innerHTML"` (o `outerHTML`): Cómo se inserta la respuesta.
- `hx-trigger="click"` (o `keyup changed delay:500ms` para búsquedas).

### Ejemplo de Fragmento HTML en Go:
```go
// Handler que devuelve solo un fragmento de HTML para HTMX
func userTableFragmentHandler(w http.ResponseWriter, r *http.Request) {
	// Se renderiza únicamente la porción <table>...</table>, no el <html> completo
	tmpl := template.Must(template.New("fragment").Parse(`
		<tr id="user-{{.ID}}">
			<td>{{.Name}}</td>
			<td>{{.Email}}</td>
		</tr>
	`))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, usuario)
}
```

