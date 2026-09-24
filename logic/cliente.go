package logic

import (
	"errors"
	"strings"
)

func ValidateCliente(nombre, apellido, email, direccion string) error {
	if strings.TrimSpace(nombre) == "" {
		return errors.New("el nombre no puede estar vacio")
	}
	if strings.TrimSpace(apellido) == "" {
		return errors.New("el apellido no puede estar vacio")
	}
	if strings.TrimSpace(email) == "" || !strings.Contains(email, "@") {
		return errors.New("el email es invalido")
	}
	if strings.TrimSpace(direccion) == "" {
		return errors.New("la direccion no puede estar vacia")
	}
	return nil
}
