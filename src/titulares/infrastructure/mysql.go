// mysql.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Romieb26/Arquitectura--hexagonal/src/core"
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain/entities"
)

type MysqlTitular struct {
	conn *sql.DB
}

func NewMysqlTitularRepository() domain.ITitulares {
	conn := core.GetDB()
	return &MysqlTitular{conn: conn}
}

// Save implementa el método de la interfaz ITitulares.
func (mysql *MysqlTitular) Save(titular entities.Titular) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO titulares (nombre, apellido, email, dni, telefono, direccion) VALUES (?, ?, ?, ?, ?, ?)",
		titular.Nombre,
		titular.Apellido,
		titular.Email,
		titular.DNI,
		titular.Telefono,
		titular.Direccion,
	)
	if err != nil {
		log.Println("Error al guardar el titular:", err)
		return err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return err
	}

	titular.SetID(int(idInserted))
	return nil
}

// Update implementa el método de la interfaz ITitulares.
func (mysql *MysqlTitular) Update(id int, titular entities.Titular) error {
	result, err := mysql.conn.Exec(
		"UPDATE titulares SET nombre = ?, apellido = ?, dni = ?, telefono = ?, direccion = ? WHERE id = ?",
		titular.Nombre,
		titular.Apellido,
		titular.DNI,
		titular.Telefono,
		titular.Direccion,
		id,
	)
	if err != nil {
		log.Println("Error al actualizar el titular:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró el titular con ID:", id)
		return fmt.Errorf("titular con ID %d no encontrado", id)
	}

	return nil
}

// Delete implementa el método de la interfaz ITitulares.
func (mysql *MysqlTitular) Delete(id int) error {
	_, err := mysql.conn.Exec("DELETE FROM titulares WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar el titular:", err)
		return err
	}
	return nil
}

// FindByID implementa el método de la interfaz ITitulares.
func (mysql *MysqlTitular) FindByID(id int) (entities.Titular, error) {
	var titular entities.Titular
	row := mysql.conn.QueryRow("SELECT id, nombre, apellido, dni, telefono, direccion FROM titulares WHERE id = ?", id)

	err := row.Scan(
		&titular.ID,
		&titular.Nombre,
		&titular.Apellido,
		&titular.DNI,
		&titular.Telefono,
		&titular.Direccion,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Titular no encontrado:", err)
			return entities.Titular{}, fmt.Errorf("titular con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el titular por ID:", err)
		return entities.Titular{}, err
	}

	return titular, nil
}

// GetAll es un método adicional para obtener todos los titulares.
func (mysql *MysqlTitular) GetAll() ([]entities.Titular, error) {
	var titulares []entities.Titular

	rows, err := mysql.conn.Query("SELECT id, nombre, apellido, dni, telefono, direccion FROM titulares")
	if err != nil {
		log.Println("Error al obtener todos los titulares:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var titular entities.Titular
		err := rows.Scan(
			&titular.ID,
			&titular.Nombre,
			&titular.Apellido,
			&titular.DNI,
			&titular.Telefono,
			&titular.Direccion,
		)
		if err != nil {
			log.Println("Error al escanear titular:", err)
			return nil, err
		}
		titulares = append(titulares, titular)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return titulares, nil
}
