// entities/Titulares.go
package entities

import "github.com/Romieb26/Arquitectura--hexagonal/src/core"

type Titular struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Email     string `json:"email"`
	DNI       string `json:"dni"`
	Telefono  int16  `json:"telefono"`
	Direccion string `json:"direccion"`
	DNIRaw    string `json:"-"`
}

type EmailData struct {
	FromID  int    `json:"from_id"`
	ToID    int    `json:"to_id"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func NewTitular(id int, nombre, apellido, email, dni string, telefono int16, direccion string) (*Titular, error) {
	dniEncrypted, err := core.EncryptPassword(dni)
	if err != nil {
		return nil, err
	}

	return &Titular{
		ID:        id,
		Nombre:    nombre,
		Apellido:  apellido,
		Email:     email,
		DNI:       dniEncrypted,
		DNIRaw:    dni,
		Telefono:  telefono,
		Direccion: direccion,
	}, nil
}

func (t *Titular) GetID() int {
	return t.ID
}

func (t *Titular) SetID(id int) {
	t.ID = id
}

func (t *Titular) GetNombre() string {
	return t.Nombre
}

func (t *Titular) SetNombre(nombre string) {
	t.Nombre = nombre
}

func (t *Titular) GetApellido() string {
	return t.Apellido
}

func (t *Titular) SetApellido(apellido string) {
	t.Apellido = apellido
}

func (t *Titular) GetDNI() string {
	return t.DNI
}

func (t *Titular) SetDNI(dni string) {
	t.DNI = dni
}

func (t *Titular) GetTelefono() int16 {
	return t.Telefono
}

func (t *Titular) SetTelefono(telefono int16) {
	t.Telefono = telefono
}

func (t *Titular) GetDireccion() string {
	return t.Direccion
}

func (t *Titular) SetDireccion(direccion string) {
	t.Direccion = direccion
}
