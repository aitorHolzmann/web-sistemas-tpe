package logic

import (
	"errors"
	"strings"
)

type Categoria struct {
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

// valida que el nombre de la categoria no este vacio
func ValidateCategoria(nombre string) error {
	if strings.TrimSpace(nombre) == "" {
		return errors.New("el nombre de la categoria no puede estar vacio")
	}
	return nil
}
