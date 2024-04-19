package model

type Marca struct {
	ID     int    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`  // Define la columna ID como llave primaria con auto-incremento
	Nombre string `gorm:"size:100;not null;column:nombre" json:"nombre"` // Define la columna Nombre como una cadena de longitud máxima 100 y no nula
}

func (Marca) TableName() string {
	return "public.marca"
}
