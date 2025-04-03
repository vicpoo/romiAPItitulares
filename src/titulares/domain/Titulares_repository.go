// Titulares_repository.go
package domain

import "github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain/entities"

type ITitulares interface {
	Save(titular entities.Titular) error
	Update(id int, titular entities.Titular) error
	Delete(id int) error
	FindByID(id int) (entities.Titular, error)
	GetAll() ([]entities.Titular, error)
}
