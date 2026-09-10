-- Reemplazá esta tabla por la entidad principal de TU dominio del TP1.
-- Este es solo el ejemplo que sugiere el enunciado del TP2 (lista de tareas).

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
