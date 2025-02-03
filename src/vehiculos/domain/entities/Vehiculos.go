package entities

type Vehiculo struct {
	ID     int    `json:"id"`
	Marca  string `json:"marca"`
	Modelo string `json:"modelo"`
	Año    int    `json:"año"`
	Color  string `json:"color"`
	Placa  string `json:"placa"`
}

func NewVehiculo(id int, marca, modelo string, año int, color, placa string) *Vehiculo {
	return &Vehiculo{
		ID:     id,
		Marca:  marca,
		Modelo: modelo,
		Año:    año,
		Color:  color,
		Placa:  placa,
	}
}

// Getters
func (v *Vehiculo) GetID() int {
	return v.ID
}

func (v *Vehiculo) GetMarca() string {
	return v.Marca
}

func (v *Vehiculo) GetModelo() string {
	return v.Modelo
}

func (v *Vehiculo) GetAño() int {
	return v.Año
}

func (v *Vehiculo) GetColor() string {
	return v.Color
}

func (v *Vehiculo) GetPlaca() string {
	return v.Placa
}

// Setters
func (v *Vehiculo) SetID(id int) {
	v.ID = id
}

func (v *Vehiculo) SetMarca(marca string) {
	v.Marca = marca
}

func (v *Vehiculo) SetModelo(modelo string) {
	v.Modelo = modelo
}

func (v *Vehiculo) SetAño(año int) {
	v.Año = año
}

func (v *Vehiculo) SetColor(color string) {
	v.Color = color
}

func (v *Vehiculo) SetPlaca(placa string) {
	v.Placa = placa
}
