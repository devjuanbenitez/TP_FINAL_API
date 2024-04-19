package model

// Producto es el modelo de datos que representa la tabla public.producto en la base de datos
type Producto struct {
	ID          int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`  // Define la columna ID como llave primaria con auto-incremento
	Nombre      string    `gorm:"size:100;not null;column:nombre" json:"nombre"` // Define la columna Nombre como una cadena de longitud máxima 100 y no nula
	MarcaID     int       `gorm:"column:id_marca" json:"marca"`                  // Define la columna id_marca
	Marca       Marca     `gorm:"foreignKey:MarcaID" json:"-"`                   // Relación con Marca (clave foránea)
	CategoriaID int       `gorm:"column:id_categoria" json:"categoria"`          // Define la columna id_categoria
	Categoria   Categoria `gorm:"foreignKey:CategoriaID" json:"-"`               // Relación con Categoria (clave foránea)
	Precio      float64   `gorm:"not null;column:precio" json:"precio"`          // Define la columna Precio como float y no nula
	Stock       int       `gorm:"not null;column:stock" json:"stock"`            // Define la columna Stock como un entero y no nulo
}

// TableName es una función opcional que puedes utilizar para especificar el nombre de la tabla en la base de datos
func (Producto) TableName() string {
	return "public.producto"
}
