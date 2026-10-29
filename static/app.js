// ==============================================================================
// MarkeTandil - Capa de Presentación (Frontend con Vanilla JS y DOM)
// Trabajo Práctico Especial 4 - Programación Web (UNICEN)
// Cátedra: Dr. Alejandro Zunino & Dr. Alfredo Teyseyre
// ==============================================================================

// URL base para interactuar con la API REST de productos
const ENDPOINT_PRODUCTOS = '/productos';

// Esperamos a que el documento HTML este completamente parseado
document.addEventListener('DOMContentLoaded', () => {
    // Referencias a los elementos del DOM que necesitamos manipular
    const listaProductos = document.getElementById('lista-productos');
    const formProducto = document.getElementById('form-producto');
    const mensajeEstado = document.getElementById('mensaje-estado');
    const btnRefrescar = document.getElementById('btn-refrescar');

    // 1. Cargar el listado inicial de productos al abrir la pagina
    cargarProductos();

    // 2. Escuchar el envio del formulario para crear un producto (POST)
    formProducto.addEventListener('submit', (e) => {
        // Prevenimos el comportamiento por defecto del navegador (recarga de pagina)
        e.preventDefault();
        crearProducto();
    });

    // 3. Delegacion de eventos para los botones de borrado (DELETE)
    // Como los botones de eliminar se generan dinamicamente cuando recibimos los datos,
    // colocamos el escuchador en el contenedor padre existente (#lista-productos).
    listaProductos.addEventListener('click', (e) => {
        // Buscamos si el clic provino de un boton con la clase "btn-eliminar"
        const btnEliminar = e.target.closest('.btn-eliminar');
        if (!btnEliminar) return;

        // Obtenemos el ID del producto almacenado en el atributo data-id
        const id = btnEliminar.dataset.id;
        const nombre = btnEliminar.dataset.nombre || 'este producto';

        if (confirm(`¿Estas seguro de que queres eliminar "${nombre}" (ID: ${id})?`)) {
            eliminarProducto(id);
        }
    });

    // Boton opcional para refrescar manualmente la lista
    if (btnRefrescar) {
        btnRefrescar.addEventListener('click', () => {
            cargarProductos();
        });
    }

    // ==========================================================================
    // FUNCIONES DE COMUNICACION CON LA API Y MANIPULACION DEL DOM
    // ==========================================================================

    /**
     * Realiza una peticion GET /productos para obtener el listado en formato JSON
     * y genera dinamicamente los elementos HTML en el contenedor.
     */
    function cargarProductos() {
        listaProductos.innerHTML = '<li class="cargando">Cargando productos...</li>';

        fetch(ENDPOINT_PRODUCTOS)
            .then((response) => {
                if (!response.ok) {
                    throw new Error(`Error del servidor (${response.status}): ${response.statusText}`);
                }
                return response.json();
            })
            .then((productos) => {
                renderizarProductos(productos);
            })
            .catch((error) => {
                console.error('Error al obtener los productos:', error);
                listaProductos.innerHTML = `<li class="error">No se pudieron cargar los productos: ${error.message}</li>`;
            });
    }

    /**
     * Dibuja dinamicamente en el DOM la lista de productos recibida por parametro.
     */
    function renderizarProductos(productos) {
        // Limpiamos el contenedor
        listaProductos.innerHTML = '';

        if (!productos || productos.length === 0) {
            listaProductos.innerHTML = '<li class="vacio">No hay productos registrados actualmente.</li>';
            return;
        }

        // Creamos cada elemento <li> con sus datos y su boton de borrado
        productos.forEach((p) => {
            const li = document.createElement('li');
            li.className = 'item-producto';

            // Armamos la descripcion limpia o un texto por defecto
            const desc = p.descripcion ? p.descripcion : 'Sin descripcion';

            li.innerHTML = `
                <div class="info-producto">
                    <span class="id-producto">#${p.id}</span>
                    <strong class="nombre-producto">${escaparHTML(p.nombre)}</strong>
                    <span class="categoria-producto">Cat. ID: ${p.id_categoria}</span>
                    <p class="desc-producto">${escaparHTML(desc)}</p>
                    <div class="detalles-producto">
                        <span class="precio-producto">$${Number(p.precio).toFixed(2)}</span>
                        <span class="stock-producto">Stock: ${p.stock}</span>
                    </div>
                </div>
                <div class="acciones-producto">
                    <button type="button" class="btn-eliminar" data-id="${p.id}" data-nombre="${escaparHTML(p.nombre)}">
                        Eliminar
                    </button>
                </div>
            `;

            listaProductos.appendChild(li);
        });
    }

    /**
     * Toma los datos del formulario, arma el objeto JSON y envia una peticion POST /productos.
     */
    function crearProducto() {
        // Extraemos y parseamos los valores ingresados por el usuario
        const nombreInput = document.getElementById('nombre');
        const descInput = document.getElementById('descripcion');
        const precioInput = document.getElementById('precio');
        const stockInput = document.getElementById('stock');
        const catInput = document.getElementById('id_categoria');

        const nuevoProducto = {
            nombre: nombreInput.value.trim(),
            descripcion: descInput.value.trim(),
            precio: parseFloat(precioInput.value),
            stock: parseFloat(stockInput.value),
            id_categoria: parseInt(catInput.value, 10),
        };

        // Deshabilitamos el boton mientras se procesa la peticion
        const btnGuardar = document.getElementById('btn-crear');
        btnGuardar.disabled = true;

        fetch(ENDPOINT_PRODUCTOS, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(nuevoProducto),
        })
            .then(async (response) => {
                if (!response.ok) {
                    // Si el servidor devolvio un error (ej. 400 Bad Request), leemos el mensaje de texto
                    const errorMsg = await response.text();
                    throw new Error(errorMsg || `Error al crear (${response.status})`);
                }
                return response.json();
            })
            .then((productoCreado) => {
                mostrarMensaje(`Producto "${productoCreado.nombre}" creado exitosamente.`, 'exito');
                // Limpiamos los campos del formulario
                formProducto.reset();
                // Dejamos la categoria por defecto en 1
                catInput.value = '1';
                // Refrescamos la lista para mostrar el nuevo item agregado
                cargarProductos();
            })
            .catch((error) => {
                console.error('Error al crear el producto:', error);
                mostrarMensaje(`Error: ${error.message}`, 'error');
            })
            .finally(() => {
                btnGuardar.disabled = false;
            });
    }

    /**
     * Realiza una peticion DELETE /productos/{id} y refresca la lista tras la respuesta exitosa.
     */
    function eliminarProducto(id) {
        fetch(`${ENDPOINT_PRODUCTOS}/${id}`, {
            method: 'DELETE',
        })
            .then(async (response) => {
                if (!response.ok) {
                    const errorMsg = await response.text();
                    throw new Error(errorMsg || `Error al eliminar (${response.status})`);
                }
                // Status 204 No Content indica exito
                mostrarMensaje(`Producto #${id} eliminado correctamente.`, 'exito');
                // Volvemos a pedir la lista completa para sincronizar la vista con la base de datos
                cargarProductos();
            })
            .catch((error) => {
                console.error('Error al eliminar el producto:', error);
                mostrarMensaje(`Error al eliminar: ${error.message}`, 'error');
            });
    }

    /**
     * Muestra un mensaje temporal de exito o error al usuario.
     */
    function mostrarMensaje(texto, tipo) {
        mensajeEstado.textContent = texto;
        mensajeEstado.className = `mensaje ${tipo}`;

        // Quitar el mensaje automaticamente luego de 4 segundos
        setTimeout(() => {
            mensajeEstado.className = 'mensaje-oculto';
            mensajeEstado.textContent = '';
        }, 4000);
    }

    /**
     * Escapa caracteres especiales para evitar inyecciones XSS al renderizar texto del usuario.
     */
    function escaparHTML(texto) {
        if (!texto) return '';
        const div = document.createElement('div');
        div.textContent = texto;
        return div.innerHTML;
    }
});

