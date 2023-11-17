package models

/*
es mejor conservar un estándar entre las etiquetas de json y db para no tener problemas al parsear
de json a db en el método ActualizarUnAmigo
*/
type Amigo struct {
	Id          int    `db:"id" json:"id"`
	Nombre      string `db:"nombre" json:"nombre"`
	Edad        uint   `db:"edad" json:"edad"`
	Descripcion string `db:"descripcion" json:"descripcion"`
	EstaCerca   bool   `db:"esta_cerca" json:"esta_cerca"`
	Direccion   string `db:"direccion" json:"direccion"`
}
